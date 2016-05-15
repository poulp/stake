package core

import (
	"time"
	"fmt"
	"github.com/libgit2/git2go"
)

type StakeYear map[int]int

type StakeRepository struct {
	Start_date time.Time
	End_date time.Time

	/* Commits */
	Total_commits int
	Year_commits map[int]int
	Month_commits map[time.Month]int
}

func NewStakeRepository() *StakeRepository {
	return &StakeRepository{
		Year_commits: make(map[int]int),
		Month_commits: make(map[time.Month]int),
	}
}

func (sc *StakeRepository) Perform(commit *git.Commit){
	date := commit.Author().When
	year := date.Year()
	month := date.Month()

	if sc.Total_commits == 0 {
		sc.End_date = date
	}
	sc.Start_date = date

	/* Commits */
	sc.Total_commits++
	// By year
	_, exist := sc.Year_commits[year]
	if exist {
		sc.Year_commits[year]++
	} else {
		sc.Year_commits[year] = 1
	}
	// By month (global)
	_, exist = sc.Month_commits[month]
	if exist {
		sc.Month_commits[month]++
	} else {
		sc.Month_commits[month] = 1
	}
}

func (sc *StakeRepository) String() string {
	total_commits := fmt.Sprintf("Total commits : %d", sc.Total_commits)
	start_date := fmt.Sprintf("Start date : %v", sc.Start_date)
	end_date := fmt.Sprintf("End date : %v", sc.End_date)
	year_commits := fmt.Sprintf("Year commits: %v", sc.Year_commits)
	month_commits := fmt.Sprintf("Month commits: %v", sc.Month_commits)

	final := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", total_commits, start_date, end_date, year_commits, month_commits)
	return final
}