// Package repository provides the repository interfaces for the domain
package repository

import (
	"context"

	"github.com/project-inari/orch-auth-external/dto"
	"github.com/project-inari/orch-auth-external/pkg/httpclient"
)

// CoreAuthAPIRepository represents the repository layer functions of core auth API repository
type CoreAuthAPIRepository interface {
	CallSignUp(ctx context.Context, req dto.CoreAuthSignUpReq, h dto.CoreAuthSignUpReqHeader) (*httpclient.Response[dto.CoreAuthSignUpRes], error)
	CallDeleteUser(ctx context.Context, req dto.CoreAuthDeleteUserReq) (*httpclient.Response[dto.CoreAuthDeleteUserRes], error)
}

// CoreUserAPIRepository represents the repository layer functions of core user API repository
type CoreUserAPIRepository interface {
	CallSignUp(ctx context.Context, req dto.CoreUserSignUpReq) (*httpclient.Response[dto.CoreUserSignUpRes], error)
}
