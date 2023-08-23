package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SubInterfaceSecurityZone struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type SubInterfacesResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []SubInterfaceResponse `json:"items"`
}

type SubInterfaceRequest struct {
	ID             string                   `json:"id"`
	Type           string                   `json:"type"`
	Name           string                   `json:"name"`
	Description    string                   `json:"description"`
	VLANId         int                      `json:"vlanId"`
	SubInterfaceId int                      `json:"subIntfId"`
	MTU            int                      `json:"MTU"`
	Priority       int                      `json:"priority"`
	ManagementOnly bool                     `json:"managementOnly"`
	Ifname         string                   `json:"ifname"`
	Enabled        bool                     `json:"enabled"`
	SecurityZone   SubInterfaceSecurityZone `json:"securityZone"`
	IPv4           IPv4                     `json:"ipv4,omitempty"`
}

type SubInterfaceResponse struct {
	Type           string                   `json:"type"`
	ID             string                   `json:"id"`
	Name           string                   `json:"name"`
	Description    string                   `json:"description"`
	VLANId         int                      `json:"vlanId"`
	SubInterfaceId int                      `json:"subIntfId"`
	MTU            int                      `json:"MTU"`
	Enabled        bool                     `json:"enabled"`
	ManagementOnly bool                     `json:"managementOnly"`
	Mode           string                   `json:"mode"`
	Ifname         string                   `json:"ifname"`
	SecurityZone   SubInterfaceSecurityZone `json:"securityZone"`
	Priority       int                      `json:"priority"`
}

/*
*
This method return the subinterfaces ID based on device-id and subinterfaces Name
*/
func (v *Client) GetFmcSubInterface(ctx context.Context, deviceID string, name string) (*SubInterfaceResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/subinterfaces?expanded=true", v.domainBaseURL, deviceID)
	log.Printf("GET FMC SUBInterface based on deviceid=%s and name=%s", deviceID, name)
	log.Printf("Get FMC SUBInterface URL=%s", url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("error while creating context for getting subinterfacess: %s - %s", url, err.Error())
	}

	subintrfs := &SubInterfacesResponse{}
	err = v.DoRequest(req, subintrfs, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("error while getting SUBInterface: %s - %s", url, err.Error())
	}
	for _, subintf := range subintrfs.Items {
		log.Printf("SUBInterface.ID=%s  SUBInterface.Name=%s  SUBInterface.Type=%s", subintf.ID, subintf.Name, subintf.Type)

		if subintf.Name == name {
			return &SubInterfaceResponse{
				ID:   subintf.ID,
				Name: subintf.Name,
			}, nil
		}
	}
	log.Printf("no SUBInterface found with name %s", name)
	return nil, fmt.Errorf("no SUBInterface found with name %s", name)
}

/*
*
This method return the all subinterfaces details based on subinterfaces ID
*/
func (v *Client) GetFmcSubInterfaceDetails(ctx context.Context, deviceID string, subintfId string) (*SubInterfaceResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/subinterfaces/%s", v.domainBaseURL, deviceID, subintfId)

	log.Printf("GET FMC SUBInterface details based on deviceid=%s and subinterfaces ID=%s", deviceID, subintfId)
	log.Printf("Get FMC SUBInterface details URL=%s", url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("error while creating context for getting SUBInterface details: %s - %s", url, err.Error())
	}

	subintid := &SubInterfaceResponse{}
	err = v.DoRequest(req, subintid, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("error while getting subinterfaces by id: %s - %s", url, err.Error())
	}

	return &SubInterfaceResponse{
		ID:             subintid.ID,
		Name:           subintid.Name,
		Ifname:         subintid.Ifname,
		Type:           subintid.Type,
		Description:    subintid.Description,
		VLANId:         subintid.VLANId,
		Enabled:        subintid.Enabled,
		SubInterfaceId: subintid.SubInterfaceId,
		MTU:            subintid.MTU,
		ManagementOnly: subintid.ManagementOnly,
		Mode:           subintid.Mode,
		SecurityZone:   subintid.SecurityZone,
		Priority:       subintid.Priority,
	}, nil

}

/*
*
This method creates the subinterfaces
*/

func (v *Client) CreateFmcSubInterface(ctx context.Context, deviceID string, object *SubInterfaceRequest) (*SubInterfaceResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/subinterfaces", v.domainBaseURL, deviceID)

	log.Printf("Create subinterfaces URL=%s", url)

	body, err := json.Marshal(&object)

	log.Printf("Create subinterfaces Request=%s", body)

	if err != nil {
		return nil, fmt.Errorf("invalid request json : %s - %s", url, err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))

	if err != nil {
		return nil, fmt.Errorf("error while creating context for subinterfaces: %s - %s", url, err.Error())
	}

	response := &SubInterfaceResponse{}
	err = v.DoRequest(req, response, http.StatusCreated)

	if err != nil {
		return nil, fmt.Errorf("error while cretaing subinterfaces: %s - %s", url, err.Error())
	}

	log.Printf("subinterfaces interface created, response=%s", response.ID)

	return response, nil
}

/*
*
This method update the subinterfaces details
*/

func (v *Client) UpdateFmcSubInterface(ctx context.Context, deviceID string, id string, object *SubInterfaceRequest) (*SubInterfaceResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/subinterfaces/%s", v.domainBaseURL, deviceID, id)

	log.Printf("Update subinterfaces URL=%s", url)

	body, err := json.Marshal(&object)

	log.Printf("Update subinterfaces Request=%s", body)

	if err != nil {
		return nil, fmt.Errorf("invalid request json : %s - %s", url, err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))

	if err != nil {
		return nil, fmt.Errorf("error while Updating context for subinterfaces: %s - %s", url, err.Error())
	}

	response := &SubInterfaceResponse{}
	err = v.DoRequest(req, response, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("error while Updating subinterfaces: %s - %s", url, err.Error())
	}

	log.Printf("subinterfaces Update, response=%s", response.ID)

	return response, nil
}

/*
*
This method delete the subinterfaces based in subinterfaces ID
*/
func (v *Client) DeleteFmcSubInterfaceDetails(ctx context.Context, deviceID string, id string) (*SubInterfaceResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/subinterfaces/%s", v.domainBaseURL, deviceID, id)

	log.Printf("Delete FMC subinterfaces details based on deviceid=%s and subinterfaces ID=%s", deviceID, id)
	log.Printf("Delete FMC subinterfaces details URL=%s", url)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)

	if err != nil {
		return nil, fmt.Errorf("error while creating context for Delete subinterfaces details: %s - %s", url, err.Error())
	}

	response := &SubInterfaceResponse{}
	err = v.DoRequest(req, response, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("error while Deleting subinterfaces by id: %s - %s", url, err.Error())
	}

	log.Printf("subinterfaces deleted, response=%s", response.ID)

	return response, nil
}
