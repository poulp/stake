package core

import (
	"time"
	"fmt"
)

type StakeInfo struct {
	Version string
	BuildDate time.Time
	GenerationTime time.Duration
}

func (si *StakeInfo) SetInfo(version string, start_generation time.Time){
	si.Version = version
	si.GenerationTime = time.Since(start_generation)
	si.BuildDate = time.Now()

	fmt.Printf("%v", si)
}

func (si *StakeInfo) String()string {
	title := fmt.Sprintf("***** Stake *****")
	version := fmt.Sprintf("Version : %s", si.Version)
	generation_time := fmt.Sprintf("Generation time : %v seconds", si.GenerationTime.Seconds())
	build_date := fmt.Sprintf("Build date : %v", si.BuildDate)

	return fmt.Sprintf("%s\n%s\n%s\n%s\n", title, version, generation_time, build_date)
}