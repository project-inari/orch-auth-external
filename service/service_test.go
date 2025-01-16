package service

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/project-inari/orch-auth-external/dto"
	"github.com/project-inari/orch-auth-external/pkg/httpclient"
	"github.com/stretchr/testify/assert"
)

type mockCoreAuthAPIRepository struct {
	signupResp     *httpclient.Response[dto.CoreAuthSignUpRes]
	deleteUserResp *httpclient.Response[dto.CoreAuthDeleteUserRes]
	err            error
}

func (m *mockCoreAuthAPIRepository) CallSignUp(_ context.Context, _ dto.CoreAuthSignUpReq, _ dto.CoreAuthSignUpReqHeader) (*httpclient.Response[dto.CoreAuthSignUpRes], error) {
	return m.signupResp, m.err
}

func (m *mockCoreAuthAPIRepository) CallDeleteUser(_ context.Context, _ dto.CoreAuthDeleteUserReq) (*httpclient.Response[dto.CoreAuthDeleteUserRes], error) {
	return m.deleteUserResp, m.err
}

type mockCoreUserAPIRepository struct {
	signupResp *httpclient.Response[dto.CoreUserSignUpRes]
	err        error
}

func (m *mockCoreUserAPIRepository) CallSignUp(_ context.Context, _ dto.CoreUserSignUpReq) (*httpclient.Response[dto.CoreUserSignUpRes], error) {
	return m.signupResp, m.err
}

const (
	mockFirstName = "firstName"
	mockLastName  = "lastName"
	mockEmail     = "email"
	mockPhoneNo   = "phoneNo"
	mockUsername  = "username"
	mockUID       = "uid"
	mockToken     = "token"
	mockPassword  = "password"

	mockENAcceptLocale = "EN"
)

func TestSignUp(t *testing.T) {
	ctx := context.Background()

	req := dto.SignUpReq{
		Username:  mockUsername,
		Email:     mockEmail,
		Password:  mockPassword,
		PhoneNo:   mockPhoneNo,
		FirstName: mockFirstName,
		LastName:  mockLastName,
	}

	h := dto.SignUpReqHeader{
		AcceptLocale: mockENAcceptLocale,
	}

	t.Run("success", func(t *testing.T) {
		authSignUpRes := &httpclient.Response[dto.CoreAuthSignUpRes]{
			HTTPStatusCode: http.StatusOK,
			Response: dto.CoreAuthSignUpRes{
				Username: mockUsername,
				UID:      mockUID,
				Token:    mockToken,
			},
		}

		userSignUpRes := &httpclient.Response[dto.CoreUserSignUpRes]{
			HTTPStatusCode: http.StatusOK,
			Response: dto.CoreUserSignUpRes{
				Username: mockUsername,
				UID:      mockUID,
				Success:  true,
			},
		}

		coreAuthAPIRepository := &mockCoreAuthAPIRepository{
			signupResp: authSignUpRes,
			err:        nil,
		}

		coreUserAPIRepository := &mockCoreUserAPIRepository{
			signupResp: userSignUpRes,
			err:        nil,
		}

		s := New(Dependencies{
			CoreAuthAPIRepository: coreAuthAPIRepository,
			CoreUserAPIRepository: coreUserAPIRepository,
		})

		res, err := s.SignUp(ctx, req, h)

		assert.Nil(t, err)
		assert.Equal(t, mockUsername, res.Username)
		assert.Equal(t, mockUID, res.UID)
		assert.Equal(t, mockToken, res.Token)
	})

	t.Run("error - when core auth CallSignUp http status code not 200", func(t *testing.T) {
		authSignUpRes := &httpclient.Response[dto.CoreAuthSignUpRes]{
			HTTPStatusCode: http.StatusInternalServerError,
			Response:       dto.CoreAuthSignUpRes{},
		}

		coreAuthAPIRepository := &mockCoreAuthAPIRepository{
			signupResp: authSignUpRes,
			err:        nil,
		}

		coreUserAPIRepository := &mockCoreUserAPIRepository{
			signupResp: nil,
			err:        nil,
		}

		s := New(Dependencies{
			CoreAuthAPIRepository: coreAuthAPIRepository,
			CoreUserAPIRepository: coreUserAPIRepository,
		})

		res, err := s.SignUp(ctx, req, h)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("error - when core auth CallSignUp client error", func(t *testing.T) {
		coreAuthAPIRepository := &mockCoreAuthAPIRepository{
			signupResp: nil,
			err:        errors.New("error"),
		}

		coreUserAPIRepository := &mockCoreUserAPIRepository{
			signupResp: nil,
			err:        nil,
		}

		s := New(Dependencies{
			CoreAuthAPIRepository: coreAuthAPIRepository,
			CoreUserAPIRepository: coreUserAPIRepository,
		})

		res, err := s.SignUp(ctx, req, h)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("error - when core user CallSignUp client error and should call core auth CallDeleteUser", func(t *testing.T) {
		authSignUpRes := &httpclient.Response[dto.CoreAuthSignUpRes]{
			HTTPStatusCode: http.StatusOK,
			Response: dto.CoreAuthSignUpRes{
				Username: mockUsername,
				UID:      mockUID,
				Token:    mockToken,
			},
		}

		authDeleteUserRes := &httpclient.Response[dto.CoreAuthDeleteUserRes]{
			HTTPStatusCode: http.StatusOK,
			Response: dto.CoreAuthDeleteUserRes{
				Success: true,
			},
		}

		coreAuthAPIRepository := &mockCoreAuthAPIRepository{
			signupResp:     authSignUpRes,
			deleteUserResp: authDeleteUserRes,
			err:            nil,
		}

		coreUserAPIRepository := &mockCoreUserAPIRepository{
			signupResp: nil,
			err:        errors.New("error"),
		}

		s := New(Dependencies{
			CoreAuthAPIRepository: coreAuthAPIRepository,
			CoreUserAPIRepository: coreUserAPIRepository,
		})

		res, err := s.SignUp(ctx, req, h)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("error - when core user CallSignUp http status code not 200 and should call core auth CallDeleteUser", func(t *testing.T) {
		authSignUpRes := &httpclient.Response[dto.CoreAuthSignUpRes]{
			HTTPStatusCode: http.StatusOK,
			Response: dto.CoreAuthSignUpRes{
				Username: mockUsername,
				UID:      mockUID,
				Token:    mockToken,
			},
		}

		authDeleteUserRes := &httpclient.Response[dto.CoreAuthDeleteUserRes]{
			HTTPStatusCode: http.StatusOK,
			Response: dto.CoreAuthDeleteUserRes{
				Success: true,
			},
		}

		userSignUpRes := &httpclient.Response[dto.CoreUserSignUpRes]{
			HTTPStatusCode: http.StatusInternalServerError,
			Response:       dto.CoreUserSignUpRes{},
		}

		coreAuthAPIRepository := &mockCoreAuthAPIRepository{
			signupResp:     authSignUpRes,
			deleteUserResp: authDeleteUserRes,
			err:            nil,
		}

		coreUserAPIRepository := &mockCoreUserAPIRepository{
			signupResp: userSignUpRes,
			err:        nil,
		}

		s := New(Dependencies{
			CoreAuthAPIRepository: coreAuthAPIRepository,
			CoreUserAPIRepository: coreUserAPIRepository,
		})

		res, err := s.SignUp(ctx, req, h)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}
