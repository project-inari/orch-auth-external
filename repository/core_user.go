package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/project-inari/orch-auth-external/dto"
	"github.com/project-inari/orch-auth-external/pkg/httpclient"
)

type coreUserAPIRepository struct {
	baseURL    string
	signupPath string
	client     *http.Client
}

// CoreUserAPIRepositoryConfig represents the configuration for core user API repository
type CoreUserAPIRepositoryConfig struct {
	BaseURL    string
	SignupPath string
}

// CoreUserAPIRepositoryDependencies represents the dependencies for core auth API repository
type CoreUserAPIRepositoryDependencies struct {
	Client *http.Client
}

// NewCoreUserAPIRepository creates a new instance of core user API repository
func NewCoreUserAPIRepository(c CoreUserAPIRepositoryConfig, d CoreUserAPIRepositoryDependencies) CoreUserAPIRepository {
	return &coreUserAPIRepository{
		baseURL:    c.BaseURL,
		signupPath: c.SignupPath,
		client:     d.Client,
	}
}

// SignUp registers a new user in the core-auth-server service to Firebase
func (r *coreUserAPIRepository) SignUp(ctx context.Context, req dto.CoreUserSignUpReq) (*httpclient.Response[dto.CoreUserSignUpRes], error) {
	url := fmt.Sprintf("%s%s", r.baseURL, r.signupPath)
	return httpclient.Post[dto.CoreUserSignUpReq, dto.CoreUserSignUpRes](ctx, r.client, url, map[string]string{}, req)
}
