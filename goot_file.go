package main

import (
	"fmt"
	"io"
	"os"
)

var minimalPackage = []string{"doc", "main", "test"}

type gootFile struct {
	anchor, packName, fileName      string
	template, okMessage, subcommand string
	output                          io.Writer
}

// Checks if a given file is part of the mininmal
// version of the package.
func (gf gootFile) isMinimalFile() bool {
	for _, value := range minimalPackage {
		if gf.anchor == value {
			return true
		}
	}
	return false
}

// Creates the based on the construction
// passed on the gootstrap file
func (gf gootFile) createFile() error {

	// Creates the file and defer its closing
	fileCreate, err := os.Create(gf.fileName)
	if err != nil {
		return err
	}
	defer fileCreate.Close()

	// Test is the current file has any template
	// attached to it and,then, writes the output
	// to os.Stdout.
	if gf.template != "" {
		fileCreate.WriteString(gf.template)
	}
	fmt.Fprintln(gf.output, gf.okMessage)

	return nil
}

// Performs creation based on the subcommand passed
func (gf gootFile) performCreation() error {

	// Checks if the subcommand is either ""
	// or --minimal in order to perform the
	// creation on the correct files.
	switch gf.subcommand {
	case "":
		err := gf.createFile()
		if err != nil {
			return err
		}
	case "--minimal":
		if gf.isMinimalFile() {
			err := gf.createFile()
			if err != nil {
				return err
			}
		}
	}
	return nil

}
