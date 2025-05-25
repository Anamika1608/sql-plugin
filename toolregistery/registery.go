package toolregistry

import (
	"context"
	"fmt"
)

type client interface {
	InstallTool(ctx context.Context, name, version, script string) (path string, err error)
}

type Registry struct {
	client client
}

func NewRegistry(client client) *Registry {
	return &Registry{client: client}
}

func (r *Registry) SQLClient(ctx context.Context, version, dbType string) (string, error) {
	switch dbType {
	case "postgres":
		return r.client.InstallTool(ctx, "psql", version, installPostgresScript)
	case "mysql":
		return r.client.InstallTool(ctx, "mysql", version, installMySQLScript)
	default:
		return "", fmt.Errorf("unsupported dbType: %s", dbType)
	}
}