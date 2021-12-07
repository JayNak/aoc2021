package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/jaynak/aoc2021/pkg/aoc"
)

func main() {

	// Pull in the function map
	fns := aoc.GetFunctions()
	dataPath := "."

	// Check for passed in arguments
	if len(os.Args) > 1 {
		dataPath = os.Args[1]
	} else {
		// fall through to the default data folder
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		dataPath = exPath + "/../data"
	}

	// Validate that the folder exists!
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		panic(err)
	}

	// Loop through the files in the data folder and call functions
	r := regexp.MustCompile("^([0-9]+).+$")

	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		panic(err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(writer, "file name\tpart 1\tpart 2\telapsed time")

	for _, f := range files {

		m := r.FindAllStringSubmatch(f.Name(), -1)

		// First character of the filename indicates the day
		if len(m) == 0 {
			continue
		}

		n, err := strconv.Atoi(m[0][1])
		if err != nil {
			panic(err)
		}

		// fns is a slice of the functions
		if n < len(fns) {
			start := time.Now()
			a, b := fns[n](dataPath + "/" + f.Name())
			elapsed := time.Since(start)
			fmt.Fprintf(writer, "%v\t%v\t%v\t%s\n", f.Name(), a, b, elapsed)

			// fmt.Printf("%v: %v, %v in %s\n", f.Name(), a, b, elapsed)
		}
	}

	writer.Flush()
}
