package payment

import (
	"net/url"

	cartDomain "go.aoe.com/flamingo/core/cart/domain/cart"
	"go.aoe.com/flamingo/framework/web"
)

type (
	PaymentMethod struct {
		Title               string
		Code                string
		IsExternalPayment   bool
		ExternalRedirectUri string
	}

	PaymentProvider interface {
		GetCode() string
		// GetPaymentMethods returns the Payment Providers available Payment Methods
		GetPaymentMethods() []PaymentMethod
		// RedirectExternalPayment starts a Redirect to an external Payment Page (if applicable)
		RedirectExternalPayment(web.Context, *PaymentMethod, *url.URL) (web.Response, error)
		// ProcessPayment
		ProcessPayment(web.Context, *PaymentMethod) (*cartDomain.CartPayment, error)
		IsActive() bool
	}
)
