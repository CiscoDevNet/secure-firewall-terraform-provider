package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SourceInterface struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type NveNeighborAddress struct {
	Literal struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"literal"`
	Object struct {
		Overridable bool   `json:"overridable"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		Type        string `json:"type"`
	} `json:"object"`
}
type VTEPEntry struct {
	SourceInterface          SourceInterface    `json:"sourceInterface"`
	NveNeighborAddress       NveNeighborAddress `json:"nveNeighborAddress"`
	NveVtepId                int                `json:"nveVtepId"`
	NveDestinationPort       int                `json:"nveDestinationPort"`
	NveEncapsulationType     string             `json:"nveEncapsulationType"`
	// NveNeighborDiscoveryType string             `json:"nveNeighborDiscoveryType"`
}
type VTEPPolicyResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	ID          string      `json:"id"`
	// Name        string      `json:"name"`
	NveEnable   bool        `json:"nveEnable"`
	VTEPEntries []VTEPEntry `json:"vtepEntries"`
	Type        string      `json:"type"`
}

type VTEPPoliciesResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items  []VTEPPolicyResponse `json:"items"`
	Paging struct {
		OffsetMTU int `json:"offset"`
		Limit     int `json:"limit"`
		Count     int `json:"count"`
		Pages     int `json:"pages"`
	} `json:"paging"`
}

type VTEPPolicy struct {
	NveEnable   bool        `json:"nveEnable"`
	VTEPEntries []VTEPEntry `json:"vtepEntries"`
	Type        string      `json:"type"`
}

// // /api/fmc_config/v1/domain/{domainUUID}/devices/devicerecords/{containerUUID}/vteppolicies/{objectId}
// func (v *Client) GetVTEPPoliciesByName(ctx context.Context, id string) (*VTEPPolicyResponse, error) {
// 	url := fmt.Sprintf("%s/devices/devicerecords/%s/vteppolicies", v.domainBaseURL, id)
// 	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting vteppolicies by name: %s - %s", url, err.Error())
// 	}
// 	resp := &VTEPPoliciesResponse{}
// 	err = v.DoRequest(req, resp, http.StatusOK)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting vteppolicies by name: %s - %s", url, err.Error())
// 	}
// 	switch l := len(resp.Items); {
// 	case l == 1:
// 		return v.GetVTEPPolicies(ctx, resp.Items[0].ID)
// 	case l > 1:
// 		for _, item := range resp.Items {
// 			if item.Name == name {
// 				return v.GetVTEPPolicies(ctx, item.ID)
// 			}
// 		}
// 		return nil, fmt.Errorf("duplicates found, no exact match, length of response is: %d, expected 1, please search using a unique id, name or value", l)
// 	case l == 0:
// 		return nil, fmt.Errorf("no access policies found, length of response is: %d, expected 1, please check your filter", l)
// 	}
// 	return nil, fmt.Errorf("no vteppolicies found with name %s", name)
// }

// /api/fmc_config/v1/domain/{domainUUID}/devices/devicerecords/{containerUUID}/vteppolicies
func (v *Client) GetVTEPPolicies(ctx context.Context, deviceID string) (*VTEPPolicyResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/vteppolicies", v.domainBaseURL, deviceID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting vteppolicies: %s - %s", url, err.Error())
	}
	resp := &VTEPPoliciesResponse{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting vteppolicies: %s - %s", url, err.Error())
	}
	for _, vtepPolicy := range resp.Items {
			return &VTEPPolicyResponse{
				ID:   vtepPolicy.ID,
				Type: vtepPolicy.Type,
			}, nil
	}
	//return resp, fmt.Errorf("no vteppolicies found for the device-ID%s", deviceID)
	return nil, fmt.Errorf("no vteppolicies found for the device-ID%s", deviceID)
}

func (v *Client) CreateVTEPPolicies(ctx context.Context, acpId string, vtepPolicy *VTEPPolicy) (*VTEPPolicyResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/vteppolicies", v.domainBaseURL, acpId)
	body, err := json.Marshal(&vtepPolicy)
	if err != nil {
		return nil, fmt.Errorf("creating vtep policies: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating vtep policies: %s - %s", url, err.Error())
	}
	item := &VTEPPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("creating vtep policies: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateVTEPPolicies(ctx context.Context, acpId string, vtepPolicy *VTEPPolicy) (*VTEPPolicyResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/vteppolicies", v.domainBaseURL, acpId)
	body, err := json.Marshal(&vtepPolicy)
	if err != nil {
		return nil, fmt.Errorf("updating vtep policies: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating vtep policies: %s - %s", url, err.Error())
	}
	item := &VTEPPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("updating vtep policies: %s - %s,%+v", url, err.Error(), vtepPolicy)
	}
	return item, nil
}

func (v *Client) DeleteVTEPPolicies(ctx context.Context, acpId string) error {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/vteppolicies", v.domainBaseURL, acpId)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting vtep policies: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}