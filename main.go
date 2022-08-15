package main

import (
	"io"
	"io/fs"
	stdlog "log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/urfave/cli/v2"
)

var log = stdlog.New(os.Stderr, "", 0)

func main() {
	var migRoot, upFile string
	var verbose bool

	app := cli.App{
		Name:  "mig-docker",
		Usage: "mig-docker <dir>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "migrations",
				Value:       "./db/migrations",
				Destination: &migRoot,
				Aliases:     []string{"m"},
			},
			&cli.StringFlag{
				Name:        "up",
				Value:       "up.sql",
				Destination: &upFile,
				Aliases:     []string{"u"},
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Destination: &verbose,
				Aliases:     []string{"v"},
			},
		},
		Action: func(c *cli.Context) error {
			outdir := c.Args().First()
			if outdir == "" {
				return cli.NewExitError("missing output directory", 1)
			}

			if err := os.MkdirAll(outdir, 0755); err != nil {
				log.Fatal(err)
			}

			fs.WalkDir(os.DirFS(migRoot), ".", func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					log.Println(err)
					return nil
				}

				if d.Name() != upFile {
					return nil
				}

				dir := filepath.Dir(path)
				if _, err := strconv.Atoi(dir); err != nil {
					return nil
				}

				in := filepath.Join(migRoot, path)
				out := filepath.Join(outdir, dir+".sql")
				if err := CopyFile(in, out); err != nil {
					return err
				}

				if verbose {
					log.Printf("%s -> %s", in, out)
				}
				return nil
			})
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
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
