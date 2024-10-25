package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/paymentintent"
)

type item struct {
	Id     string
	Amount int64
}

func calculateOrderAmount(items []item) int64 {
	// Calculate the order total on the server to prevent
	// people from directly manipulating the amount on the client
	var total int64
	total = 0
	for _, item := range items {
		total += int64(item.Amount)
	}
	return total
}

func HandleCreatePaymentIntent(c *fiber.Ctx) error {

	stripe.Key = "sk_test_51Q5bFgCQw7GcQLHIDImiHQkgmiLHlPmGOb6groWAuKwFsqjUahUETjuVm1PzDgPFIqjPoCErf8N8qTyLb8CqINaR00meLMPOjH"

	var req struct {
		Items []item `json:"items"`
	}

	// Parse the request body into the `req` struct
	if err := c.BodyParser(&req); err != nil {
		log.Printf("c.BodyParser: %v", err)
		return c.Status(http.StatusBadRequest).SendString("Bad Request")
	}

	// Create a PaymentIntent with amount and currency
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculateOrderAmount(req.Items)),
		Currency: stripe.String(string(stripe.CurrencyBRL)),
		// In the latest version of the API, specifying the `automatic_payment_methods` parameter is optional because Stripe enables its functionality by default.
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	log.Printf("pi.New: %v", pi.ClientSecret)

	if err != nil {
		log.Printf("pi.New: %v", err)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(fiber.Map{
		"clientSecret":   pi.ClientSecret,
		"dpmCheckerLink": fmt.Sprintf("https://dashboard.stripe.com/settings/payment_methods/review?transaction_id=%s", pi.ID),
	})
}
