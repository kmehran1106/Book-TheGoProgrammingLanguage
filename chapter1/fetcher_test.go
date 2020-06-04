package main

import "testing"

// BenchmarkInefficienctFunc : This is to benchmark the ineffient echo function
func BenchmarkRegularFetcher(b *testing.B) {
	args := []string{"https://www.google.com", "https://app.getalice.ai"}
	for i := 0; i < b.N; i++ {
		RegularFetch(args)
	}
}

// BenchmarkEfficienctFunc : This is to benchmark the efficient echo function
func BenchmarkConcurrentFetcher(b *testing.B) {
	args := []string{"https://www.google.com", "https://app.getalice.ai"}
	for i := 0; i < b.N; i++ {
		ConcurrentFetch(args)
	}
}
