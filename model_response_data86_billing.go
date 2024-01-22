/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Billing-related information
type ResponseData86Billing struct {
	Period *BillingPeriod `json:"period"`
	// Date the payment is due
	PaymentDueDate string `json:"payment_due_date"`
	// The minimum payment amount
	MinimumPayment float32 `json:"minimum_payment"`
	// The amount that is past due
	PastDueAmount float32 `json:"past_due_amount"`
}
