package main

import (
	"fmt"
	"github.com/samtoya/paystack-client/paystack"
)

func main() {
	client := paystack.New("sk_test_c6017df73901c9c13dc8404e103c72507dbc2ad1")
	transaction := client.NewTransaction()
	resp, err := transaction.Initialize("joseph@elogence.com", 5000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response: %v", resp.Data.Reference)

	//resp, err := transaction.Verify("c03auk3y2wccjtm")
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//_ = fmt.Sprintf("response: %v", resp.Data)
}
