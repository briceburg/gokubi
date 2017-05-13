package readers

import (
  "testing"
  "github.com/briceburg/gokubi"
)

func BenchmarkFileReader(b *testing.B) {
  data := make(gokubi.Data)
  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    if err := FileReader("../fixtures/json/music.json", &data); err != nil {
      b.Skipf("BenchmarkFileReader: failure: %v", err)
    }
  }
}
