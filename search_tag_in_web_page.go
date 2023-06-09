package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

const url = "https://habr.com/ru/flows/develop/top10/"
const searchTag = "meta name=\"viewport\""

func main() {
	resp, err := http.Get(url)
	if err != nil {
		handleError(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		handleError(errors.Errorf("response status code: %d", resp.StatusCode))
	}

	search := []byte(searchTag)

	curPos := 0
	for {
		bodyBuf := make([]byte, 14336)
		b, err := resp.Body.Read(bodyBuf)
		if err != nil && err != io.EOF {
			handleError(err)
		}
		if err == nil && b == 0 {
			handleError(errors.New("body read 0 byte"))
		}
		for _, b := range bodyBuf {
			if curPos == len(search) {
				break
			}
			if b == search[curPos] {
				curPos++
			} else {
				curPos = 0
			}
		}
		if err != nil && err == io.EOF {
			break
		}
	}

	if curPos == len(search) {
		handleResult("MATCH")
		return
	}

	handleResult("NOT MATCH")
}

func handleError(err error) {
	log.Fatal(err)
}

func handleResult(res string) {
	fmt.Println(res)
}
