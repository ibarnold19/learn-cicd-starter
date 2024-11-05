package auth

import (
	"log"
	"net/http"
	"testing"
)

func TestValidAuthHeader(t *testing.T) {

	req, err := http.NewRequest("GET", "http://localhost:8080", nil)

	if err != nil {

		t.Skip("Unable to create mock request.")

	}

	req.Header.Add("Authorization", "ApiKey 1234567890")

	api_key_result, err := GetAPIKey(req.Header)

	if err != nil || api_key_result == "" {

		log.Printf("Expected a valid auth header! Header passed: %+v", req.Header)

		log.Printf("Api key response: %+s. Error: %+v", api_key_result, err)

		t.Fail()

	}

}

func TestMissingAuthHeader(t *testing.T) {

	req, err := http.NewRequest("GET", "http://localhost:8080", nil)

	if err != nil {

		t.Skip("Unable to create mock request.")

	}

	api_key_result, err := GetAPIKey(req.Header)

	if err == nil || api_key_result != "" {

		log.Printf("Expected no auth header! Header passed: %+v", req.Header)

		log.Printf("Api key response: %+s. Error: %+v", api_key_result, err)

		t.Fail()

	}

}

func TestCombinedApiKey(t *testing.T) {

	req, err := http.NewRequest("GET", "http://localhost:8080", nil)

	if err != nil {

		t.Skip("Unable to create mock request.")

	}

	// Combining 'ApiKey' with the key itself
	req.Header.Add("Authorization", "ApiKey1234567890")

	api_key_result, err := GetAPIKey(req.Header)

	if err == nil || api_key_result != "" {

		log.Printf("Expected bad ApiKey value! Header passed: %+v", req.Header)

		log.Printf("Api key response: %+s. Error: %+v", api_key_result, err)

		t.Fail()

	}

}
func TestMispelledApiKey(t *testing.T) {

	req, err := http.NewRequest("GET", "http://localhost:8080", nil)

	if err != nil {

		t.Skip("Unable to create mock request.")

	}

	// Mispell ApiKey
	req.Header.Add("Authorization", "Apiey 1234567890")

	api_key_result, err := GetAPIKey(req.Header)

	if err == nil || api_key_result != "" {

		log.Printf("Expected bad ApiKey value! Header passed: %+v", req.Header)

		log.Printf("Api key response: %+s. Error: %+v", api_key_result, err)

		t.Fail()

	}

}
