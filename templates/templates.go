package templates

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

// Templates is the map of available templates by name
var Templates = map[string]*Template{}

// Add adds an available template
func Add(template *Template) {
	Templates[template.Name] = template
}

type Template struct {
	Name   string
	Inputs []*UserInput
	Files  TemplateFiles
}

type TemplateFiles func() []*os.File

func (t *Template) TemplateFiles(inputs map[string]string) error {
	engine := template.New("infrastructure")
	for _, file := range t.Files() {
		fmt.Println("Name: " + file.Name())
		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("Error reading: " + file.Name())
			os.Exit(1)
		}
		dest, err := os.Create(file.Name())
		if err != nil {
			fmt.Println("Error creating output file: " + file.Name())
			os.Exit(1)
		}
		tmpl, err := engine.Parse(string(data))
		if err != nil {
			fmt.Println("Error parsing data from: " + file.Name())
			os.Exit(1)
		}
		err = tmpl.Execute(dest, &Input{
			Variables: inputs,
		})
		if err != nil {
			fmt.Println("Error executing templating file: " + file.Name())
			os.Exit(1)
		}
	}
	return nil
}

type UserInput struct {
	Name        string
	Default     string
	Description string
}

type Input struct {
	Variables map[string]string
}
