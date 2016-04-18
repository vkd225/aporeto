package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

var strip = regexp.MustCompile("[^a-z0-9A-Z]+")

func main() {
	var waitgrp sync.WaitGroup
	for i := 1; i < len(os.Args); i++ {
		waitgrp.Add(1)
		go func(url string) {
			process(url)
			waitgrp.Done()
		}(os.Args[i])
	}

	waitgrp.Wait()
}

func process(url string) {
	wordMap := make(map[string]int)

	resp, respErr := http.Get(url)
	if respErr != nil {
		fmt.Println(respErr.Error())
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("url responded with:", resp.StatusCode)
		return
	}

	data, dataErr := ioutil.ReadAll(resp.Body)
	if dataErr != nil {
		fmt.Println(dataErr.Error())
		return
	}

	tokens := strings.Split(string(strip.ReplaceAll(data, []byte(" "))), " ")
	for i := 0; i < len(tokens); i++ {
		if _, ok := wordMap[tokens[i]]; !ok {
			wordMap[tokens[i]] = 0
		}

		wordMap[tokens[i]]++
	}

	fmt.Println(url, wordMap)
}
