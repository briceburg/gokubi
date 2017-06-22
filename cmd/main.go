package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"github.com/briceburg/gokubi"
)

// command
var action func()
var exitCode int = 1
var usageTemplate = `gokubi merges configuration and either renders a template or prints the result

Usage: {{.Cmd}} <options> [--] [paths...]

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

var data gokubi.Data

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

	if err := loadData(); err != nil {
		panic(err)
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
	data = make(gokubi.Data)
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

func die(msg string) {
	if msg == "" {
		msg = "bailing out"
	}
	fmt.Fprintf(os.Stderr, "\n⚡⚡\n⚡ %s \n⚡⚡\n", msg)
	os.Exit(exitCode)
}

///

func loadData() error {
	err := data.LoadPaths(flag.Args())
	if err == nil && isStdin() {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err == nil {
			return data.Decode(bytes, inputFormat)
		}
	}
	return err
}

func isStdin() bool {
	stat, _ := os.Stdin.Stat()
	mode := stat.Mode()
	return (mode & os.ModeNamedPipe) != 0
}

func output() {
	bytes, err := data.Encode(outputFormat)
	if err != nil {
		die(fmt.Sprintf("failed encoding: %s ", err.Error()))
	}
	os.Stdout.Write(bytes)
}

func render() {
	fmt.Println("RENDER!")
}
