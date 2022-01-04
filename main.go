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

func ParseTemplate(certificate_authority_data, server, client_certificate_data, client_key_data string) bool {
	t, err := template.ParseFiles("/config_gotemplate")
	if err != nil {
		log.Print(err)
		return false
	}

	log.Println("Template loaded")

	f, err := os.Create("config")
	if err != nil {
		log.Println("create file: ", err)
		return false
	}

	log.Println("File config created")

	err = t.Execute(f, map[string]string{
		"certificate_authority_data": certificate_authority_data,
		"server":                     server,
		"client_certificate_data":    client_certificate_data,
		"client_key_data":            client_key_data,
	})

	if err != nil {
		log.Print("execute: ", err)
		return false
	}
	log.Println("Template executed successfully")

	f.Close()
	return true
}

func main() {
	certificate_authority_data := GetEnvironment("CERTIFICATE_AUTHORITY_DATA", "localhost")
	server := GetEnvironment("SERVER", "27017")
	client_certificate_data := GetEnvironment("CLIENT_CERTIFICATE_DATA", "")
	client_key_data := GetEnvironment("CLIENT_KEY_DATA", "")

	successfull := ParseTemplate(certificate_authority_data, server, client_certificate_data, client_key_data)

	if successfull {
		log.Println("All successfull")
		os.Exit(0)
	} else {
		log.Println("Something went wrong")
		os.Exit(1)
	}
}
