package impl

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type publicIp struct {
	Origin string
}

func getIp() string {

	log.Printf("[Order] Invoking httpbin IP service")

	https_insecure, err := strconv.ParseBool(os.Getenv("HTTPBIN_INSECURE"))

	if err != nil {
		https_insecure = false
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: https_insecure},
	}

	client := &http.Client{Transport: tr}

	httpbin_protocol := os.Getenv("HTTPBIN_PROTOCOL")
	httpbin_host := os.Getenv("HTTPBIN_HOST")
	httpbin_port := os.Getenv("HTTPBIN_PORT")

	if httpbin_protocol == "" {
		httpbin_protocol = "https"
	}

	if httpbin_host == "" {
		httpbin_host = "httpbin.org"
	}

	if httpbin_port == "" {
		httpbin_port = "443"
	}

	response, err := client.Get(httpbin_protocol + "://" + httpbin_host + ":" + httpbin_port + "/ip")

	var public_ip string

	if err != nil {
		log.Fatalf("[Order] httpbin IP service failed with error: %s", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		var response publicIp

		json.Unmarshal([]byte(data), &response)

		public_ip = response.Origin

		log.Printf("[Order] httpbin IP service: %s", public_ip)
	}

	return public_ip
}
