package main

import (
	"encoding/json"
	"net/http"
)

func getPythonGreets() interface{} {
	resp, err := http.Get("http://127.0.0.1:8000/greet/history")
	if err != nil {
		panic(err)
	}
	var respJSON interface{}
	err = json.NewDecoder(resp.Body).Decode(&respJSON)
	if err != nil {
		panic(err)
	}
	return respJSON

}
