package main

import (
	"strings"
	"testing"
)

//TODO test with // "-no-passthrough",
var params = []string{
	"goos: windows",
	"goarch: amd64",
	"pkg: test",
	"BenchmarkDominantColors/SmallImages2_Lena-4				225           4870923 ns/op",
	"BenchmarkDominantColors/SmallImages2_WorstCase-4			230           4742135 ns/op",
	"PASS",
	"ok      test  5.128s",
}

var expected = []string{
	"benchmark                                          iter    time/iter",
	"---------                                          ----    ---------",
	"BenchmarkDominantColors/SmallImages2_Lena-4         225   4.87 ms/op",
	"BenchmarkDominantColors/SmallImages2_WorstCase-4    230   4.74 ms/op",
	"",
}

// TODO processing failing benchmarks

func TestBenchmarkOutput(t *testing.T) {
	paramsString := strings.NewReader(strings.Join(params, "\n"))
	benchmarks := processBenchmark(paramsString)

	actual := strings.Split(benchmarks[0].String(), "\n")
	for _, sa := range actual {
		t.Logf("%s\n", sa)
	}

	actualLen := len(actual)
	expectedLen := len(expected)
	if actualLen != expectedLen {
		t.Errorf("actual output is having unexpected line count ACTUAL:%d EXPECTED:%d", actualLen, expectedLen)
	}

	for i := 0; i < expectedLen; i++ {
		actualString := actual[i]
		expectedString := expected[i]
		if actualString != expectedString {
			t.Errorf("Benchmark returns incorrect output (ACTUAL|EXPTECTED): \n\n%s\n%s", actualString, expectedString)
		}
	}
}
