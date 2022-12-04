package transaction

import (
	"fmt"

	"github.com/samtoya/paystack-client/paystack/common"
	"github.com/samtoya/paystack-client/paystack/common/utils"
	"github.com/samtoya/paystack-client/paystack/config"
	"github.com/samtoya/paystack-client/paystack/dtos"
)

type Client struct {
	AccessKey string
}

func (t *Client) Initialize(req map[string]any) (*common.ApiResponse[dtos.InitializeTransactionDto], error) {
	body, err := utils.MakePostRequest[common.ApiResponse[dtos.InitializeTransactionDto]](config.InitializeTransactionEndpoint, req)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (t *Client) Verify(reference string) (*common.ApiResponse[dtos.VerifyTransactionDto], error) {
	url := fmt.Sprintf("/%s/%s", config.VerifyTransactionEndpoint, reference)
	body, err := utils.MakeGetRequest[common.ApiResponse[dtos.VerifyTransactionDto]](url)
	if err != nil {
		return nil, err
	}

	return body, nil
}
