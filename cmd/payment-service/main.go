package main

import (
	"fmt"

	"github.com/ahmedaabouzied/iyzipay-go/iyzipay"
)

func main() {

	api_key := "sandbox-afXhZPW0MQlE4dCUUlHcEopnMBgXnAZI"
	secret_key := "sandbox-wbwpzKIiplZxI3hh5ALI4FJyAcZKL6kq"
	base_url := "https://sandbox-api.iyzipay.com"

	options := iyzipay.Options{}
	options.New(api_key, secret_key, base_url)

	// paymentCard := iyzipay.PaymentCard{
	// 	CardHolderName: "John Doe",
	// 	CardNumber:     "5528790000000008",
	// 	ExpireMonth:    "12",
	// 	ExpireYear:     "2030",
	// 	Cvc:            "123",
	// 	RegisterCard:   "0",
	// }

	buyer := iyzipay.Buyer{
		Id:                  "BY789",
		Name:                "John",
		Surname:             "Doe",
		IdentityNumber:      "74300864791",
		Email:               "email@email.com",
		RegistrationAddress: "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
		City:                "Istanbul",
		Country:             "Turkey",
		Ip:                  "::1",
	}

	address := iyzipay.Address{
		ContactName: "Jane Doe",
		City:        "Istanbul",
		Country:     "Turkey",
		Address:     "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
	}

	basketItems := []iyzipay.BasketItem{
		iyzipay.BasketItem{
			Id:        "BI101",
			Name:      "Binocular",
			Category1: "Subscriptions",
			ItemType:  "VIRTUAL",
			Price:     "0.3",
		},
	}

	request := iyzipay.CreateCheckoutFormInitializeRequest{
		Locale:          "tr",
		Price:           "0.3",
		PaidPrice:       "15.0",
		Currency:        "TRY",
		Buyer:           buyer,
		ShippingAddress: address,
		BillingAddress:  address,
		BasketItems:     basketItems,
		CallbackUrl:     "https://www.merchant.com/callback",
	}

	payment := iyzipay.CheckoutFormInitialize{}.Create(request, options)
	fmt.Println(payment)
}
