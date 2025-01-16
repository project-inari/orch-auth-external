package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/project-inari/orch-auth-external/dto"
	"github.com/project-inari/orch-auth-external/pkg/httpclient"
)

type coreAuthAPIRepository struct {
	baseURL        string
	signupPath     string
	deleteUserPath string
	client         *http.Client
}

// CoreAuthAPIRepositoryConfig represents the configuration for core auth API repository
type CoreAuthAPIRepositoryConfig struct {
	BaseURL        string
	SignupPath     string
	DeleteUserPath string
}

// CoreAuthAPIRepositoryDependencies represents the dependencies for core auth API repository
type CoreAuthAPIRepositoryDependencies struct {
	Client *http.Client
}

// NewCoreAuthAPIRepository creates a new instance of core auth API repository
func NewCoreAuthAPIRepository(c CoreAuthAPIRepositoryConfig, d CoreAuthAPIRepositoryDependencies) CoreAuthAPIRepository {
	return &coreAuthAPIRepository{
		baseURL:        c.BaseURL,
		signupPath:     c.SignupPath,
		deleteUserPath: c.DeleteUserPath,
		client:         d.Client,
	}
}

// SignUp registers a new user in the core-auth-server service to Firebase
func (r *coreAuthAPIRepository) SignUp(ctx context.Context, req dto.CoreAuthSignUpReq, h dto.CoreAuthSignUpReqHeader) (*httpclient.Response[dto.CoreAuthSignUpRes], error) {
	url := fmt.Sprintf("%s%s", r.baseURL, r.signupPath)
	return httpclient.Post[dto.CoreAuthSignUpReq, dto.CoreAuthSignUpRes](ctx, r.client, url, h.ToMap(), req)
}

// DeleteUser deletes a user in the core-auth-server service from Firebase
func (r *coreAuthAPIRepository) DeleteUser(ctx context.Context, req dto.CoreAuthDeleteUserReq) (*httpclient.Response[dto.CoreAuthDeleteUserRes], error) {
	url := fmt.Sprintf("%s%s", r.baseURL, r.deleteUserPath)
	return httpclient.Delete[dto.CoreAuthDeleteUserReq, dto.CoreAuthDeleteUserRes](ctx, r.client, url, map[string]string{}, req)
}
