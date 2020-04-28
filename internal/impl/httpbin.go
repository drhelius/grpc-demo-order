package impl

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type publicIp struct {
	Origin string
}

func getIp() string {

	log.Printf("[Order] Invoking httpbin IP service")

	response, err := http.Get("https://httpbin/ip")

	var public_ip string

	if err != nil {
		log.Fatalf("[Order] httpbin IP service failed with error: %s", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		var response publicIp

		json.Unmarshal([]byte(data), &response)

		public_ip = response.Origin

		log.Printf("[Order] httpbin IP service: %s", string(data))
	}

	return public_ip
}
