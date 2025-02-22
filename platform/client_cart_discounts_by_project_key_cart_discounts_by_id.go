package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCartDiscountsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyCartDiscountsByIDRequestBuilder) Get() *ByProjectKeyCartDiscountsByIDRequestMethodGet {
	return &ByProjectKeyCartDiscountsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/cart-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCartDiscountsByIDRequestBuilder) Post(body CartDiscountUpdate) *ByProjectKeyCartDiscountsByIDRequestMethodPost {
	return &ByProjectKeyCartDiscountsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/cart-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCartDiscountsByIDRequestBuilder) Delete() *ByProjectKeyCartDiscountsByIDRequestMethodDelete {
	return &ByProjectKeyCartDiscountsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/cart-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
