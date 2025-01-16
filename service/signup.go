package service

import (
	"context"
	"net/http"

	"github.com/project-inari/orch-auth-external/dto"
)

func (s *service) SignUp(ctx context.Context, req dto.SignUpReq, h dto.SignUpReqHeader) (*dto.SignUpRes, error) {
	authSignUpReq := constructCoreAuthSignUpReq(req)
	authSignUpReqHeader := constructCoreAuthSignUpReqHeader(h)

	authSignUpRes, err := s.coreAuthAPIRepository.SignUp(ctx, authSignUpReq, authSignUpReqHeader)
	if err != nil || authSignUpRes.HTTPStatusCode != http.StatusOK {
		return nil, err
	}

	userSignUpReq := constructCoreUserSignUpReq(req, authSignUpRes.Response.UID, h.AcceptLocale)

	userSignUpRes, err := s.coreUserAPIRepository.SignUp(ctx, userSignUpReq)
	if err != nil || userSignUpRes.HTTPStatusCode != http.StatusOK {
		deleteUserReq := constructCoreAuthDeleteUserReq(authSignUpRes.Response.UID)
		_, _ = s.coreAuthAPIRepository.DeleteUser(ctx, deleteUserReq)
		return nil, err
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
