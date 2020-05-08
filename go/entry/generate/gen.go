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

var typeDir = "types"

func main() {
	types, err := readTypes()
	if err != nil {
		log.Fatal(err)
	}
	err = printMethods(types)
	if err != nil {
		log.Fatal(err)
	}
	err = goFmt(types)
	if err != nil {
		log.Fatal(err)
	}
}

func printMethods(types []string) error {
	t, err := loadTemplate(typeDir + "/_out.go")
	if err != nil {
		return err
	}
	for _, typ := range types {
		err := printOutFile(typ, t)
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

func printOutFile(typ string, tmpl *template.Template) error {
	buf := bytes.Buffer{}
	err := tmpl.ExecuteTemplate(&buf, "out", typ)
	if err != nil {
		return err
	}
	path := outFilePath(typ)
	err = ioutil.WriteFile(path, buf.Bytes(), 0644)
	if err != nil {
		return err
	}
	fmt.Printf("written: %v\n", path)
	return nil
}

func outFilePath(typ string) string {
	outName := "out_gen.go"
	if isMedia(typ) {
		return fmt.Sprintf("%v/media/%v/%v", typeDir, typ, outName)
	}
	return fmt.Sprintf("%v/%v/%v", typeDir, typ, outName)
}

func goFmt(types []string) error {
	for _, typ := range types {
		cmd := exec.Command("go", "fmt", outFilePath(typ))
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("%s\n", out)
		}
	}
	return nil
}

func loadTemplate(path string) (*template.Template, error) {
	t := template.New("").Funcs(template.FuncMap{
		"receiver": func(typ string) string {
			return fmt.Sprintf("(e *%v)", strings.Title(typ))
		},
		"isMedia": isMedia,
		"isTree":  isTree,
		"title":   strings.Title,
		"typeDef": func(name string) string {
			return fmt.Sprintf("*%v.%v", name, strings.Title(name))
		},
	})
	return t.ParseFiles(path)
}

func firstChar(typ string) string {
	if len(typ) > 1 {
		return string(typ[0])
	}
	return typ
}

func isMedia(typ string) bool {
	switch typ {
	case "tree", "set":
		return false
	}
	return true
}

func isTree(typ string) bool {
	return typ == "tree"
}
