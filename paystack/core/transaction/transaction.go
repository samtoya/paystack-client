package transaction

import (
	"encoding/json"
	"fmt"
	. "github.com/samtoya/paystack-client/paystack/common"
	"github.com/samtoya/paystack-client/paystack/common/utils"
	"github.com/samtoya/paystack-client/paystack/config"
	"github.com/samtoya/paystack-client/paystack/dtos"
)

type Client struct {
	AccessKey string
}

func (t *Client) Initialize(email string, amount float32) (*ApiResponse[dtos.InitializeTransactionDto], error) {
	req := make(map[string]string)
	req["email"] = email
	req["amount"] = fmt.Sprintf("%f", amount)
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := utils.MakePostRequest(config.InitializeTransactionEndpoint, req)
	if err != nil {
		return nil, err
	}
	response := new(ApiResponse[dtos.InitializeTransactionDto])
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	fmt.Printf("response: %v", response)

	return response, nil
}

func (t *Client) Verify(reference string) (*ApiResponse[dtos.VerifyTransactionDto], error) {
	url := fmt.Sprintf("%s/%s", config.VerifyTransactionEndpoint, reference)
	body, err := utils.MakeGetRequest(url)
	if err != nil {
		return nil, err
	}

	response := new(ApiResponse[dtos.VerifyTransactionDto])
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	fmt.Printf("response: %v", response)

	return response, nil
}
