package community_api_go

import (
	"context"
	"errors"
	"github.com/214alphadev/community-api-go/scalars"
	cd "github.com/214alphadev/community-bl"
	vo "github.com/214alphadev/community-bl/value_objects"
)

type loginInput struct {
	EmailAddress          scalars.EmailAddress
	MemberAccessPublicKey scalars.Ed25519PublicKey
	ConfirmationCode      scalars.ConfirmationCode
}

type loginResponse struct {
	error       error
	accessToken string
	member      *Member
}

func (r loginResponse) Error() *string {
	switch r.error {
	case nil:
		return nil
	default:
		e := r.error.Error()
		return &e
	}
}

func (r loginResponse) AccessToken() *string {
	switch r.accessToken {
	case "":
		return nil
	default:
		a := r.accessToken
		return &a
	}
}

func (r loginResponse) Member() *Member {
	return r.member
}

func (r *rootResolver) Login(ctx context.Context, args struct{ Input loginInput }) (loginResponse, error) {

	input := args.Input

	memberAccessPublicKey, err := vo.NewMemberAccessPublicKey(input.MemberAccessPublicKey.PublicKey)
	if err != nil {
		return loginResponse{}, err
	}

	accessToken, err := r.community.Login(input.EmailAddress.EmailAddress, memberAccessPublicKey, input.ConfirmationCode.Code)
	switch err {
	case nil:

		member, err := newMemberType(r.community, accessToken.Subject, accessToken.Subject)
		if err != nil {
			return loginResponse{}, err
		}

		return loginResponse{
			accessToken: accessToken.SignedAccessToken(),
			member:      &member,
		}, nil

	case cd.LoginErrorMemberAccessKeyHasAlreadyBeenUsed:
		return loginResponse{error: errors.New("MemberAccessKeyHasAlreadyBeenUsed")}, nil
	case cd.LoginErrorConfirmationCodeNotFound:
		return loginResponse{error: errors.New("ConfirmationCodeNotFound")}, nil
	case cd.LoginErrorConfirmationCodeExpired:
		return loginResponse{error: errors.New("ConfirmationCodeExpired")}, nil
	case cd.LoginErrorConfirmationCodeMemberMismatch:
		return loginResponse{error: errors.New("ConfirmationCodeMemberMismatch")}, nil
	case cd.LoginErrorConfirmationCodeAlreadyUsed:
		return loginResponse{error: errors.New("ConfirmationCodeAlreadyUsed")}, nil
	default:
		return loginResponse{}, err
	}

}
