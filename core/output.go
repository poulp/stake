package core

import (
	"encoding/json"
	"os"
	"fmt"
)

func GetStakeOutput(arg string)StakeOutput{
	switch arg {
	case "json":
		return StakeJSONOutput{filename: "stake.json"}
	case "html":
		return StakeHTMLOutput{dir_name: "html"}
	default:
		return StakeJSONOutput{filename: "stake.json"}
	}
}

/* Interface */
type StakeOutput interface {
	render(*StakeRepository, *StakeAuthors)
	displayBefore()
}

/***********
 * JSON
 **********/
type StakeJSONOutput struct {
	filename string
}

func (sjo StakeJSONOutput) render(r *StakeRepository, a *StakeAuthors){
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

func (sjo StakeJSONOutput) displayBefore(){
	fmt.Println("Output format : JSON")
	fmt.Println("Filename : %s", sjo.filename)
}

/***********
 * HTML
 **********/
type StakeHTMLOutput struct {
	dir_name string
}

func (sho StakeHTMLOutput) render(r *StakeRepository, a *StakeAuthors){
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

func (sjo StakeHTMLOutput) displayBefore(){
	fmt.Println("Output format : HTML")
	fmt.Printf("Directory : %s\n", sjo.dir_name)
}