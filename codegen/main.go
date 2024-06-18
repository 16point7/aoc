package main

import (
	"embed"
	"flag"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

//go:embed sources
var sources embed.FS

func main() {
	now := time.Now()

	var year int
	flag.IntVar(&year, "year", now.Year(), "year to generate")
	var day int
	flag.IntVar(&day, "day", now.Day(), "day to generate")

	flag.Parse()

	_, caller, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to determine output directory")
	}
	dstDir := filepath.Join(filepath.Dir(caller), "..", strconv.Itoa(year), strconv.Itoa(day))

	err := os.MkdirAll(dstDir, 0777)
	if err != nil {
		log.Fatal(err)
	}

	fs.WalkDir(sources, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		dstPath := filepath.Join(dstDir, d.Name())

		srcFile, err := sources.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}

		return nil
	})

}
