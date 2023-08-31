package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type NetworkGroupObjectUpdateInput struct {
	Name        string                       `json:"name"`
	Objects     []NetworkGroupObjectObjects  `json:"objects"`
	Literals    []NetworkGroupObjectLiterals `json:"literals"`
	Description string                       `json:"description"`
	Type        string                       `json:"type"`
	ID          string                       `json:"id"`
}

type NetworkGroupObjectObjects struct {
	Name string `json:"name"`
	Type string `json:"type"`
	ID   string `json:"id"`
}

type NetworkGroupObjectLiterals struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type NetworkGroupObject struct {
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Objects     []NetworkGroupObjectObjects  `json:"objects"`
	Literals    []NetworkGroupObjectLiterals `json:"literals"`
	Type        string                       `json:"type"`
}

type NetworkGroupObjectResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Type        string                       `json:"type"`
	Literals    []NetworkGroupObjectLiterals `json:"literals"`
	Objects     []NetworkGroupObjectObjects  `json:"objects"`
	Overridable bool                         `json:"overridable"`
	Name        string                       `json:"name"`
	ID          string                       `json:"id"`
	Description string                       `json:"description"`
	Metadata    struct {
		Timestamp int `json:"timestamp"`
		Lastuser  struct {
			Name string `json:"name"`
		} `json:"lastUser"`
		Domain struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"domain"`
	} `json:"metadata"`
}
type NetworkGroupObjectsResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Name string `json:"name"`
	} `json:"items"`
	Paging struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Count  int `json:"count"`
		Pages  int `json:"pages"`
	} `json:"paging"`
}

// /fmc_config/v1/domain/DomainUUID/object/networkgroups?bulk=true ( Bulk POST operation on network objects. )
func (v *Client) GetFmcNetworkGroupObjectByName(ctx context.Context, name string) (*NetworkGroupObjectResponse, error) {
	url := fmt.Sprintf("%s/object/networkgroups?filter=nameOrValue:%s", v.domainBaseURL, name)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting network group object: %s - %s", url, err.Error())
	}
	networkGroupObjects := &NetworkGroupObjectsResponse{}
	err = v.DoRequest(req, networkGroupObjects, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("getting network group object: %s - %s", url, err.Error())
	}

	for _, NetworkGroupObjects := range networkGroupObjects.Items {
		if NetworkGroupObjects.Name == name {
			return &NetworkGroupObjectResponse{
				ID:   NetworkGroupObjects.ID,
				Name: NetworkGroupObjects.Name,
				Type: NetworkGroupObjects.Type,
			}, nil
		}
	}

	return nil, fmt.Errorf("no network group found with name %s", name)
}
func (v *Client) CreateFmcNetworkGroupObject(ctx context.Context, object *NetworkGroupObject) (*NetworkGroupObjectResponse, error) {
	url := fmt.Sprintf("%s/object/networkgroups", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("creating network objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating network objects: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &NetworkGroupObjectResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("getting network objects: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) GetFmcNetworkGroupObject(ctx context.Context, id string) (*NetworkGroupObjectResponse, error) {
	url := fmt.Sprintf("%s/object/networkgroups/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting network objects: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &NetworkGroupObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return item, fmt.Errorf("getting network objects: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) UpdateFmcNetworkGroupObject(ctx context.Context, id string, object *NetworkGroupObjectUpdateInput) (*NetworkGroupObjectResponse, error) {
	url := fmt.Sprintf("%s/object/networkgroups/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating network objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating network objects: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &NetworkGroupObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting network objects: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) DeleteFmcNetworkGroupObject(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/networkgroups/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting network objects: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	err = v.DoRequest(req, nil, http.StatusOK)
	Log.line()
	return err
}