package gokubi

import (
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func BenchmarkInterfaceMapToStringMap(b *testing.B) {
	body, err := ioutil.ReadFile("fixtures/music.yml")
	if err != nil {
		b.Skipf("BenchmarkInterfaceMapToStringMap: failure: %v", err)
	}

	var yamlMap map[interface{}]interface{}
	if err := yaml.Unmarshal(body, &yamlMap); err != nil {
		b.Skipf("BenchmarkInterfaceMapToStringMap: failure: %v", err)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = InterfaceMapToStringMap(yamlMap)
	}
}

func BenchmarkEnvMarshal(b *testing.B) {
	data := make(Data)
	if err := FileReader("fixtures/music.yml", &data); err != nil {
		b.Skipf("BenchmarkEnvMarshal: failure: %v", err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if _, err := EnvMarshal(data); err != nil {
			b.Skipf("BenchmarkEnvMarshal: failure: %v", err)
		}
	}
}
