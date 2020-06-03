package main

import "testing"

// BenchmarkInefficienctFunc : This is to benchmark the ineffient echo function
func BenchmarkInefficienctFunc(b *testing.B) {
	args := []string{"a", "b", "c", "d", "e", "f", "g"}
	for i := 0; i < b.N; i++ {
		InefficientEcho(args)
	}
}

// BenchmarkEfficienctFunc : This is to benchmark the efficient echo function
func BenchmarkEfficienctFunc(b *testing.B) {
	args := []string{"a", "b", "c", "d", "e", "f", "g"}
	for i := 0; i < b.N; i++ {
		EfficientEcho(args)
	}
}
