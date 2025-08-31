package main

import (
	"flag"
	"fmt"
	"localgitstats/internal/scan"
	"localgitstats/internal/stats"
)

type User struct {
	// User Id
	Uid string
	// Group Id
	Gid      string
	Username string
	Name     string
	HomeDir  string
}

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "specify folder")
	flag.StringVar(&email, "email", "your@gmail.com", "the email to scan")
	flag.Parse()

	fmt.Println(folder)

	if folder != "" {
		scan.Scan(folder)
	} else {
		stats.Stats(email)
	}
}
