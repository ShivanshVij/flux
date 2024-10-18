package models

type DiscoveryData struct {
	MachineName     string `json:"MachineName"`     // Machine Name
	MachineModel    string `json:"MachineModel"`    // Machine Model
	BrandName       string `json:"BrandName"`       // Brand Name
	MachineIP       string `json:"MachineIP"`       // Motherboard IP Address
	MachineID       string `json:"MachineID"`       // Motherboard ID (16-bit)
	ProtocolVersion string `json:"ProtocolVersion"` // Protocol Version
	FirmwareVersion string `json:"FirmwareVersion"` // Firmware Version
}

type DiscoveryResponse struct {
	Discovered []*DiscoveryData `json:"discovered"`
}
