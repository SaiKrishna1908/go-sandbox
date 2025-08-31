package scan

import (
	"fmt"
	"localgitstats/internal/utils"
	"log"
	"os"
	"slices"
	"strings"
)

func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")

	f, err := os.Open(folder)

	if err != nil {
		if os.IsPermission(err) {
			return nil
		}
	}

	files, err := f.ReadDir(-1)

	if err != nil {
		log.Fatal(err)
	}

	var path string

	foldersToIgnore := []string{"vendor", "node_modules", "venv"}

	for _, file := range files {
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

func Scan(folder string) {
	fmt.Printf("Found Folders:\n\n")
	repositories := recursiveScanFolder(folder)
	filePath := utils.GetDotFilePath()
	addNewSliceElementsToFile(filePath, repositories)
	fmt.Printf("\n\nSuccessfully Added\n\n")
}

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func addNewSliceElementsToFile(filePath string, newRepos []string) {
	existingRepos := utils.ParseFileLinesToSlice(filePath)
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
