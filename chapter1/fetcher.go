package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func fetchRegular(url string) string {
	r, e := http.Get(url)
	if e != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", e)
		os.Exit(1)
	}
	_, e = ioutil.ReadAll(r.Body)
	r.Body.Close()
	if e != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, e)
		os.Exit(1)
	}
	s := fmt.Sprintf("Handled url %s with StatusCode %d\n", url, r.StatusCode)
	return s
}

func fetchConcurrent(url string, ch chan<- string) {
	r, e := http.Get(url)
	if e != nil {
		ch <- fmt.Sprint(e)
		return
	}
	_, e = ioutil.ReadAll(r.Body)
	r.Body.Close()
	if e != nil {
		ch <- fmt.Sprintf("fetch: reading %s: %v\n", url, e)
		return
	}
	s := fmt.Sprintf("Handled url %s with StatusCode %d\n", url, r.StatusCode)
	ch <- s
}

// RegularFetch : N
func RegularFetch(urls []string) string {
	s := ""
	for _, url := range urls {
		s += fetchRegular(url)
	}
	return s
}

// ConcurrentFetch : N
func ConcurrentFetch(urls []string) string {
	s := ""
	ch := make(chan string)
	for _, url := range urls {
		go fetchConcurrent(url, ch)
	}
	for range urls {
		s += <-ch
	}
	return s
}

func main() {
	fmt.Println(RegularFetch(os.Args[1:]))
	fmt.Println(ConcurrentFetch(os.Args[1:]))
}
