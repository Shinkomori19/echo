package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Takes 1 line of stdin
	scanner.Scan()
	input := scanner.Text()

	ps := url.Values{}
	ps.Add("content", input)

	res, err := http.PostForm("http://localhost:8000/echo", ps)
	if err != nil {
			log.Fatal(err)
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	fmt.Print(string(body))
}