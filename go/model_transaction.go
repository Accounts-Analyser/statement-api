/*
 * Accounts Analyser
 *
 * This API is for managing bank account statements, for use with [https://github.com/Accounts-Analyser](https://github.com/Accounts-Analyser) 
 *
 * API version: 1.0.0
 * Contact: colincampbell321123@hotmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Transaction struct {

	Id int64 `json:"id,omitempty"`

	TransactionId int64 `json:"transactionId,omitempty"`

	Date string `json:"date,omitempty"`

	TransactionType string `json:"transactionType,omitempty"`

	Description string `json:"description,omitempty"`

	PaidOut float32 `json:"paidOut,omitempty"`

	PaidIn float32 `json:"paidIn,omitempty"`

	Balance float32 `json:"balance,omitempty"`
}
