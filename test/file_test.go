package test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFile(t *testing.T) {
	dir := http.Dir("../")
	fmt.Println(dir)
	server := http.FileServer(dir)
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		return
	}
}
