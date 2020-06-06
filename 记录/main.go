package main

import (
	"fmt"
	"net/http"
)

type t1 struct{}
type t2 struct{}

func (t *t1) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("t1"))
}
func (t *t2) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("t2"))
}
func main() {
	// http.Handle("/hello/", new(t1))
	// http.Handle("/hello/server", new(t2))
	// http.ListenAndServe(":8000", nil)
	var str = "pwwkecdefwghtij"
	r := MaxString(str)
	fmt.Println(r)

}

func MaxString(s string) int {
	m := make(map[byte]int)
	maxLength := 0
	start := 0
	for k, v := range []byte(s) {
		d, ok := m[v]
		if ok && d >= start {
			//if k-d > maxLength {
			//fmt.Println(maxLength, d-start, k-(d+1))
			maxLength = max(maxLength, d-start, k-(d+1)) //k - d)
			//}
			start = d + 1
		}
		m[v] = k
	}
	if len(s)-start-1 > maxLength {
		//fmt.Println(len(s)-start-1, maxLength)
		maxLength = len(s) - start - 1
	}
	return maxLength + 1
}

func max(i ...int) int {
	max := 0
	for _, v := range i {
		if v > max {
			max = v
		}
	}
	return max
}
