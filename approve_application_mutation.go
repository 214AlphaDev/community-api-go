package community_api_go

import (
	"context"
	"errors"
	"github.com/214alphadev/community-api-go/scalars"
	am "github.com/214alphadev/community-authentication-middleware"
)

type approveApplicationResponse struct {
	error       error
	application *application
}

func (r approveApplicationResponse) Error() *string {
	switch r.error {
	case nil:
		return nil
	default:
		e := r.error.Error()
		return &e
	}
}

func (r approveApplicationResponse) Application() *application {
	return r.application
}

func (r *rootResolver) ApproveApplication(ctx context.Context, args struct{ ApplicationID scalars.UUIDV4 }) (applyForVerificationResponse, error) {

	reviewer := am.GetAuthenticateMember(ctx)

	switch reviewer {
	case nil:

		return applyForVerificationResponse{error: errors.New("Unauthenticated")}, nil

	default:

		err := r.community.ApproveApplication(args.ApplicationID.UUID, *reviewer)
		switch err {
		case errors.New("InsufficientPermissions"):
			return applyForVerificationResponse{error: err}, nil
		case errors.New("ApplicationDoesNotExist"):
			return applyForVerificationResponse{error: err}, nil
		case errors.New("AlreadyReviewed"):
			return applyForVerificationResponse{error: err}, nil
		default:

			application, err := r.community.GetApplication(args.ApplicationID.UUID)
			if err != nil {
				return applyForVerificationResponse{}, err
			}

			gqlApplication, err := newApplicationTypeFromEntity(r.community, application)
			if err != nil {
				return applyForVerificationResponse{}, err
			}

			return applyForVerificationResponse{application: &gqlApplication}, nil

		}

	}

}
