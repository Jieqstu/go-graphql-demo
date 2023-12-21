package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func importJSONFromFile(fileName string, result interface{}) (isOK bool) {
	isOK = true
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("Error:", err)
		isOK = false
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		isOK = false
		fmt.Print("Error:", err)
	}
	return
}

func FetchOdds(url string) []odds {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	list := []odds{}
	err = json.Unmarshal(data, &list)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return list
}

func FetchOdd(url string) *odds {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	odd := &odds{}
	err = json.Unmarshal(data, odd)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return odd
}

func GetApiKey() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	apiKey := os.Getenv("API_KEY")
	return apiKey
}
