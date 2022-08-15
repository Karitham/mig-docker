package main

import (
	"flag"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var migRoot = flag.String("migrations", "./db/migrations", "migrations directory")
var upFile = flag.String("up", "up.sql", "up file name")

func init() {
	flag.Parse()

	if len(os.Args) < 2 {
		log.Fatal("Usage: database-migrator <out dir>")
	}
}

func main() {
	outdir := os.Args[1]
	if err := os.MkdirAll(outdir, 0755); err != nil {
		log.Fatal(err)
	}

	fs.WalkDir(os.DirFS(*migRoot), ".", func(path string, d fs.DirEntry, _ error) error {
		if d.Name() != *upFile {
			return nil
		}

		dir := filepath.Dir(path)
		if _, err := strconv.Atoi(dir); err != nil {
			return nil
		}

		CopyFile(filepath.Join(*migRoot, path), filepath.Join(outdir, dir+".sql"))
		return nil
	})
}

// CopyFile copies a file from src to dst.
func CopyFile(src string, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}
