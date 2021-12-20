// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInventoryRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyInventoryRequestBuilder) WithId(id string) *ByProjectKeyInventoryByIDRequestBuilder {
	return &ByProjectKeyInventoryByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInventoryRequestBuilder) Get() *ByProjectKeyInventoryRequestMethodGet {
	return &ByProjectKeyInventoryRequestMethodGet{
		url:    fmt.Sprintf("/%s/inventory", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInventoryRequestBuilder) Post(body InventoryEntryDraft) *ByProjectKeyInventoryRequestMethodPost {
	return &ByProjectKeyInventoryRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/inventory", rb.projectKey),
		client: rb.client,
	}
}
