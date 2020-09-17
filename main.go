package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	generate "github.com/Wryte/graphql-client-generator/generate"
	graphql "github.com/Wryte/graphql-client-generator/graphql"
)

type schemaWrap struct {
	Test string `json:"test"`
	Data struct {
		Schema graphql.Schema `json:"__schema"`
	} `json:"data"`
}

func main() {
	jsonFile, err := os.Open("schema.json")

	if err != nil {
		fmt.Printf("error reading file: %#v", err)
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Printf("error reading whole file: %#v", err)
	}

	w := schemaWrap{}

	err = json.Unmarshal(bytes, &w)
	if err != nil {
		fmt.Printf("error unmarshalling json: %#v", err)
	}

	s := w.Data.Schema
	err = generate.Write(os.Stdout, s)

	if err != nil {
		fmt.Printf("error generating files: %#v", err)
	}
}
