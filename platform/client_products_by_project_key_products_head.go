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

type ByProjectKeyProductsRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductsRequestMethodHeadInput
}

func (r *ByProjectKeyProductsRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductsRequestMethodHeadInput struct {
	Where []string
}

func (input *ByProjectKeyProductsRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyProductsRequestMethodHead) Where(v []string) *ByProjectKeyProductsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodHeadInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodHead) WithQueryParams(input ByProjectKeyProductsRequestMethodHeadInput) *ByProjectKeyProductsRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductsRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyProductsRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Check if Products exist. Responds with a `200 OK` status if any Products match the Query Predicate, or `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyProductsRequestMethodHead) Execute(ctx context.Context) error {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.head(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		return nil

	case 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return result
	}

}
