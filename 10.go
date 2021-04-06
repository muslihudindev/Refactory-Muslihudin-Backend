package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getjson(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
	}

	return body
}

func main() {

	var user []interface{}
	var post []interface{}

	err := json.Unmarshal(getjson("https://jsonplaceholder.typicode.com/users"), &user)
	if err != nil {
		panic(err)
	}
	err2 := json.Unmarshal(getjson("https://jsonplaceholder.typicode.com/posts"), &post)
	if err2 != nil {
		panic(err2)
	}

	result := []map[string]interface{}{}
	var m map[string]interface{}
	var mu map[string]interface{}
	for k, _ := range post {
		m = post[k].(map[string]interface{})
		for ku, _ := range user {
			mu = user[ku].(map[string]interface{})
			if m["userId"] == mu["id"] {
				m["user"] = mu
				result = append(result, m)
			}
		}

	}
	//fmt.Println(result)
	jsonString, err := json.Marshal(result)
	fmt.Println(string(jsonString))
}
