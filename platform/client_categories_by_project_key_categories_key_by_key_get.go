package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyCategoriesKeyByKeyRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCategoriesKeyByKeyRequestMethodGetInput
}

func (r *ByProjectKeyCategoriesKeyByKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCategoriesKeyByKeyRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyCategoriesKeyByKeyRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCategoriesKeyByKeyRequestMethodGet) Expand(v []string) *ByProjectKeyCategoriesKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCategoriesKeyByKeyRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCategoriesKeyByKeyRequestMethodGet) WithQueryParams(input ByProjectKeyCategoriesKeyByKeyRequestMethodGetInput) *ByProjectKeyCategoriesKeyByKeyRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCategoriesKeyByKeyRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyCategoriesKeyByKeyRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Either the [scope](/../api/scopes) `view_products:{projectKey}` or `view_categories:{projectKey}` is required.
*
 */
func (rb *ByProjectKeyCategoriesKeyByKeyRequestMethodGet) Execute(ctx context.Context) (result *Category, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.get(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj

	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
