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

type ByProjectKeyStatesKeyByKeyRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyStatesKeyByKeyRequestMethodDeleteInput
}

func (r *ByProjectKeyStatesKeyByKeyRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyStatesKeyByKeyRequestMethodDeleteInput struct {
	Version int
	Expand  []string
}

func (input *ByProjectKeyStatesKeyByKeyRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyStatesKeyByKeyRequestMethodDelete) Version(v int) *ByProjectKeyStatesKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyStatesKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyStatesKeyByKeyRequestMethodDelete) Expand(v []string) *ByProjectKeyStatesKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyStatesKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyStatesKeyByKeyRequestMethodDelete) WithQueryParams(input ByProjectKeyStatesKeyByKeyRequestMethodDeleteInput) *ByProjectKeyStatesKeyByKeyRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyStatesKeyByKeyRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyStatesKeyByKeyRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyStatesKeyByKeyRequestMethodDelete) Execute(ctx context.Context) (result *State, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.delete(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		nil,
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
	case 409, 400, 401, 403, 500, 502, 503:
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
