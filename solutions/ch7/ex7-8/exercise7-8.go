/*
Exercise 7.8: Many GUI sprovide a table widget with a stateful multi-tier sort:the primary
sort key is the most recently clicked column head, the secondary sort key is the second-most
recently clicked column head, and soon. Deﬁne an implementation of sort.Interface for
use by such a table.Compare that approach with repeated sorting using sort.Stable.
*/

//finish date:20170717 7:37am

// See page 187.

package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Jave", "From the Jave Up", 2017, length("3m38s")},
	{"Go", "Jave", "From the Jave Up1", 2016, length("3m38s")},
	{"Go", "Jave", "From the Jave Up2", 2015, length("3m38s")},
	{"Go", "Jave", "From the Jave Up3", 2015, length("3m38s")},
	{"Go", "Jave", "From the Jave Up4", 2015, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!-main

//!+printTracks
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

//!-printTracks

//!+artistcode
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-artistcode

//!+yearcode
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-yearcode
const BY_NULL = 0
const BY_TITLE = 1
const BY_ARTIST = 2
const BY_ALBUM = 3
const BY_YEAR = 4
const BY_LENGTH = 5

var RecentlyTier1 int = BY_NULL
var RecentlyTier2 int = BY_NULL
var RecentlyTier3 int = BY_NULL

func SetRecentlySortTier(sortKey int) {
	RecentlyTier3 = RecentlyTier2
	RecentlyTier2 = RecentlyTier1
	RecentlyTier1 = sortKey
}

func SortKeyEnable(key int, x, y *Track) bool {
	if key != BY_NULL {
		switch key {
		case BY_TITLE:
			return x.Title != y.Title
		case BY_ARTIST:
			return x.Artist != y.Artist
		case BY_ALBUM:
			return x.Album != y.Album
		case BY_YEAR:
			return x.Year != y.Year
		case BY_LENGTH:
			return x.Length != y.Length
		}
	}
	return false
}

func SortKeyImp(key int, x, y *Track) bool {

	switch key {
	case BY_TITLE:
		return x.Title < y.Title
	case BY_ARTIST:
		return x.Artist < y.Artist
	case BY_ALBUM:
		return x.Album < y.Album
	case BY_YEAR:
		return x.Year < y.Year
	case BY_LENGTH:
		return x.Length < y.Length
	}

	return false
}

func main() {
	fmt.Println("byArtist:")
	sort.Sort(byArtist(tracks))

	printTracks(tracks)

	fmt.Println("\nReverse(byArtist):")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println("\nbyYear:")
	sort.Sort(byYear(tracks))

	printTracks(tracks)

	//TEST CODE
	SetRecentlySortTier(BY_ALBUM)
	SetRecentlySortTier(BY_ARTIST)
	SetRecentlySortTier(BY_YEAR)
	//END TEST CODE

	fmt.Println("\nCustom:")
	//!+customcall
	sort.Sort(customSort{tracks, func(x, y *Track) bool {

		if SortKeyEnable(RecentlyTier1, x, y) == true {
			return SortKeyImp(RecentlyTier1, x, y)
		}
		if SortKeyEnable(RecentlyTier2, x, y) == true {
			return SortKeyImp(RecentlyTier2, x, y)
		}
		if SortKeyEnable(RecentlyTier3, x, y) == true {
			return SortKeyImp(RecentlyTier3, x, y)
		}

		return false
	}})
	//!-customcall
	printTracks(tracks)
}

/*
//!+artistoutput
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go Ahead    Alicia Keys     As I Am            2007  4m36s
Go          Delilah         From the Roots Up  2012  3m38s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
Go          Moby            Moby               1992  3m37s
//!-artistoutput

//!+artistrevoutput
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go          Moby            Moby               1992  3m37s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
Go          Delilah         From the Roots Up  2012  3m38s
Go Ahead    Alicia Keys     As I Am            2007  4m36s
//!-artistrevoutput

//!+yearoutput
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go          Moby            Moby               1992  3m37s
Go Ahead    Alicia Keys     As I Am            2007  4m36s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
Go          Delilah         From the Roots Up  2012  3m38s
//!-yearoutput

//!+customout
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go          Moby            Moby               1992  3m37s
Go          Delilah         From the Roots Up  2012  3m38s
Go Ahead    Alicia Keys     As I Am            2007  4m36s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
//!-customout
*/

//!+customcode
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

//!-customcode

func init() {
	//!+ints
	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	sort.Ints(values)
	fmt.Println(values)                     // "[1 1 3 4]"
	fmt.Println(sort.IntsAreSorted(values)) // "true"
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	//!-ints
}
