// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// MobileAppIntentAndStateRequestBuilder is request builder for MobileAppIntentAndState
type MobileAppIntentAndStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppIntentAndStateRequest
func (b *MobileAppIntentAndStateRequestBuilder) Request() *MobileAppIntentAndStateRequest {
	return &MobileAppIntentAndStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppIntentAndStateRequest is request for MobileAppIntentAndState
type MobileAppIntentAndStateRequest struct{ BaseRequest }

// Get performs GET request for MobileAppIntentAndState
func (r *MobileAppIntentAndStateRequest) Get(ctx context.Context) (resObj *MobileAppIntentAndState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileAppIntentAndState
func (r *MobileAppIntentAndStateRequest) Update(ctx context.Context, reqObj *MobileAppIntentAndState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileAppIntentAndState
func (r *MobileAppIntentAndStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}