package models

type MachineRegisterRequest struct {
	MachineID string `json:"machine_id"`
	MachineIP string `json:"machine_ip"`
}
