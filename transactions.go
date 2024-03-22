package httpshopify

import (
	"encoding/json"
	"fmt"

	"github.com/MOHC-LTD/httpshopify/v2/internal/http"
	"github.com/MOHC-LTD/shopify/v2"
)

type transactionRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newTransactionRepository(client http.Client, createURL func(endpoint string) string) transactionRepository {
	return transactionRepository{
		client,
		createURL,
	}
}

func (repository transactionRepository) Get(orderID, id int64) (shopify.Transaction, error) {
	url := repository.createURL(fmt.Sprintf("orders/%v/transactions/%v.json", orderID, id))

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return shopify.Transaction{}, err
	}

	var resultDTO struct {
		Transaction TransactionDTO `json:"transaction"`
	}

	if resultDTO.Transaction.ID == 0 {
		return shopify.Transaction{}, shopify.NewErrTransactionNotFound(id)
	}

	json.Unmarshal(body, &resultDTO)

	return resultDTO.Transaction.ToShopify(), nil
}

func (repository transactionRepository) List(orderID int64) (shopify.Transactions, error) {
	url := repository.createURL(fmt.Sprintf("orders/%v/transactions.json", orderID))

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return shopify.Transactions{}, err
	}

	var resultDTO struct {
		Transactions TransactionDTOs `json:"transactions"`
	}

	json.Unmarshal(body, &resultDTO)

	return resultDTO.Transactions.ToShopify(), nil
}

// TransactionDTOs represents a list of shopify Transactions in HTTP requests and responses
type TransactionDTOs []TransactionDTO

// ToShopify converts the DTO to the Shopify equivalent
func (dtos TransactionDTOs) ToShopify() shopify.Transactions {
	Transactions := make(shopify.Transactions, 0, len(dtos))

	for _, dto := range dtos {
		Transactions = append(Transactions, dto.ToShopify())
	}

	return Transactions
}

// PaymentDetailsDTO represents the Payment Details of a Shopify Transaction in HTTP requests and responses
type PaymentDetailsDTO struct {
	CreditCardNumber  string `json:"credit_card_number,omitempty"`
	CreditCardCompany string `json:"credit_card_company,omitempty"`
}

// TransactionDTO represents a Shopify Transaction in HTTP requests and responses
type TransactionDTO struct {
	ID             int64             `json:"id,omitempty"`
	PaymentDetails PaymentDetailsDTO `json:"payment_details,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto TransactionDTO) ToShopify() shopify.Transaction {
	return shopify.Transaction{
		ID: dto.ID,
		PaymentDetails: shopify.PaymentDetails{
			CreditCardNumber:  dto.PaymentDetails.CreditCardNumber,
			CreditCardCompany: dto.PaymentDetails.CreditCardCompany,
		},
	}
}
