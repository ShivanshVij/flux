package models

type DiscoveryData struct {
	MachineName     string `json:"MachineName"`     // Machine Name
	MachineModel    string `json:"MachineModel"`    // Machine Model
	BrandName       string `json:"BrandName"`       // Brand Name
	MainboardIP     string `json:"MainboardIP"`     // Motherboard IP Address
	MainboardID     string `json:"MainboardID"`     // Motherboard ID (16-bit)
	ProtocolVersion string `json:"ProtocolVersion"` // Protocol Version
	FirmwareVersion string `json:"FirmwareVersion"` // Firmware Version
}

type DiscoveryResponse struct {
	Discovered []*DiscoveryData `json:"discovered"`
}
