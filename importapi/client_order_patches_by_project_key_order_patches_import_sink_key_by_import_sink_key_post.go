// Generated file, please do not change!!!
package importapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyRequestMethodPost struct {
	body    OrderPatchImportRequest
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Creates a new import request for order patches
 */
func (rb *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyRequestMethodPost) Execute(ctx context.Context) (result *ImportResponse, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	queryParams := url.Values{}
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
	case 201:
		err = json.Unmarshal(content, &result)
		return result, nil
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
