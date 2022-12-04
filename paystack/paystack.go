package paystack

import (
	"github.com/samtoya/paystack-client/paystack/core/customer"
	"github.com/samtoya/paystack-client/paystack/core/transaction"
	"github.com/spf13/viper"
)

type PClient struct {
	//Charge                  *charge.Client
	Customer *customer.Client
	//Dispute                 *dispute.Client
	//DedicatedVirtualAccount *dva.Client
	//Invoice                 *invoice.Client
	//Miscellaneous           *miscellaneous.Client
	//PaymentPage             *paymentpage.Client
	//Plan                    *plan.Client
	//Product                 *product.Client
	//Refund                  *refund.Client
	//Settlement              *settlement.Client
	//SubAccount              *subaccount.SubAccount
	//Subscription            *subscription.Client
	Transaction *transaction.Client
	//Transfer                *transfer.Client
	//TransferControl         *transfercontrol.Client
	//TransferRecipient       *transferrecipient.Client
}

func New(secretKey string) *PClient {
	// Store key in config global
	viper.Set("secretKey", secretKey)
	// Initialize the paystack client
	client := new(PClient)

	return client
}

func (c *PClient) NewTransaction() *transaction.Client {
	return new(transaction.Client)
}
