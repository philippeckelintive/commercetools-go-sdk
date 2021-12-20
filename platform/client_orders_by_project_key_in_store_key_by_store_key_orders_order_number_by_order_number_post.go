// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost struct {
	body    OrderUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPostInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPostInput) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Updates an order in the store specified by {storeKey}. The {storeKey} path parameter maps to a Store's key.
*	If the order exists in the commercetools project but does not have the store field,
*	or the store field references a different store, this method returns a ResourceNotFound error.
*	In case the orderNumber does not match the regular expression [a-zA-Z0-9_\-]+,
*	it should be provided in URL-encoded format.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost) Execute(ctx context.Context) (result *Order, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.post(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		data,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 409, 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
		}
		return nil, result
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
