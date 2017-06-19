package main

import (
	"fmt"

	"github.com/briceburg/gokubi"
)

func main() {
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
	//fmt.Println(data.EncodeBash())

	/*
		if err := readers.DirectoryReader("fixtures", &data); err != nil {
			panic(err)
		}
		fmt.Println(data.String())
		fmt.Println("%+v", data)
	*/
}
