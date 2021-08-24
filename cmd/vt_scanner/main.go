package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	vt "github.com/VirusTotal/vt-go"
)

var apikey = flag.String("apikey", "", "VirusTotal API key")
var sha256 = flag.String("sha256", "", "SHA-256 of some file")
var hashList = flag.String("hashlist", "", "File containing hashes")
var out = flag.String("out", "", "Output file")

func main() {

	hashes := []string{}

	flag.Parse()

	if *apikey == "" {
		fmt.Println("Must pass --apikey argument.")
		os.Exit(1)
	}

	if *hashList != "" {
		fileWithHashes, err := os.ReadFile(*hashList)
		if err != nil {
			log.Fatal(err)
		}
		hashes = strings.Split(string(fileWithHashes), "\n")
	}

	if *sha256 != "" {
		hashes = append(hashes, *sha256)
	}

	client := vt.NewClient(*apikey)

	reports := []*vt.Object{}
	distributors := []string{}

	for _, hash := range hashes {
		report, err := getReport(client, hash)
		if err != nil {
			log.Print(err)
			continue
		}
		reports = append(reports, report)
		ls, err := report.GetTime("last_submission_date")
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("File %s was submitted for the last time on %v\n", report.ID(), ls)

		distributor, err := getDistributors(report)
		if err != nil {
			log.Print(err)
			continue
		}
		if distributor, ok := distributor.([]string); ok {
			distributors = append(distributors, distributor[0])
		}
	}

	output := os.Stdout

	if *out != "" {
		output, err := os.OpenFile(*out, os.O_WRONLY|os.O_CREATE, 0644)
		defer output.Close()

		if err != nil {
			log.Fatal(err)
		}
	}

	// _, err = f.Write(data)
	_, err := output.Write([]byte(strings.Join(distributors, "\n")))

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
}

func getReport(client *vt.Client, hash string) (*vt.Object, error) {
	return client.GetObject(vt.URL("files/%s", *sha256))
}

func getDistributors(report *vt.Object) (interface{}, error) {
	return report.Get("distributors")
}
