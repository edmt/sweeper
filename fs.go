package main

import (
	l4g "code.google.com/p/log4go"
	"os"
	"path/filepath"
)

func ListFiles(globPattern string) (matches []string, err error) {
	return filepath.Glob(globPattern)
}

func GetGlobPatternList(baseDir string) (output []string) {
	rfcList, _ := getRFCList(baseDir)

	for _, value := range rfcList {
		l4g.Debug("Probando directorio: %s",
			filepath.Join(value, "CFDS_Recibidos"))
		dirExists, _ := exists(filepath.Join(value, "CFDS_Recibidos"))
		if dirExists {
			output = append(output,
				filepath.Join(value, "CFDS_Recibidos", "*", "*", "*", "*.xml"))
		}
	}
	return
}

func getRFCList(baseDir string) (matches []string, err error) {
	return filepath.Glob(filepath.Join(baseDir, "*"))
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
