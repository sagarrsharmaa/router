package proxy

import (
	"context"
	"net/http"
	"testing"

	"weave-os/router/internal/requestcontext"

	"github.com/stretchr/testify/assert"
)

func TestUpstreamGatewayModeHidesTerminalSurfaces(t *testing.T) {
	ctx := requestcontext.WithUpstreamGateway(context.Background(), requestcontext.UpstreamGatewayHeaders{}, http.Header{})
	assert.True(t, hideTerminalSurfacesForRequest(ctx))
	assert.False(t, hideTerminalSurfacesForRequest(context.Background()))
}
