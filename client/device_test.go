package client

import (
	"testing"

	"github.com/Marquis42/ari-proxy/internal/integration"
)

func TestDeviceData(t *testing.T) {
	integration.TestDeviceData(t, &srv{})
}

func TestDeviceDelete(t *testing.T) {
	integration.TestDeviceDelete(t, &srv{})
}

func TestDeviceUpdate(t *testing.T) {
	integration.TestDeviceUpdate(t, &srv{})
}

func TestDeviceList(t *testing.T) {
	integration.TestDeviceList(t, &srv{})
}
