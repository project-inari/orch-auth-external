// Package repository provides the repository interfaces for the domain
package repository

import (
	"context"

	"github.com/orch-auth-external/dto"
	"github.com/orch-auth-external/pkg/httpclient"
)

// WiremockAPIRepository represents the repository layer functions of wiremock API repository
type WiremockAPIRepository interface {
	GetTest(ctx context.Context, h dto.WiremockGetTestHeader) (*httpclient.Response[dto.WiremockGetTestResponse], error)
}
