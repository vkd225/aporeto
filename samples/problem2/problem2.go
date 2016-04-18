package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var strip = regexp.MustCompile("[^a-z0-9A-Z]+")

func main() {
	for i := 1; i < len(os.Args); i++ {
		wordMap := make(map[string]int)

		resp, respErr := http.Get(os.Args[i])
		if respErr != nil {
			fmt.Println(respErr.Error())
			continue
		}
		if resp.StatusCode != 200 {
			fmt.Println("url responded with:", resp.StatusCode)
			continue
		}

		data, dataErr := ioutil.ReadAll(resp.Body)
		if dataErr != nil {
			fmt.Println(dataErr.Error())
			continue
		}

		tokens := strings.Split(string(strip.ReplaceAll(data, []byte(" "))), " ")
		for i := 0; i < len(tokens); i++ {
			if _, ok := wordMap[tokens[i]]; !ok {
				wordMap[tokens[i]] = 0
			}

			wordMap[tokens[i]]++
		}

		fmt.Println(os.Args[i], wordMap)
	}
}
