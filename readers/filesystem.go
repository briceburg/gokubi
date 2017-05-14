package readers

import(
  "github.com/briceburg/gokubi"
  "path/filepath"
  "fmt"
  "os"
  "io/ioutil"
  "errors"
//  "strings"
)

// mapping of extension => gokubi method
var ExtMethodMap = map[string]string{
  ".json": "DecodeJSON",
  ".yaml": "DecodeYML",
  ".yml": "DecodeYML",
  ".html": "DecodeXML",
  ".xml": "DecodeXML",
  ".hcl": "DecodeHCL"}


func FileReader(p string, d *gokubi.Data) error {
  ext := filepath.Ext(p)
  method, ok := ExtMethodMap[ext]
  if ok {
    body, err := ioutil.ReadFile(p)
    if err != nil {
       fmt.Fprintf(os.Stderr, "gokubi/filesystem.FileReader: %v", err)
       return err
    }
    return d.Decode(method, body)
  }

  return errors.New(fmt.Sprintf("gokubi/filesystem.FileReader: unsupported path: %v", p))
}

// reads supported files in a directory in lexical order
func DirectoryReader(p string, d gokubi.Data) error{
  files, _ := ioutil.ReadDir("./")
   for _, f := range files {
           fmt.Println(f.Name())
   }
   return nil
}


// reads supported files in a directory and its subdirectories in lexical order
func RecursiveDirectoryReader(p string, d gokubi.Data) error{
  filepath.Walk(p, walkFn)
  return nil
}


// reads supported files in a directory concurrently (disregards order)
func DirectoryReaderFast(p string, d gokubi.Data) error{
  return errors.New("DirectoryReaderFast not implemented")
}


// reads supported in a directory and its subdirectories concurrently (disregards order)
func RecursiveDirectoryReaderFast(p string, d gokubi.Data) error{
  return errors.New("RecursiveDirectoryReaderFast not implemented")
}

func walkFn(p string, info os.FileInfo, err error) error {
  if err != nil {
    fmt.Fprintf(os.Stderr, "gokubi/filesystem.walkFn: failed on %s: %v", p, err)
    return err
  }

  if info.IsDir() {
    return nil
  }

  //ext := strings.ToLower(filepath.Ext(p))
  return nil
}
