package walker

import (
	"fmt"
	"image-compressor/src/constants"
	"io/fs"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

func getPaths(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func filterByExtend(paths []string, extends []string) []string {
	var result []string
	regString := strings.Join(extends, "|")
	for i := 0; i < len(paths); i++ {
		match, _ := regexp.MatchString(regString, paths[i])
		if match {
			result = append(result, paths[i])
		}
	}
	return result
}

func GetImagesPath(path string, extends []string) {
	files, err := getPaths(path)
	if err != nil {
		log.Fatal(err)
	}
	images := filterByExtend(files, []string{".jpg"})
	fmt.Println(images)
}
