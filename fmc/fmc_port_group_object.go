package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PortGroupObjectUpdateInput struct {
	Name        string                   `json:"name"`
	Objects     []PortGroupObjectObjects `json:"objects"`
	Description string                   `json:"description"`
	Type        string                   `json:"type"`
	ID          string                   `json:"id"`
}

type PortGroupObjectObjects struct {
	Name string `json:"name"`
	Type string `json:"type"`
	ID   string `json:"id"`
}

type PortGroupObject struct {
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Objects     []PortGroupObjectObjects `json:"objects"`
	Type        string                   `json:"type"`
}

type PortGroupObjectResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Type        string                   `json:"type"`
	Objects     []PortGroupObjectObjects `json:"objects"`
	Overridable bool                     `json:"overridable"`
	Description string                   `json:"description"`
	Name        string                   `json:"name"`
	ID          string                   `json:"id"`
}

type PortGroupObjectsResponse struct {
	Items []PortGroupObjectResponse `json:"items"`
}

// /fmc_config/v1/domain/DomainUUID/object/portobjectgroups?bulk=true ( Bulk POST operation on port group objects. )

func (v *Client) CreateFmcPortGroupObject(ctx context.Context, object *PortGroupObject) (*PortGroupObjectResponse, error) {
	url := fmt.Sprintf("%s/object/portobjectgroups?limit=1000", v.domainBaseURL)
	urlCreate := fmt.Sprintf("%s/object/portobjectgroups", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("creating port group objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", urlCreate, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating port group objects: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &PortGroupObjectResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("getting port group objects: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) GetFmcPortGroupObject(ctx context.Context, id string) (*PortGroupObjectResponse, error) {
	url := fmt.Sprintf("%s/object/portobjectgroups/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting port group objects: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &PortGroupObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return item, fmt.Errorf("getting port group objects: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) GetFmcPortGroupObjectByName(ctx context.Context, name string) (*PortGroupObjectResponse, error) {
	url := fmt.Sprintf("%s/object/portobjectgroups", v.domainBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting port group object: %s - %s", url, err.Error())
	}
	portGroupObjects := &PortGroupObjectsResponse{}
	err = v.DoRequest(req, portGroupObjects, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("getting port group object: %s - %s", url, err.Error())
	}

	for _, PortGroupObjects := range portGroupObjects.Items {
		if PortGroupObjects.Name == name {
			return &PortGroupObjectResponse{
				ID:   PortGroupObjects.ID,
				Name: PortGroupObjects.Name,
				Type: PortGroupObjects.Type,
			}, nil
		}
	}

	return nil, fmt.Errorf("no port group  object found with name %s", name)
}

func (v *Client) UpdateFmcPortGroupObject(ctx context.Context, id string, object *PortGroupObjectUpdateInput) (*PortGroupObjectResponse, error) {
	url := fmt.Sprintf("%s/object/portobjectgroups/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating port group objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating port group objects: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &PortGroupObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting port group objects: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) DeleteFmcPortGroupObject(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/portobjectgroups/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting port group objects: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	err = v.DoRequest(req, nil, http.StatusOK)
	Log.line()
	return err
}