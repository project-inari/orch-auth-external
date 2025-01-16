package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/project-inari/orch-auth-external/dto"
	"github.com/project-inari/orch-auth-external/pkg/httpclient"
)

func (s *service) SignUp(ctx context.Context, req dto.SignUpReq, h dto.SignUpReqHeader) (*dto.SignUpRes, error) {
	authSignUpReq := constructCoreAuthSignUpReq(req)
	authSignUpReqHeader := constructCoreAuthSignUpReqHeader(h)

	authSignUpRes, err := s.coreAuthAPIRepository.CallSignUp(ctx, authSignUpReq, authSignUpReqHeader)
	if err != nil {
		return nil, err
	}
	if authSignUpRes.HTTPStatusCode != http.StatusOK {
		return nil, fmt.Errorf("error - [service.SignUp] - core-auth-server http status code returns %v", authSignUpRes.HTTPStatusCode)
	}

	userSignUpReq := constructCoreUserSignUpReq(req, authSignUpRes.Response.UID, h.AcceptLocale)

	userSignUpRes, err := s.coreUserAPIRepository.CallSignUp(ctx, userSignUpReq)
	if err != nil {
		s.handleCoreUserSignUpResponseError(ctx, authSignUpRes)
		return nil, err
	}
	if userSignUpRes.HTTPStatusCode != http.StatusOK {
		s.handleCoreUserSignUpResponseError(ctx, authSignUpRes)
		return nil, fmt.Errorf("error - [service.SignUp] - core-user-server http status code returns %v", authSignUpRes.HTTPStatusCode)
	}

	return constructSignUpResponse(userSignUpReq.Username, userSignUpReq.UID, authSignUpRes.Response.Token), nil
}

func constructCoreAuthSignUpReq(req dto.SignUpReq) dto.CoreAuthSignUpReq {
	return dto.CoreAuthSignUpReq{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		PhoneNo:  req.PhoneNo,
	}
}

func constructCoreAuthSignUpReqHeader(h dto.SignUpReqHeader) dto.CoreAuthSignUpReqHeader {
	return dto.CoreAuthSignUpReqHeader{
		AcceptLocale: h.AcceptLocale,
	}
}

func constructCoreUserSignUpReq(req dto.SignUpReq, uid string, selectedLocale string) dto.CoreUserSignUpReq {
	return dto.CoreUserSignUpReq{
		Username:       req.Username,
		UID:            uid,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		PhoneNo:        req.PhoneNo,
		Email:          req.Email,
		SelectedLocale: selectedLocale,
	}
}

func constructCoreAuthDeleteUserReq(uid string) dto.CoreAuthDeleteUserReq {
	return dto.CoreAuthDeleteUserReq{
		UID: uid,
	}
}

func constructSignUpResponse(username, uid, token string) *dto.SignUpRes {
	return &dto.SignUpRes{
		Username: username,
		UID:      uid,
		Token:    token,
	}
}

func (s *service) handleCoreUserSignUpResponseError(ctx context.Context, authSignUpRes *httpclient.Response[dto.CoreAuthSignUpRes]) {
	deleteUserReq := constructCoreAuthDeleteUserReq(authSignUpRes.Response.UID)
	_, _ = s.coreAuthAPIRepository.CallDeleteUser(ctx, deleteUserReq)
}
