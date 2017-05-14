package main

import(
  "github.com/briceburg/gokubi"
  "github.com/briceburg/gokubi/readers"
  "fmt"
)

func main(){
  data := make(gokubi.Data)
  if err := readers.FileReader("fixtures/music.xml", &data); err != nil {
    panic(err)
  }
  fmt.Println(data.String())
  fmt.Println(data.EncodeXML())
  fmt.Println(data.EncodeYML())


}
