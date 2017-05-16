package main

import (
	"fmt"

	"github.com/briceburg/gokubi"
)

func main() {
	data := make(gokubi.Data)

	if err := gokubi.FileReader("fixtures/music.yml", &data); err != nil {
		panic(err)
	}
	fmt.Println(data.String())
	//fmt.Println(data.EncodeXML())
	fmt.Println(data.EncodeENV())

	/*
		if err := readers.DirectoryReader("fixtures", &data); err != nil {
			panic(err)
		}
		fmt.Println(data.String())
		fmt.Println("%+v", data)
	*/
}
