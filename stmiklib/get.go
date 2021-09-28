// package stmiklib is a good package
package stmiklib

// password: iOEWS3DTue
// https://ctp.stmik.ru/api/ctp/telemetry/read.php

import (
	curl "github.com/andelf/go-curl"
)

// curl-client configuration
const (
	url             = "https://ctp.stmik.ru/api/ctp/telemetry/read.php?JSON_PRETTY_PRINT"
	ssl_cert_type   = "P12"
	ssl_verify_peer = false // Discard usage of host certificate files in exchange for the potential spoofing threat
)

// Get transfers bytes of data from STMIK-server using HTTPS-protocol
func Get(p12_cert_path string, p12_password string) ([]byte, error) {
	session := curl.EasyInit()
	defer session.Cleanup()

	session.Setopt(curl.OPT_URL, url)
	session.Setopt(curl.OPT_SSLCERTTYPE, ssl_cert_type)
	session.Setopt(curl.OPT_SSLCERT, p12_cert_path)      //"acservice01.p12"
	session.Setopt(curl.OPT_SSLCERTPASSWD, p12_password) //"iOEWS3DTue"
	session.Setopt(curl.OPT_SSL_VERIFYPEER, ssl_verify_peer)

	// make up a callback function
	var message []byte
	callback := func(buf []byte, userdata interface{}) bool {
		message = append(message, buf...)
		return true
	}
	session.Setopt(curl.OPT_WRITEFUNCTION, callback)
	err := session.Perform()
	return message, err
}
