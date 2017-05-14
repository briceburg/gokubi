package main

import(
  "github.com/briceburg/gokubi"
  "github.com/briceburg/gokubi/readers"
  "fmt"
)

func main(){
  data := make(gokubi.Data)
  if err := readers.FileReader("fixtures/yml/music.yml", &data); err != nil {
    panic(err)
  }
  fmt.Println(data.String())

  data = make(gokubi.Data)
  if err := readers.FileReader("fixtures/json/music.json", &data); err != nil {
    panic(err)
  }
  fmt.Println(data.String())
  //readers.FileReader("fixtures/yml/music.yml", &data)


}
