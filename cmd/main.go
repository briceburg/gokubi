package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/briceburg/gokubi"
)

// command
var action func()
var exitCode int = 1
var usageTemplate = `gokubi merges configuration and either renders a template or prints the result

Usage: {{.Cmd}} [options] [--] [paths...]

Examples:
  # merge configuration from xml, yml, and .properties files and output as JSON
  gokubi .properties conf.d/*.xml conf.d/*.yml -o json

  # read JSON from stdin and use it to render template.j2
  curl https://json.api/ | gokubi --render-file template.j2

  # read configuration variables into a bash shell
  eval $(gokubi conf.d/ -o bash)

Options:
`

// command flags
var inputFormat string
var outputFormat string
var templateStr string
var templatePath string

///

func main() {

	flag.StringVar(&inputFormat, "i", "json", "Input `format` (used when configuration is passed as stdin). Supported Values: "+strings.Join(gokubi.InputFormats, ", "))
	flag.StringVar(&outputFormat, "o", "", "Output `format`. Supported Values: "+strings.Join(gokubi.OutputFormats, ", "))
	flag.StringVar(&templateStr, "template", "", "Template string to render. A value of '-' reads template from stdin.")
	flag.StringVar(&templatePath, "template-path", "", "Path to template file to render.")
	flag.Parse()
	flag.Visit(setAction)

	if action == nil {
		flag.Usage()
		die("please specify an output format or template to render")
	}
	action()
}

// we use init to define [flag] command usage
func init() {
	ctx := struct {
		Cmd string
	}{
		os.Args[0],
	}
	t := template.Must(template.New("usage").Parse(usageTemplate))
	flag.Usage = func() {
		t.Execute(os.Stderr, ctx)
		flag.PrintDefaults()
	}
}

// determines command path based on a passed flag. last "command" flag wins.
func setAction(flag *flag.Flag) {
	switch flag.Name {
	case "o":
		action = output
	case "template", "template-path":
		action = render
	}
}

///

func die(msg string) {
	if msg == "" {
		msg = "bailing out"
	}
	fmt.Fprintf(os.Stderr, "\n⚡⚡\n⚡ %s \n⚡⚡\n", msg)
	os.Exit(exitCode)
}

func output() {
	fmt.Println("OUT!")
	/*
		fmt.Println(*inputFormat)
		fmt.Println(*outputFormat)
		fmt.Println(*template)
		fmt.Println(*templatePath)
		/*

			data := make(gokubi.Data)

			if err := gokubi.FileReader("formats/json/fixtures/music.json", &data); err != nil {
				panic(err)
			}
			if err := gokubi.FileReader("formats/yaml/fixtures/music.yml", &data); err != nil {
				panic(err)
			}
			fmt.Println(data.String())

			out, _ := data.EncodeYAML()
			fmt.Println(string(out))

	*/
	//fmt.Println(data.EncodeBash())

	/*
		if err := readers.DirectoryReader("fixtures", &data); err != nil {
			panic(err)
		}
		fmt.Println(data.String())
		fmt.Println("%+v", data)
	*/
}

func render() {
	fmt.Println("RENDER!")
}
