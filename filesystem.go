package gokubi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// mapping of extension => gokubi method
var ExtMethodMap = map[string]string{
	".json": "DecodeJSON",
	".yaml": "DecodeYML",
	".yml":  "DecodeYML",
	".html": "DecodeXML",
	".xml":  "DecodeXML",
	".hcl":  "DecodeHCL"}

func FileReader(p string, d *Data) error {
	ext := filepath.Ext(p)
	method, ok := ExtMethodMap[ext]
	if ok {
		body, err := ioutil.ReadFile(p)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gokubi/filesystem.FileReader: %v", err)
			return err
		}
		return d.Decode(method, body)
	}

	return fmt.Errorf("gokubi/filesystem.FileReader: unsupported path: %v", p)
}

// reads supported files in a directory in lexical order
func DirectoryReader(p string, d *Data) error {
	return lexicalLoadFiles(p, false, d)
}

// reads supported files in a directory and its subdirectories in lexical order
func RecursiveDirectoryReader(p string, d *Data) error {
	return lexicalLoadFiles(p, true, d)
}

// reads supported files in a directory concurrently (disregards order)
func DirectoryReaderFast(p string, d *Data) error {
	return errors.New("DirectoryReaderFast not implemented")
}

// reads supported in a directory and its subdirectories concurrently (disregards order)
func RecursiveDirectoryReaderFast(p string, d *Data) error {
	return errors.New("RecursiveDirectoryReaderFast not implemented")
}

func lexicalLoadFiles(dir string, recurse bool, d *Data) error {
	files, err := lexicalWalk(dir, recurse, onlyConfigFiles)
	if err != nil {
		return fmt.Errorf("gokubi/filesystem.DirectoryReader: %v", err)
	}

	for _, p := range files {
		if err := FileReader(p, d); err != nil {
			// @TODO continue execution on individual file errs?
			return err
		}
	}

	return nil
}

func lexicalWalk(pwd string, recurse bool, filterFn func(string, os.FileInfo) bool) ([]string, error) {
	var list []string
	stat, err := os.Stat(pwd)
	if err != nil {
		return list, err
	}
	if !stat.IsDir() {
		return list, fmt.Errorf("stat: %s: must be a directory", pwd)
	}

	walkFn := func(p string, stat os.FileInfo, err error) error {
		if err != nil {
			fmt.Fprintf(os.Stderr, "gokubi/filesystem.walkFn: failed on %s: %v", p, err)
			return err
		}

		if stat.IsDir() {
			if recurse || p == pwd {
				return nil
			}
			// halt recursion
			return filepath.SkipDir
		}

		if filterFn(p, stat) {
			list = append(list, p)
		}

		return nil
	}

	return list, filepath.Walk(pwd, walkFn)
}

func onlyConfigFiles(p string, stat os.FileInfo) bool {
	_, ok := ExtMethodMap[filepath.Ext(stat.Name())]
	return ok
}
