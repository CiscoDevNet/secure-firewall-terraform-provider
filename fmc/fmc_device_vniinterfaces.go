package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type FTDVNIInterfaceResponse struct {
	Metadata struct {
		Timestamp int `json:"timestamp"`
		Domain    struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		}
	} `json:"metadata"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	ID                             string                         `json:"id"`
	Name                           string                         `json:"name"`
	Type                           string                         `json:"type"`
	VNIID                          int                            `json:"vniId"`
	MulticastGroupAddress          string                         `json:"multicastGroupAddress"`
	VTEPID                         int                            `json:"vtepID"`
	Enabled                        bool                           `json:"enabled"`
	EnableProxy                    bool                           `json:"enableProxy"`
	SegmentId                      int                            `json:"segmentId"`
	IPv4                           IPv4                           `json:"ipv4"`
	IPv6                           IPv6_VNI                       `json:"ipv6"`
	EnableAntiSpoofing             bool                           `json:"enableAntiSpoofing"`
	OverrideDefaultFragmentSetting OverrideDefaultFragmentSetting `json:"overrideDefaultFragmentSetting"`
}
type Prefix struct {
	Address       string `json:"address"`
	Default       bool   `json:"default"`
	Advertisement struct {
		Offlink        bool `json:"offlink"`
		AutoConfig     bool `json:"autoConfig"`
		PreferLifeTime struct {
			Duration struct {
				PreferLifeTime string `json:"preferLifeTime"`
				ValidLifeTime  bool   `json:"validLifeTime"`
			} `json:"duration"`
		} `json:"preferLifeTime"`
	} `json:"advertisement"`
}

type Addresses struct {
	Address      string `json:"address"`
	Prefix       string `json:"prefix"`
	EnforceEUI64 bool   `json:"enforceEUI64"`
}

type FTDVNIInterfacesResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items  []FTDVNIInterfaceResponse `json:"items"`
	Paging struct {
		OffsetMTU int `json:"offset"`
		Limit     int `json:"limit"`
		Count     int `json:"count"`
		Pages     int `json:"pages"`
	} `json:"paging"`
}
type OverrideDefaultFragmentSetting struct {
	Size    int `json:"size"`
	Timeout int `json:"timeout"`
	Chain   int `json:"chain"`
}
type ArpConfig struct {
	IpAddress   string `json:"ipAddress"`
	MacAddress  string `json:"macAddress"`
	EnableAlias bool   `json:"enableAlias"`
}
type SecurityZone_VNI struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
type IPv4 struct {
	Static Static `json:"static"`
}
type Static struct {
	Address string `json:"address"`
	Netmask string `json:"netmask"`
}
type IPv6_VNI struct {
	Prefixes                []Prefix    `json:"prefixes"`
	EnableIPV6              bool        `json:"enableIPV6"`
	LinkLocalAddress        string      `json:"linkLocalAddress"`
	RaLifeTime              int         `json:"raLifeTime"`
	RaInterval              int         `json:"raInterval"`
	Addresses               []Addresses `json:"addresses"`
	EnableRA                bool        `json:"enableRA"`
	EnforceEUI64            bool        `json:"enforceEUI64"`
	EnableAutoConfig        bool        `json:"enableAutoConfig"`
	EnableDHCPAddrConfig    bool        `json:"enableDHCPAddrConfig"`
	EnableDHCPNonAddrConfig bool        `json:"enableDHCPNonAddrConfig"`
	DadAttempts             int         `json:"dadAttempts"`
	NsInterval              int         `json:"nsInterval"`
	ReachableTime           int         `json:"reachableTime"`
}
type FmcAccessConfig struct {
	AllowedNetworks []AllowedNetworks `json:"allowedNetworks"`
	EnableAccess    bool              `json:"enableAccess"`
}
type AllowedNetworks struct {
	Overridable bool      `json:"overridable"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Overrides   Overrides `json:"overrides"`
}
type Overrides struct {
	Description string     `json:"description"`
	Parent      IReference `json:"parent"`
	Target      IReference `json:"target"`
}
type IReference struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	ID          string `json:"id"`
	Description string `json:"description"`
}
type FTDVNIInterfaces struct {
	Name                           string                         `json:"name"`
	Type                           string                         `json:"type"`
	VNIID                          int                            `json:"vniId"`
	MulticastGroupAddress          string                         `json:"multicastGroupAddress"`
	VTEPID                         int                            `json:"vtepID"`
	Enabled                        bool                           `json:"enabled"`
	EnableAntiSpoofing             bool                           `json:"enableAntiSpoofing"`
	SegmentId                      int                            `json:"segmentId"`
	Ifname                         string                         `json:"ifname"`
	EnableProxy                    bool                           `json:"enableProxy"`
	OverrideDefaultFragmentSetting OverrideDefaultFragmentSetting `json:"overrideDefaultFragmentSetting"`
	ArpConfig                      []ArpConfig                    `json:"arpConfig"`
	SecurityZone                   SecurityZone_VNI               `json:"securityZone"`
	IPv4                           IPv4                           `json:"ipv4"`
	IPv6                           IPv6_VNI                       `json:"ipv6"`
	FmcAccessConfig                FmcAccessConfig                `json:"fmcAccessConfig"`
}

// /api/fmc_config/v1/domain/{domainUUID}/devices/devicerecords/{containerUUID}/vniinterfaces/{objectId}
// func (v *Client) GetVNIInterfacesByName(ctx context.Context, id, name string) (*FTDVNIInterfaceResponse, error) {
// 	url := fmt.Sprintf("%s/devices/devicerecords/%s/vniinterfaces?expanded=false&filter=name:", v.domainBaseURL, id)
// 	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting vniinterfaces by name: %s - %s", url, err.Error())
// 	}
// 	resp := &FTDVNIInterfacesResponse{}
// 	err = v.DoRequest(req, resp, http.StatusOK)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting vniinterfaces by name: %s - %s", url, err.Error())
// 	}
// 	switch l := len(resp.Items); {
// 	case l == 1:
// 		return v.GetVNIInterfaces(ctx, resp.Items[0].ID)
// 	case l > 1:
// 		for _, item := range resp.Items {
// 			if item.Name == name {
// 				return v.GetVNIInterfaces(ctx, item.ID)
// 			}
// 		}
// 		return nil, fmt.Errorf("duplicates found, no exact match, length of response is: %d, expected 1, please search using a unique id, name or value", l)
// 	case l == 0:
// 		return nil, fmt.Errorf("no vni interface found, length of response is: %d, expected 1, please check your filter", l)
// 	}
// 	return nil, fmt.Errorf("no vni interface found found with name %s", name)
// }

// /api/fmc_config/v1/domain/{domainUUID}/devices/devicerecords/{containerUUID}/vniinterfaces
func (v *Client) GetVNIInterfaces(ctx context.Context, id string, name string) (*FTDVNIInterfaceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/vniinterfaces", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting vniinterfaces by name: %s - %s", url, err.Error())
	}
	resp := &FTDVNIInterfacesResponse{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting vniinterfaces by name: %s - %s", url, err.Error())
	}

	for _, vniInterface := range resp.Items {
		if vniInterface.Name == name {
			return &FTDVNIInterfaceResponse{
				ID:   vniInterface.ID,
				Name: vniInterface.Name,
				Type: vniInterface.Type,
			}, nil
		}
	}

	return nil, fmt.Errorf("no vniinterfaces found with name %s",name)
}

func (v *Client) CreateVNIInterfaces(ctx context.Context, acpId string, ftdVNIInterfaces *FTDVNIInterfaces) (*FTDVNIInterfaceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/vniinterfaces", v.domainBaseURL, acpId)
	body, err := json.Marshal(&ftdVNIInterfaces)
	if err != nil {
		return nil, fmt.Errorf("creating vni interfaces: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating vni interfaces: %s - %s", url, err.Error())
	}
	item := &FTDVNIInterfaceResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("creating vni interfaces: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateVNIInterfaces(ctx context.Context, acpId string, ftdVNIInterfaces *FTDVNIInterfaces) (*FTDVNIInterfaceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/vniinterfaces", v.domainBaseURL, acpId)
	body, err := json.Marshal(&ftdVNIInterfaces)
	if err != nil {
		return nil, fmt.Errorf("updating vni interfaces: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating vni interfaces: %s - %s", url, err.Error())
	}
	item := &FTDVNIInterfaceResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("updating vni interfaces: %s - %s,%+v", url, err.Error(), ftdVNIInterfaces)
	}
	return item, nil
}

func (v *Client) DeleteVNIInterfaces(ctx context.Context, acpId string) error {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/vniinterfaces", v.domainBaseURL, acpId)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting vni interfaces: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
