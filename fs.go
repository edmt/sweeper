package main

import (
	l4g "code.google.com/p/log4go"
	"io"
	"os"
	"path/filepath"
)

func ListFiles(globPattern string) (matches []string, err error) {
	return filepath.Glob(globPattern)
}

func GetGlobPatternList(baseDir string, year, month, day string) (output []string) {
	rfcList, _ := getRFCList(baseDir)

	output = make([]string, len(rfcList))
	for i, value := range rfcList {
		output[i] = filepath.Join(value, "CFDs_Expedidos", year, month, day, "*.xml")
	}
	return
}

func getRFCList(baseDir string) (matches []string, err error) {
	return filepath.Glob(filepath.Join(baseDir, "*"))
}

func Copy(source, destination string) (err error) {
	in, err := os.Open(source)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(destination)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func Mkdir(path string) {
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		l4g.Error("Error al crear la estructura de directorios %s. Error:%s", path, err.Error())
	}
}

func BackUp(source, destination string) {
	Mkdir(destination)
	Copy(source, destination)
}
