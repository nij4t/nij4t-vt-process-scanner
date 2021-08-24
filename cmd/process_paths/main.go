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
			f, err := os.ReadFile(v)
			if err != nil {
				log.Print(err)
				continue
			}

			hash := sha256.Sum256(f)
			encoder := hex.NewEncoder(os.Stdout)
			encoder.Write(hash[:])
			fmt.Println()
		}
		os.Exit(0)
	}

	fmt.Println(strings.Join(proc, "\n"))
}
