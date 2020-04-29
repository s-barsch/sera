package main

import (
	"text/template"
	"ioutil"
)

func main() {
}

func readTypes() ([]string, error) {
	types := []string{}
	ioutil.ReadDir("../types")
}
