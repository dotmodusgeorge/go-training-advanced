package main

import (
	"sync"
	"fmt"
	"flag"
	"os"
	"go-training-advanced/internal"
)

var wg sync.WaitGroup

func main() {
	output := flag.String("output", "./csv_output/", "Output of csvs")
	args := os.Args
	if (args[1] == "output") {
		args = args[2:]
	} else {
		args = args[1:]
	}
	fmt.Println(args)
	
	for _, arg := range args {
		wg.Add(1)
		go func() {
			internal.HandleJSON(arg, *output)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("asdf")
}
