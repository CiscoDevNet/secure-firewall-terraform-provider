package fmc

import (
	"context"
	"fmt"
	"net/http"
)

type AnyconnectObj struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PayloadFile string `json:"payloadfile"`
}

type AnyconnectObjResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	FileName    string `json:"filename"`
	Type        string `json:"type"`
}

type AnyconnectObjsResponse struct {
	Items []struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	}
}

func (v *Client) GetFmcAnyconnectPackageByNameOrValue(ctx context.Context, nameOrValue string) (*AnyconnectObjResponse, error) {
	url := fmt.Sprintf("%s/object/anyconnectpackages?filter=nameOrValue:%s", v.domainBaseURL, nameOrValue)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting anyconnect object by name/value: %s - %s", url, err.Error())
	}

	resp := &AnyconnectObjsResponse{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting anyconnect object by name/value: %s - %s", url, err.Error())
	}

	switch l := len(resp.Items); {
	case l == 1:
		return v.GetFmcAnyconnectPackageObject(ctx, resp.Items[0].ID)
	case l > 1:
		for _, item := range resp.Items {
			if item.Name == nameOrValue {
				return v.GetFmcAnyconnectPackageObject(ctx, item.ID)
			}
		}
		return nil, fmt.Errorf("duplicates found, no exact match, length of response is: %d, expected 1, please search using a unique id, name or value", l)
	case l == 0:
		return nil, fmt.Errorf("no anyconnect objects found, length of response is: %d, expected 1, please check your filter", l)
	}
	return nil, fmt.Errorf("this should not be reachable, this is a bug")
}

func (v *Client) GetFmcAnyconnectPackageObject(ctx context.Context, id string) (*AnyconnectObjResponse, error) {
	url := fmt.Sprintf("%s/object/anyconnectpackages/:%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting anyconnect objects: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")

	item := &AnyconnectObjResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return item, fmt.Errorf("getting anyconnect objects: %s - %s", url, err.Error())
	}

	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) CreateFmcAnyconnectPackageObject(ctx context.Context, object *AnyconnectObj) (*AnyconnectObjResponse, error) {
	return nil, nil
}
