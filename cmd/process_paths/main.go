package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nij4t/vt-process-scanner/pkg/process"
)

var outType = flag.String("out-type", "path", "Output type (path | hash)")

func main() {

	flag.Parse()

	proc, err := process.GetProcessPaths()

	if err != nil {
		log.Fatal(err)
	}

	if *outType == "hash" {
		for _, v := range proc {
			hash := sha256.Sum256([]byte(v))
			encoder := hex.NewEncoder(os.Stdout)
			encoder.Write(hash[:])
			fmt.Println()
		}
		os.Exit(0)
	}

	fmt.Println(strings.Join(proc, "\n"))
}
