package readers

import(
  "github.com/briceburg/gokubi"
  "fmt"
  "os"
  "io/ioutil"
)

// supported exentions
var SupportedExts = []string{"json", "yaml", "yml"}

func FileReader(p string, d *gokubi.Data) error {
  body, err := ioutil.ReadFile(p)
  if err != nil {
     fmt.Fprintf(os.Stderr, "gokubi/filesystem.FileReader: %v", err)
     return err
  }
  return d.DecodeJSON(body)
}

func DirectoryReader(p string, d gokubi.Data) error{
  return nil
}
