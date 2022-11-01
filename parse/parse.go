package parse

import (
	"log"
	"os"

	"golang.org/x/mod/modfile"
)

const (
	modFileName = "go.mod"
)

var (
	FilePath string
	ModFile  *modfile.File
)

func MustParse() {
	fileData, err := os.ReadFile(FilePath + "/" + modFileName)
	if err != nil {
		log.Fatalln(err)
	}
	modFile, err := modfile.Parse(modFileName, fileData, nil)
	if err != nil {
		log.Fatalln(err)
	}
	ModFile = modFile
}
