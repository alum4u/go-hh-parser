package grabber

import (
	"net/http"
	"log"
	"io"
)

func Grab(url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)

		return "" , err;
	}

	body, _:= io.ReadAll(res.Body)

	res.Body.Close()

	return string(body), nil
}