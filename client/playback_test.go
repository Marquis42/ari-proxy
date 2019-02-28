package client

import (
	"testing"

	"github.com/Marquis42/ari-proxy/internal/integration"
)

func TestPlaybackData(t *testing.T) {
	integration.TestPlaybackData(t, &srv{})
}

func TestPlaybackControl(t *testing.T) {
	integration.TestPlaybackControl(t, &srv{})
}

func TestPlaybackStop(t *testing.T) {
	integration.TestPlaybackStop(t, &srv{})
}
