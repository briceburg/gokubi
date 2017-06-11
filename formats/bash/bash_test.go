package bash

import (
	"testing"

	"github.com/briceburg/gokubi/readers"
)

func BenchmarkMarshal(b *testing.B) {
	data := make(map[string]interface{})
	if err := readers.FileReader("fixtures/music.yml", &data); err != nil {
		b.Skipf("BenchmarMarshal: failure: %v", err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if _, err := Marshal(data); err != nil {
			b.Skipf("BenchmarMarshal: failure: %v", err)
		}
	}
}
