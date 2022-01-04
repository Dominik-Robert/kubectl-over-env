package main

import (
	"log"
	"os"
	"text/template"
)

func GetEnvironment(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultValue
	}
	return val
}

func ParseTemplate(certificate_authority_data, server, client_certificate_data, client_key_data string) {
	t, err := template.ParseFiles("/config_gotemplate")
	if err != nil {
		log.Print(err)
	}

	f, err := os.Create("config")
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = t.Execute(f, map[string]string{
		"certificate_authority_data": certificate_authority_data,
		"server":                     server,
		"client_certificate_data":    client_certificate_data,
		"client_key_data":            client_key_data,
	})

	if err != nil {
		log.Print("execute: ", err)
		return
	}

	f.Close()

}

func main() {
	certificate_authority_data := GetEnvironment("CERTIFICATE_AUTHORITY_DATA", "localhost")
	server := GetEnvironment("SERVER", "27017")
	client_certificate_data := GetEnvironment("CLIENT_CERTIFICATE_DATA", "")
	client_key_data := GetEnvironment("CLIENT_KEY_DATA", "")

	ParseTemplate(certificate_authority_data, server, client_certificate_data, client_key_data)
}
