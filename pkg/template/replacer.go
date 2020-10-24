package template

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ReplaceFileTemplate represents a method which creates a new base TypeScript module for Stampede-based project while replacing base template content
func ReplaceFileTemplate(modelName string, baseName string) {
	path, _ := os.Getwd()
	srcPath := fmt.Sprintf("%s/assets/module_templates/%sTemplate.ts", path, baseName)
	destinationPath := fmt.Sprintf("%s/controllers/%sController.ts", path, modelName)

	err := replace(modelName, srcPath, destinationPath)
	if err != nil {
		log.Fatal("Unable to replace file template - failed at copying the template", err)
	}
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