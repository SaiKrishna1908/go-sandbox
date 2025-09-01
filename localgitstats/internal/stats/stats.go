package stats

import (
	"fmt"
	"localgitstats/internal/utils"
	"sort"
	"time"

	git "github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
)

const outOfRange = 99999
const daysInLastSixMonths = 182
const weeksInLastSixMonths = 26

type column []int

// stats calculates and prints the stats.
func Stats(email string) {
	commits := processRepositories(email)
	printCommitsStats(commits)
}

// getBeginningOfDay given a time.Time calculates the start time of that day
func getBeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return startOfDay
}

// countDaysSinceDate counts how many days passed since the passed `date`
func countDaysSinceDate(date time.Time) int {
	now := getBeginningOfDay(time.Now())
	startOfDay := getBeginningOfDay(date)
	duration := now.Sub(startOfDay)
	days := int(duration / (24 * time.Hour))
	if days > daysInLastSixMonths || days < 0 {
		return outOfRange
	}
	return days
}

// fillCommits given a repository found in `path`, gets the commits and
// puts them in the `commits` map, returning it when completed
func fillCommits(email string, path string, commits map[int]int) map[int]int {
	// instantiate a git repo object from path
	repo, err := git.PlainOpen(path)
	if err != nil {
		return commits
	}
	// get the HEAD reference
	ref, err := repo.Head()
	if err != nil {
		if err.Error() == "reference not found" {
			return commits
		}
		// log.DEBUG(err)
		return commits
	}
	// get the commits history starting from HEAD
	iterator, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return commits
	}
	// iterate the commits
	offset := calcOffset()
	err = iterator.ForEach(func(c *object.Commit) error {
		if c.Author.Email != email {
			return nil
		}
		days := countDaysSinceDate(c.Author.When)
		if days == outOfRange {
			return nil
		}
		daysAgo := daysInLastSixMonths + offset - days
		commits[daysAgo]++

		return nil
	})
	if err != nil {
		panic(err)
	}

	return commits
}

// processRepositories given an user email, returns the
// commits made in the last 6 months
func processRepositories(email string) map[int]int {
	filePath := utils.GetDotFilePath()
	repos := utils.ParseFileLinesToSlice(filePath)
	daysInMap := daysInLastSixMonths + 7

	commits := make(map[int]int, daysInMap)
	for i := 0; i <= daysInMap; i++ {
		commits[i] = 0
	}

	for _, path := range repos {
		commits = fillCommits(email, path, commits)
	}

	return commits
}

// calcOffset determines and returns the amount of days missing to fill
// the last row of the stats graph
func calcOffset() int {
	weekday := time.Now().Weekday()
	return int(weekday)
}

// printCell given a cell value prints it with a different format
// based on the value amount, and on the `today` flag.
func printCell(val int, today bool) {
	escape := "\033[0;37;30m"
	switch {
	case val == 0:
		// gray/white (empty day)
		escape = "\033[1;30;47m"
	case val > 0 && val <= 3:
		// very light green
		escape = "\033[1;30;102m"
	case val >= 4 && val <= 6:
		// light green
		escape = "\033[1;30;42m"
	case val >= 7 && val <= 9:
		// medium green (bold text)
		escape = "\033[1;37;42m"
	case val >= 10:
		// dark/strong shade
		escape = "\033[1;37;100m"
	}

	if today {
		escape = "\033[1;37;45m"
	}

	if val == 0 {
		fmt.Printf(escape + "  - " + "\033[0m")
		return
	}

	str := "  %d "
	switch {
	case val >= 10:
		str = " %d "
	case val >= 100:
		str = "%d "
	}

	fmt.Printf(escape+str+"\033[0m", val)
}

// printCommitsStats prints the commits stats
func printCommitsStats(commits map[int]int) {
	keys := sortMapIntoSlice(commits)
	cols := buildCols(keys, commits)
	printMonths()
	printCells(cols)
}

// sortMapIntoSlice returns a slice of indexes of a map, ordered
func sortMapIntoSlice(m map[int]int) []int {
	// order map
	// To store the keys in slice in sorted order
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return keys
}

// buildCols generates a map with rows and columns ready to be printed to screen
func buildCols(keys []int, commits map[int]int) map[int]column {
	cols := make(map[int]column)
	col := column{}
	lastWeek := -1

	for _, k := range keys {
		week := int(k / 7)
		dayinweek := k % 7

		if dayinweek == 0 {
			col = column{}
		}
		col = append(col, commits[k])
		lastWeek = week

		if dayinweek == 6 {
			cols[week] = col
		}
	}

	if len(col) > 0 {
		cols[lastWeek] = col
	}

	return cols
}

// printCells prints the cells of the graph
func printCells(cols map[int]column) {
	for j := 0; j <= 6; j++ {
		printDayCol(j)
		for i := 0; i <= weeksInLastSixMonths+1; i++ {
			v := 0
			today := false
			if col, ok := cols[i]; ok {
				if len(col) > j {
					v = col[j]
				}
			}
			if i == weeksInLastSixMonths && j == calcOffset() {
				today = true
			}
			printCell(v, today)
		}
		fmt.Printf("\n")
	}
}

// printMonths prints the month names in the first line, determining when the month
// changed between switching weeks
func printMonths() {
	week := getBeginningOfDay(time.Now()).Add(-(daysInLastSixMonths * time.Hour * 24))
	month := week.Month()
	fmt.Printf("         ")
	for {
		if week.Month() != month {
			fmt.Printf("%s ", week.Month().String()[:3])
			month = week.Month()
		} else {
			fmt.Printf("    ")
		}

		week = week.Add(7 * time.Hour * 24)
		if week.After(time.Now()) {
			break
		}
	}
	fmt.Printf("\n")
}

// printDayCol given the day number (0 is Sunday) prints the day name,
// alternating the rows (prints just 2,4,6)
func printDayCol(day int) {
	out := "     "
	switch day {
	case 1:
		out = " Mon "
	case 3:
		out = " Wed "
	case 5:
		out = " Fri "
	}

	fmt.Printf(out)
}
