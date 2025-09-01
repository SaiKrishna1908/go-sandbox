package main

import (
	"flag"
	"fmt"
	"localgitstats/internal/scan"
	"localgitstats/internal/stats"
	"localgitstats/internal/utils"
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
	var clear string

	flag.StringVar(&folder, "add", "", "specify folder")
	flag.StringVar(&email, "email", "", "the email to scan")
	flag.StringVar(&clear, "clear", "false", "clear repositories to track")
	flag.Parse()

	fmt.Println(folder)

	if clear != "false" {
		isFileTruncated := utils.ClearDotFile()
		if isFileTruncated {
			fmt.Println("Cleared tracked repositories")
		}
		return
	}

	if folder == "" && email == "" {
		fmt.Printf("please provide --add or --email")
		return
	}

	if folder != "" {
		scan.Scan(folder)
	} else if email != "" {
		stats.Stats(email)
	}
}
