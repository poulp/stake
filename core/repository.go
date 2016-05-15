package core

import (
	"time"
	"fmt"
	"github.com/libgit2/git2go"
)

type StakeRepository struct {
	Start_date time.Time
	End_date time.Time
	Total_commits int
}

func (sc *StakeRepository) Perform(commit *git.Commit){
	if sc.Total_commits == 0 {
		sc.End_date = commit.Author().When
	}
	sc.Start_date = commit.Author().When
	sc.Total_commits++
}

func (sc *StakeRepository) String() string {
	total_commits := fmt.Sprintf("Total commits : %d", sc.Total_commits)
	start_date := fmt.Sprintf("Start date : %v", sc.Start_date)
	end_date := fmt.Sprintf("end date : %v", sc.End_date)

	final := fmt.Sprintf("%s\n%s\n%s\n", total_commits, start_date, end_date)
	return final
}