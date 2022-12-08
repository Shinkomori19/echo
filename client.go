package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
)

type Data struct {
	Content string `json:"content"`
	Name string `json:"name"`
	Age int `json:"age"`
	Company string `json:"company"`
}

// Makes a POST req using json.
func main() {
	// takes inputs and sets it to data
	var inputs [3]string
	age := 2
	fmt.Scan(&inputs[0], &inputs[1], &age, &inputs[2])

	data := Data{
		Content:inputs[0],
		Name:inputs[1],
		Age:age,
		Company:inputs[2],
	}
	// fmt.Println(data)

	// Parse to json
	jsonString, err := json.Marshal(data)
	handleErr(err)

	// Make request
	url := "http://localhost:8000/echo/post"
	resp, err := http.Post(url, "application/json",bytes.NewBuffer(jsonString))
	handleErr(err)

	defer resp.Body.Close()

	// Decode response and print it out
	resjs := Data{}
	json.NewDecoder(resp.Body).Decode(&resjs)
	a,_ := PrettyStruct(resjs)
	fmt.Println(a)
}

// Pretty prints struct. Found online.
func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
			return "", err
	}
	return string(val), nil
}

// Simple err handler
func handleErr(err interface{}) {
	if err != nil {
		panic("Error")
	}
	return
}


// Previous version of main()

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	// Takes 1 line of stdin
// 	scanner.Scan()
// 	input := scanner.Text()

// 	ps := url.Values{}
// 	ps.Add("content", input)

// 	res, err := http.PostForm("http://localhost:8000/echo", ps)
// 	if err != nil {
// 			log.Fatal(err)
// 	}

// 	defer res.Body.Close()

// 	body, _ := io.ReadAll(res.Body)
// 	fmt.Print(string(body))
// }
