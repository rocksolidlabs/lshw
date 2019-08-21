package lshw

import (
	"encoding/json"

	"github.com/rocksolidlabs/RSIM/src/lib/sh"
)

type LSHW struct {
	Businfo       string      `json:"businfo,omitempty"`
	Capabilities  *LSHWCAP    `json:"capabilities,omitempty"`
	Capacity      int64       `json:"capacity,omitempty"`
	Children      []*LSHW     `json:"children,omitempty"`
	Claimed       bool        `json:"claimed,omitempty"`
	Class         string      `json:"class,omitempty"`
	Clock         int         `json:"clock,omitempty"`
	Configuration *LSHWCONFIG `json:"configuration,omitempty"`
	Date          string      `json:"date,omitempty"`
	Description   string      `json:"description,omitempty"`
	Dev           string      `json:"dev,omitempty"`
	Disabled      bool        `json:"disabled,omitempty"`
	Handle        string      `json:"handle,omitempty"`
	ID            string      `json:"id,omitempty"`
	Logicalname   interface{} `json:"logicalname,omitempty"`
	Physid        string      `json:"physid,omitempty"`
	Product       string      `json:"product,omitempty"`
	Serial        string      `json:"serial,omitempty"`
	Slot          string      `json:"slot,omitempty"`
	Size          int64       `json:"size,omitempty"`
	Units         string      `json:"units,omitempty"`
	Vendor        string      `json:"vendor,omitempty"`
	Version       string      `json:"version,omitempty"`
	Width         int         `json:"width,omitempty"`
}

type LSHWCONFIG map[string]interface{}
type LSHWCAP map[string]interface{}

func GetLSHW() (*LSHW, error) {
	// sudo lshw -json
	out, err := sh.Command("sudo", "lshw", "-json").Output()
	if err != nil {
		return nil, err
	}
	lshw := &LSHW{}
	err = json.Unmarshal(out, &lshw)
	if err != nil {
		return nil, err
	}
	return lshw, nil
}
