package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PhysicalInterface struct {
	Ifname        string                          `json:"ifname"`
	ID            string                          `json:"id"`
	Type          string                          `json:"type"`
	Mode          string                          `json:"mode"`
	Enabled       bool                            `json:"enabled"`
	Ipv4          []PhysicalInterfaceIpv4         `json:"ipv4"`
	Security_Zone []PhysicalInterfaceSecurityZone `json:"securityZone"`
	Description   string                          `json:"description"`
	Name          string                          `json:"name"`
}

type PhysicalInterfaceIpv4 struct {
	Static []PhysicalInterfaceIpv4Static `json:"static"`
	Dhcp   []PhysicalInterfaceIpv4Dhcp   `json:"dhcp"`
}
type PhysicalInterfaceSecurityZone struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type PhysicalInterfaceIpv4Static struct {
	Address string `json:"address"`
	Netmask string `json:"netmask"`
}

type PhysicalInterfaceIpv4Dhcp struct {
	EnableDefaultRouteDHCP string `json:"enableDefaultRouteDHCP"`
	DhcpRouteMetric        string `json:"dhcpRouteMetric"`
}

type PhysicalInterfaceResponse struct {
	Links struct {
		Self   string `json:"self"`
		Parent string `json:"parent"`
	} `json:"links"`
	Type        string `json:"type"`
	Ifname      string `json:"ifname"`
	Enabled     bool   `json:"enabled"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}

type PhysicalInterfacesResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []struct {
		ID     string `json:"id"`
		Ifname string `json:"ifname"`
		Type   string `json:"type"`
		Links  struct {
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

func (v *Client) GetFmcPhysicalInterface(ctx context.Context, deviceID string, name string) (*PhysicalInterfaceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/physicalinterfaces?expanded=true", v.domainBaseURL, deviceID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting physical interfaces: %s - %s", url, err.Error())
	}
	physicalInterfaces := &PhysicalInterfacesResponse{}
	err = v.DoRequest(req, physicalInterfaces, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("getting physical interfaces: %s - %s", url, err.Error())
	}

	for _, PhysicalInterface := range physicalInterfaces.Items {
		if PhysicalInterface.Name == name {
			return &PhysicalInterfaceResponse{
				ID:   PhysicalInterface.ID,
				Name: PhysicalInterface.Name,
				Type: PhysicalInterface.Type,
			}, nil
		}
	}

	return nil, fmt.Errorf("no Interface found with name %s", name)
}

func (v *Client) UpdateFmcPhysicalInterface(ctx context.Context, deviceID string, id string, object *PhysicalInterface) (*PhysicalInterfaceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/physicalinterfaces/%s", v.domainBaseURL, deviceID, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating physical interfaces: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating physical interfaces: %s - %s", url, err.Error())
	}
	item := &PhysicalInterfaceResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting physical interfaces: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcPhysicalInterface() {

}
