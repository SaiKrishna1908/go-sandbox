package main

import (
	"bufio"
	"flag"
	_ "flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"slices"
	"strings"
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

// recursively se
func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")

	f, err := os.Open(folder)

	if err != nil {
		log.Fatal(err)
	}

	files, err := f.ReadDir(-1)

	if err != nil {
		log.Fatal(err)
	}

	var path string

	foldersToIgnore := []string{"vendor", "node_modules", "venv"}

	for _, file := range files {
		// log.Println(file.Name())
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}
			if slices.Contains(foldersToIgnore, file.Name()) {
				continue
			}
			folders = scanGitFolders(folders, path)
		}
	}

	return folders
}

func scan(folder string) {
	fmt.Printf("Found Folders:\n\n")
	repositories := recursiveScanFolder(folder)
	filePath := getDotFilePath()
	addNewSliceElementsToFile(filePath, repositories)
	fmt.Printf("\n\nSuccessfully Added\n\n")
}

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func getDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "/.localgitstats"
	return dotFile
}

func addNewSliceElementsToFile(filePath string, newRepos []string) {
	existingRepos := parseFileLinesToSlice(filePath)
	repos := joinSlices(newRepos, existingRepos)
	dumpStringSliceToFile(repos, filePath)
}

func dumpStringSliceToFile(repos []string, filePath string) {
	content := strings.Join(repos, "\n")
	os.WriteFile(filePath, []byte(content), 0755)
}

func joinSlices(new []string, existing []string) []string {
	for _, i := range new {
		if !slices.Contains(existing, i) {
			existing = append(existing, i)
		}
	}
	return existing
}

func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

func parseFileLinesToSlice(filePath string) []string {
	f := openFile(filePath)
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err := scanner.Err()
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	return lines
}

func openFile(filePath string) *os.File {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(filePath)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	return f
}

func stats(email string) {
	print("stats for " + email)
}

func main() {
	var folder string
	var email string

	pd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&folder, "add", pd, "specify folder")
	flag.StringVar(&email, "email", "your@gmail.com", "the email to scan")

	if folder != "" {
		scan(folder)
	}

	stats(email)
}
