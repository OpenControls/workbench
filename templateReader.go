package main

import (
	"github.com/markbates/pkger"
	"text/template"
)

func CompileTemplate(path string) (*template.Template, error) {
	tpl := template.New(path)
	// Since Walk receives a dynamic value, pkger won't be able to find the
	// actual directory to package from the next line, which is why we used
	// pkger.Include() in main().
	file, err := pkger.Open(path)
	info, err := file.Stat()
	b := make([]byte, info.Size())

	_, _ = file.Read(b)
	defer file.Close()
	fileAsStr := string(b)
	tpl.Parse(fileAsStr)

	return tpl, err
}
