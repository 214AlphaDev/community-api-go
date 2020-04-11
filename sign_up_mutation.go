package community_api_go

import (
	"context"
	"github.com/214alphadev/community-api-go/scalars"
	cd "github.com/214alphadev/community-bl"
	vo "github.com/214alphadev/community-bl/value_objects"
)

type signUpResponse struct {
	error  error
	member *Member
}

func (r signUpResponse) Error() *string {
	switch r.error {
	case nil:
		return nil
	default:
		e := r.error.Error()
		return &e
	}
}

func (r signUpResponse) Member() *Member {
	return r.member
}

type ProperNameInput struct {
	FirstName string
	LastName  string
}

type signUpInput struct {
	Username       scalars.Username
	EmailAddress   scalars.EmailAddress
	ProperName     ProperNameInput
	ProfilePicture *scalars.ProfilePicture
}

func (r *rootResolver) SignUp(ctx context.Context, args struct{ Input signUpInput }) (signUpResponse, error) {

	input := args.Input

	properName, err := vo.NewProperName(input.ProperName.FirstName, input.ProperName.LastName)
	if err != nil {
		return signUpResponse{}, err
	}

	metadata := cd.MetadataEntity{
		ProperName: properName,
	}

	if args.Input.ProfilePicture != nil {
		metadata.ProfileImage = &args.Input.ProfilePicture.Base64String
	}

	member, err := r.community.SignUp(args.Input.Username.Username, args.Input.EmailAddress.EmailAddress, metadata)

	switch err {
	case nil:
		m := newMemberTypeFromEntity(member, member)
		return signUpResponse{member: &m}, nil
	default:

		switch err.Error() {
		case "UsernameTaken":
			return signUpResponse{error: err}, nil
		case "EmailAddressTaken":
			return signUpResponse{error: err}, nil
		default:
			return signUpResponse{}, err
		}

	}

}
