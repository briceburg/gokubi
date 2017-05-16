package gokubi

import "testing"

func BenchmarkFileReader(b *testing.B) {
	data := make(Data)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if err := FileReader("fixtures/music.json", &data); err != nil {
			b.Skipf("BenchmarkFileReader: failure: %v", err)
		}
	}
}
