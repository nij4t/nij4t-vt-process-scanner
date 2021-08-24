package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	vt "github.com/VirusTotal/vt-go"
)

var apikey = flag.String("apikey", "", "VirusTotal API key")
var sha256 = flag.String("sha256", "", "SHA-256 of some file")
var out = flag.String("out", "", "Output file")

func main() {

	flag.Parse()

	if *apikey == "" || *sha256 == "" {
		fmt.Println("Must pass both the --apikey and --sha256 arguments.")
		os.Exit(0)
	}

	client := vt.NewClient(*apikey)

	file, err := client.GetObject(vt.URL("files/%s", *sha256))
	if err != nil {
		log.Fatal(err)
	}

	ls, err := file.GetTime("last_submission_date")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File %s was submitted for the last time on %v\n", file.ID(), ls)

	data, err := file.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(data))

	if *out == "" {
		os.Exit(0)
	}

	f, err := os.OpenFile(*out, os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(data)

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
}
