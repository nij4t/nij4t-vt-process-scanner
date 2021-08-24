package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/nij4t/vt-process-scanner/pkg/process"
)

func main() {
	// taskList, err := getProcessList()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// proc, err := getProcessPath(taskList[0])
	proc, err := process.GetProcessPaths()

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(strings.Join(taskList, "\n"))
	fmt.Println(strings.Join(proc, "\n"))
}
