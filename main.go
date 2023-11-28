package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aquasecurity/trivy/pkg/types"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	publishedBefore := flag.String("published-before", "", "take vulnerabilities published before the specified timestamp (ex. 2019-11-04)")
	publishedAfter := flag.String("published-after", "", "take vulnerabilities published after the specified timestamp (ex. 2019-11-04)")
	flag.Parse()

	var before, after time.Time
	var err error
	if *publishedBefore != "" {
		before, err = time.Parse("2006-01-02", *publishedBefore)
		if err != nil {
			return err
		}
	}
	if *publishedAfter != "" {
		after, err = time.Parse("2006-01-02", *publishedAfter)
		if err != nil {
			return err
		}
	}

	var report types.Report
	if err := json.NewDecoder(os.Stdin).Decode(&report); err != nil {
		return err
	}

	var count int
	for _, result := range report.Results {
		for _, vuln := range result.Vulnerabilities {
			if (!before.IsZero() || !after.IsZero()) && vuln.PublishedDate == nil {
				continue
			}
			if (!before.IsZero() && vuln.PublishedDate.After(before)) ||
				(!after.IsZero() && vuln.PublishedDate.Before(after)) {
				continue
			}
			count += 1
		}
	}
	fmt.Printf("Number of vulnerabilities: %d\n", count)
	return nil
}
