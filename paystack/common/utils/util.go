package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/samtoya/paystack-client/paystack/config"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

func MakeGetRequest(endpoint string) ([]byte, error) {
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
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	return body, nil
}

func MakePostRequest(endpoint string, payload map[string]any) ([]byte, error) {
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
		"Content-Type":  {"application/jsonData"},
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func validateTokenPresent() (string, error) {
	secretKey := viper.GetString("secretKey")
	if secretKey == "" {
		return "", fmt.Errorf("please provide your secret key")
	}

	return secretKey, nil
}
