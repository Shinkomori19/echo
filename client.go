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

	// set input as data
	data := url.Values{
		"content": {input},
	}

	//make POSt request
	res, err := http.PostForm("http://localhost:8000/echo", data)
	if err != nil {
			log.Fatal(err)
	}

	// Not sure this is necessary.
	// defer res.Body.Close()

	// Read the response and print
	body, _ := io.ReadAll(res.Body)
	fmt.Print(string(body))
}
