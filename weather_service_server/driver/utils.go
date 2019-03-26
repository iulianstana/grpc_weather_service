package driver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func extractContent(url string, target interface{}) error {
	spaceClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &target)

	return jsonErr
}
