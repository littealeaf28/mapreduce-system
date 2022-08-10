package main

//
// start the coordinator process, which is implemented
// in ../mr/coordinator.go
//
// go run main/mrcoordinator.go wc_inputfiles/pg*.txt
//

import (
	"mr_system/mr"
	"time"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: mrcoordinator inputfiles...\n")
		os.Exit(1)
	}

	m := mr.MakeCoordinator(os.Args[1:], 10)
	for m.Done() == false {
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second)
}
