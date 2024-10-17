package sdcp

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiscovery(t *testing.T) {
	discovery, err := Discover(context.Background())
	require.NoError(t, err)
	if len(discovery) == 0 {
		t.Log("no discovery messages received")
	}
	for _, d := range discovery {
		t.Logf("ID %s: Brand Name '%s'", d.ID, d.Data.BrandName)
		t.Logf("ID %s: MachineModel '%s'", d.ID, d.Data.MachineModel)
		t.Logf("ID %s: MainboardID '%s'", d.ID, d.Data.MainboardID)
		t.Logf("ID %s: MachineName '%s'", d.ID, d.Data.MachineName)
		t.Logf("ID %s: MainboardIP '%s'", d.ID, d.Data.MainboardIP)
		t.Logf("ID %s: FirmwareVersion '%s'", d.ID, d.Data.FirmwareVersion)
		t.Logf("ID %s: ProtocolVersion '%s'", d.ID, d.Data.ProtocolVersion)
	}
}
