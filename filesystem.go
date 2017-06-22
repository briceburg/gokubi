package gokubi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// mapping of extension => gokubi format
var FormatExtensionsMap = map[string][]string{
	"bash": {".vars", ".env", ".bash", ".sh"},
	"json": {".json"},
	"yaml": {".yaml", ".yml"},
	"xml":  {".xml", ".html"},
	"hcl":  {".hcl"},
}

func (d *Data) LoadPaths(paths []string) error {
	for _, p := range paths {
		if err := PathReader(p, d); err != nil {
			return err
		}
	}
	return nil
}

// given a path, attempts to find a suitable format using the filename extension
func FormatFromPath(p string) (string, error) {
	ext := strings.ToLower(filepath.Ext(p))
	for f, list := range FormatExtensionsMap {
		for _, v := range list {
			if ext == v {
				return f, nil
			}
		}
	}
	return "", fmt.Errorf("gokubi/filesystem.MethodFromPath: failed on %v", ext)
}

func PathReader(p string, d *Data) error {
	stat, err := os.Stat(p)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return DirectoryReader(p, d)
	}
	return FileReader(p, d)
}

func FileReader(p string, d *Data) error {
	f, err := FormatFromPath(p)
	if err != nil {
		return fmt.Errorf("gokubi/filesystem.FileReader: unsupported path: %v", p)
	}
	bytes, err := ioutil.ReadFile(p)
	if err != nil {
		return fmt.Errorf("gokubi/filesystem.FileReader: %v", err)
	}

	return d.Decode(bytes, f)
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
	_, err := FormatFromPath(stat.Name())
	return err == nil
}
