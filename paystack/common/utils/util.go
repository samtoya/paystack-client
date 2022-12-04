package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/samtoya/paystack-client/paystack/config"
	"github.com/spf13/viper"
)

func MakeGetRequest[T any](endpoint string) (*T, error) {
	token, err := validateTokenPresent()
	if err != nil {
		log.Fatalf(err.Error())
	}
	url := fmt.Sprintf("%s%s", config.BaseUrl, endpoint)
	log.Println("full url:", url)

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	response := new(T)
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return response, nil
}

func MakePostRequest[T any](endpoint string, payload map[string]any) (*T, error) {
	token, err := validateTokenPresent()
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s", config.BaseUrl, endpoint)
	client := http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	log.Printf("resp status code: %v", resp.StatusCode)
	response := new(T)
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return response, nil
}

func validateTokenPresent() (string, error) {
	secretKey := viper.GetString("secretKey")
	if secretKey == "" {
		return "", fmt.Errorf("please provide your secret key")
	}

	return secretKey, nil
}
