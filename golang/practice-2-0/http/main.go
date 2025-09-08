package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	data := []byte(`{"name:"Bob"}`)

	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
