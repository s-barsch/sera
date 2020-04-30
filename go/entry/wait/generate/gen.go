// Code should be run from parent directory with
//
// `go run generate/*.go`

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"text/template"
)

func main() {
	types, err := readTypes()
	if err != nil {
		log.Fatal(err)
	}
	files := []string{
		"new.go",
		"out.go",
	}
	err = generate(types, files)
	if err != nil {
		log.Fatal(err)
	}
	err = goFmt(files)
	if err != nil {
		log.Fatal(err)
	}
}

func generate(types *Types, files []string) error {
	for _, src := range files {
		err := generateFile(types, src)
		if err != nil {
			return err
		}
	}
	return nil
}

type action struct {
	List []string
	Prop string
}

func outName(src string) string {
	return "generated_" + src
}

func generateFile(types *Types, src string) error {
	t := template.New("").Funcs(template.FuncMap{
		"action": func(list []string, prop string) *action {
			return &action{
				List: list,
				Prop: prop,
			}
		},
		"title": strings.Title,
		"typeDef": func(name string) string {
			return fmt.Sprintf("*%v.%v", name, strings.Title(name))
		},
	})
	t, err := t.ParseFiles(src)
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = t.ExecuteTemplate(&buf, "main", types)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(outName(src), buf.Bytes(), 0644)
}

func goFmt(files []string) error {
	for _, src := range files {
		cmd := exec.Command("go", "fmt", outName(src))
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("%s\n", out)
		}
	}
	return nil
}
