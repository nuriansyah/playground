package compareFast_test

import (
	"strconv"
	"testing"

	"github.com/ruang-guru/playground/backend/benchmarking-and-compiler/compare"
)

func BenchmarkUnmarshal(b *testing.B) {

	for n := 0; n < b.N; n++ {
		umur := strconv.Itoa(n)
		data := `{"Nama":"user", "Umur: "` + umur + `"}`
		compare.UnmarshallFast(data)
	}

}
