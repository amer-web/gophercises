package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

const newFileName = "Episode %03s%s"

var (
	filePattern = regexp.MustCompile(".*\\((\\d+).+\\)\\..+")
)

func main() {
	filenName := flag.String("file_name", "c-sample", "choose file")

	filepath.Walk(*filenName, wkFn)
	if err := filepath.Walk(*filenName, wkFn); err != nil {
		log.Fatal(err)
	}
	fmt.Println("rename successfully")
}

func wkFn(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	matched := filePattern.FindStringSubmatch(info.Name())
	if len(matched) == 0 {
		return nil
	}
	_, newFile := matched[0], fmt.Sprintf(newFileName, matched[1], filepath.Ext(path))

	return os.Rename(path, filepath.Join(filepath.Dir(path), newFile))

}
