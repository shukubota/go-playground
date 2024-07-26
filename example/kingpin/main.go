package main

import (
	"fmt"
	"github.com/alecthomas/kingpin/v2"
)

var (
	cmd = kingpin.Flag("cmd", "command name").Required().Short('c').String()
	//arg = kingpin.Arg("arg", "command arguments").Strings()

	count = kingpin.Arg("count", "count").Required().Int()
	name  = kingpin.Arg("name", "name").String()
	date  = kingpin.Arg("date", "date").String()
)

func main() {
	kingpin.Parse()
	fmt.Println("cmdflag:", *cmd)
	fmt.Println("count:", *count)
	fmt.Println("name:", *name)
	fmt.Println("date:", *date)
}
