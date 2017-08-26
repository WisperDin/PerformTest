package main

import (
	"testing"
	"./routine"
	"./conf"
)
func init(){
	conf.Init("./test.toml")
}

func BenchmarkEmpRoutine(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		routine.EleRoutine(1,n)
	}
}