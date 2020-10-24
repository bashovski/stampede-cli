package template

import (
	"bytes"
	"fmt"
	"github.com/dustin/go-humanize/english"
	"io/ioutil"
	"log"
	"os"
)

// ReplaceFileTemplates represents a method which replaces existing templates with new base modules
func ReplaceFileTemplates(templateNames []string, baseName string) {
	for i := 0; i < len(templateNames); i++ {
		ReplaceFileTemplate(templateNames[i], baseName)
	}
}

// ReplaceFileTemplate represents a method which creates a new base TypeScript module for Stampede-based project while replacing base template content
func ReplaceFileTemplate(modelName string, baseName string) {
	path, _ := os.Getwd()
	srcPath := fmt.Sprintf("%s/assets/module_templates/%sTemplate.ts", path, baseName)
	moduleCategory := english.PluralWord(2, baseName, "")
	destinationPath := fmt.Sprintf("%s/%s/%s%s.ts", path, moduleCategory, modelName, baseName)

	err := replace(modelName, srcPath, destinationPath)
	if err != nil {
		log.Fatal("Unable to replace file template - failed at copying the template", err)
	}
	fmt.Println("Successfully created:", fmt.Sprintf("%s%s", modelName, baseName))
}

func replace(modelName string, srcPath string, destPath string) error {
	src, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return err
	}

	src = bytes.Replace(src, []byte("__replace_me__"), []byte(modelName), -1)
	if err = ioutil.WriteFile(destPath, src, 0666); err != nil {
		return err
	}
	return nil
}