package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func callRestEndpoint() {
	response, err := http.Get("http://www.randomnumberapi.com/api/v1.0/random?min=100000000&max=999999999&count=100")
	if err != nil {
		// handle error
	}
	body, err := ioutil.ReadAll(response.Body)

	var v [100]int
	json.Unmarshal(body, &v)

	fmt.Println(v[2])
}
