package common

import (
	"net/http"
)

var (
	client http.Client{}
)

// SendData to some destinations
func SendData(data []byte, destinationURLs []string) {
	for _, URL := range destinationURLs {
		go send(data, URL)	
	}
}

func send(data []byte, URL string) {
	
}