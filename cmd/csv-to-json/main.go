package main

import (
	"sync"
	"fmt"
	"os"
	"flag"
	"go-training-advanced/internal"
) 

var wg sync.WaitGroup

func main() {
	output := flag.String("output", "./json_output/", "The output for the json file")

	args := os.Args
	if (args[1] == "output") {
		args = args[2:]
	} else {
		args = args[1:]
	}
	for _, arg := range args {
		wg.Add(1)
		go func() {
			internal.HandleCsv(arg, *output)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(args)
}
