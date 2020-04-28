package impl

import (
	"io/ioutil"
	"log"
	"net/http"
)

func getIp() string {

	log.Printf("[Order] Invoking httpbin IP service")

	response, err := http.Get("https://httpbin.org/ip")

	var data string

	if err != nil {
		log.Fatalf("[Order] httpbin IP service failed with error: %s", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		log.Printf("[Order] httpbin IP service: %s", string(data))
	}

	return string(data)
}
