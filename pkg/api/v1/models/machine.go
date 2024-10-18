package models

import "github.com/shivanshvij/flux/pkg/sdcp"

type MachineRegisterRequest struct {
	MachineID string `json:"machine_id"`
	MachineIP string `json:"machine_ip"`
}

type MachineStatusResponse struct {
	Status sdcp.Status `json:"status"`
}
