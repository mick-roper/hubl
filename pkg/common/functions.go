package common

import (
	"net/http"
)

var (
	client http.Client{}
)

// SendData to some destinations
func SendData(data []byte, subscriptions []Subscription) {
	for _, s := range subscriptions {
		reader := nil
		go send(reader, s)
	}
}

func send(data io.Reader, s *Subscription) {
	req, err := http.NewRequest(http.MethodPost, s.DestinationURL, data)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if res.StatusCode < 300 {
		s.Metrics.Success += 1
	} else {
		s.Metrics.Failure += 1
	}

	err = s.Save()

	if err != nil {
		panic(err)
	}
}