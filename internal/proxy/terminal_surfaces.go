package proxy

import (
	"context"

	"weave-os/router/internal/auth"
	"weave-os/router/internal/requestcontext"
)

func hideTerminalSurfacesForRequest(ctx context.Context) bool {
	if _, active := auth.BlindExperimentFrom(ctx); active {
		return true
	}
	if _, gateway := requestcontext.UpstreamGatewayFrom(ctx); gateway {
		return true
	}
	hide, _ := ctx.Value(InstallationHideTerminalSurfacesContextKey{}).(bool)
	return hide
}
