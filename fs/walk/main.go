package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	archive string    // archive directory
	ext     string    // extension to filter out
	size    int64     // min file size
	list    bool      // list files
	del     bool      // delete files
	wLog    io.Writer // log destination writer
}

func run(root string, out io.Writer, cfg config) error {
	delLogger := log.New(cfg.wLog, "DELETED FILE: ", log.LstdFlags)

	return filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if filterOut(path, cfg.ext, cfg.size, info) {
				return nil
			}

			// if list was set, don't do anything else
			if cfg.list {
				return listFile(path, out)
			}

			// archive files and continue if successful
			if cfg.archive != "" {
				if err := archiveFile(cfg.archive, root, path); err != nil {
					return err
				}
			}

			// delete files
			if cfg.del {
				return delFile(path, delLogger)
			}

			// if nothing else was set, list is the default option
			return listFile(path, out)
		})
}

func main() {
	// parse flags
	archive := flag.String("archive", "", "archive directory")
	del := flag.Bool("del", false, "delete files")
	ext := flag.String("ext", "", "file extension to filter out")
	list := flag.Bool("list", false, "list files only")
	logFile := flag.String("log", "", "log deletes to this file")
	root := flag.String("root", ".", "starting directory")
	size := flag.Int64("size", 0, "minimum file size")
	flag.Parse()

	var (
		f   = os.Stdout
		err error
	)

	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}

	c := config{
		archive: *archive,
		del:     *del,
		ext:     *ext,
		size:    *size,
		list:    *list,
		wLog:    f,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
