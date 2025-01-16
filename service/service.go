// Package service provides the business logic service layer for the server
package service

import (
	"context"

	"github.com/project-inari/orch-auth-external/dto"
	"github.com/project-inari/orch-auth-external/repository"
)

// Port represents the service layer functions
type Port interface {
	SignUp(ctx context.Context, req dto.SignUpReq, h dto.SignUpReqHeader) (*dto.SignUpRes, error)
}

type service struct {
	coreAuthAPIRepository repository.CoreAuthAPIRepository
	coreUserAPIRepository repository.CoreUserAPIRepository
}

// Dependencies represents the dependencies for the service
type Dependencies struct {
	CoreAuthAPIRepository repository.CoreAuthAPIRepository
	CoreUserAPIRepository repository.CoreUserAPIRepository
}

// New creates a new service
func New(d Dependencies) Port {
	return &service{
		coreAuthAPIRepository: d.CoreAuthAPIRepository,
		coreUserAPIRepository: d.CoreUserAPIRepository,
	}
}
