package main

import(
  "github.com/briceburg/gokubi"
  "github.com/briceburg/gokubi/readers"
  "fmt"
)

func main(){
  data := make(gokubi.Data)

  readers.FileReader("fixtures/json/music.json", &data)
  //readers.FileReader("fixtures/yml/music.yml", &data)

  fmt.Println(data.String())
}
