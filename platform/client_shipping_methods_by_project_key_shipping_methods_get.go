package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyShippingMethodsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyShippingMethodsRequestMethodGetInput
}

func (r *ByProjectKeyShippingMethodsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyShippingMethodsRequestMethodGetInput struct {
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyShippingMethodsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.WithTotal != nil {
		if *input.WithTotal {
			values.Add("withTotal", "true")
		} else {
			values.Add("withTotal", "false")
		}
	}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	for k, v := range input.PredicateVar {
		for _, x := range v {
			values.Set(k, x)
		}
	}
	return values
}

func (rb *ByProjectKeyShippingMethodsRequestMethodGet) Expand(v []string) *ByProjectKeyShippingMethodsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsRequestMethodGet) Sort(v []string) *ByProjectKeyShippingMethodsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsRequestMethodGet) Limit(v int) *ByProjectKeyShippingMethodsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyShippingMethodsRequestMethodGet) Offset(v int) *ByProjectKeyShippingMethodsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyShippingMethodsRequestMethodGet) WithTotal(v bool) *ByProjectKeyShippingMethodsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyShippingMethodsRequestMethodGet) Where(v []string) *ByProjectKeyShippingMethodsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyShippingMethodsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsRequestMethodGet) WithQueryParams(input ByProjectKeyShippingMethodsRequestMethodGetInput) *ByProjectKeyShippingMethodsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyShippingMethodsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyShippingMethodsRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyShippingMethodsRequestMethodGet) Execute(ctx context.Context) (result *ShippingMethodPagedQueryResponse, err error) {
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
