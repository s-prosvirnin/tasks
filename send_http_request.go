package main

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
)

const requestUrl = "https://enm3fdguu1gx.x.pipedream.net/"

func main() {
	sendPostRequest("halo!")
}

func sendPostRequest(text string) {
	var buf bytes.Buffer
	n, err := buf.WriteString(text)
	if err != nil || n == 0 {
		log.Fatal(err)
	}
	// https://public.requestbin.com/r/enm3fdguu1gx/2PQTw5B8Ws03qCUqVw1G1dqGX4P
	resp, err := http.Post(
		requestUrl,
		"text/plain",
		&buf,
	)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal("sendPostRequest error status" + strconv.Itoa(resp.StatusCode))
	}
}
