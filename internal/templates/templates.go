package templates

import (
	"bytes"
	"embed"
	"html/template"
	"io"
)

//go:embed html/*
var raw embed.FS

var t = func() *template.Template {
	t := template.New("")

	entries, _ := raw.ReadDir("html")
	for _, de := range entries {
		filename := de.Name()
		contents, _ := raw.ReadFile("html/" + filename)

		_, err := t.New(filename).Parse(string(contents))
		if err != nil {
			panic(err)
		}
	}

	return t
}()

func Render(w io.Writer, name string, params any) error {
	var temporalWriter bytes.Buffer
	err := t.ExecuteTemplate(&temporalWriter, name, params)
	if err != nil {
		return err
	}
	w.Write(temporalWriter.Bytes())
	return nil
}
