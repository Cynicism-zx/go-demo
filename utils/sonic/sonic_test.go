package sonic

import (
	"bytes"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/encoder"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

// https://github.com/bytedance/sonic?utm_source=gold_browser_extension 字节跳动序列化与反序列化库

type Schema struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

//goos: linux
//goarch: amd64
//pkg: go-demo/utils/sonic
//cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
//BenchmarkSonicMarshal-16        12871128               470.1 ns/op
//BenchmarkJsoniterMarshal-16     12473517               479.2 ns/op
//PASS
//ok      go-demo/utils/sonic     12.997s

func BenchmarkSonicMarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var data = &Schema{
			Name:    "dulipa",
			Age:     88,
			Address: "china",
		}
		var res Schema
		out, err := sonic.Marshal(&data)
		if err != nil {
			b.Error(err.Error())
			return
		}
		//b.Log(out)
		if err = sonic.Unmarshal(out, &res); err != nil {
			b.Error(err.Error())
			return
		}
		//b.Log(res)
	}
}

func BenchmarkJsoniterMarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var data = &Schema{
			Name:    "dulipa",
			Age:     88,
			Address: "china",
		}
		var res Schema
		out, err := jsoniter.Marshal(&data)
		if err != nil {
			b.Error(err.Error())
			return
		}
		//b.Log(out)
		if err = jsoniter.Unmarshal(out, &res); err != nil {
			b.Error(err.Error())
			return
		}
		//b.Log(res)
	}
}

func BenchmarkSonicStreaming(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var o1 = map[string]interface{}{
			"a": "b",
		}
		var w1 = bytes.NewBuffer(nil)
		var enc = encoder.NewStreamEncoder(w1)
		enc.Encode(o1)
		println(w1.String())
	}
}
