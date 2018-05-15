package main

import (
	"bytes"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"text/template"
)

var (
	valuesFile = flag.String("valuesFile", "", "file with values for Dockerfile template")
)

const outputFile = "Dockerfile"

// Values provided in the values file
type Values map[string]string

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalln("Dockerfile template expected")
	}

	if valuesFile == nil || len(*valuesFile) == 0 {
		log.Fatalln("A values file expected")
	}

	yamlValues, err := ioutil.ReadFile(*valuesFile)
	if err != nil {
		log.Fatalf("Failed to read the values file. Error: %v\n", err)
	}

	templateFile := args[0]
	if len(templateFile) == 0 {
		log.Fatalln("Dockerfile template is empty")
	}

	buf, err := ioutil.ReadFile(templateFile)
	if err != nil {
		log.Fatalln(err)
	}

	tmpl := template.Must(template.New("generated").Parse(string(buf[:])))
	if tmpl == nil {
		log.Fatalln("Failed to parse the template")
	}

	var values Values
	err = yaml.Unmarshal(yamlValues, &values)
	if err != nil {
		log.Fatalf("Failed to unmarshal the yaml values: Error: %v\n", err)
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, values)
	if err != nil {
		log.Fatalf("Failed to render the template. Error: %v\n", err)
	}

	if err := ioutil.WriteFile(outputFile, output.Bytes(), 0644); err != nil {
		log.Fatalf("Failed to safe the generated file. Error: %v", err)
	}
}
