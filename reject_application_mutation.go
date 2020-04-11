package community_api_go

import (
	"context"
	"errors"
	"github.com/214alphadev/community-api-go/scalars"
	am "github.com/214alphadev/community-authentication-middleware"
)

type rejectApplicationInput struct {
	ApplicationID   scalars.UUIDV4
	RejectionReason string
}

type rejectApplicationResponse struct {
	error       error
	application *application
}

func (r rejectApplicationResponse) Error() *string {
	switch r.error {
	case nil:
		return nil
	default:
		e := r.error.Error()
		return &e
	}
}

func (r rejectApplicationResponse) Application() *application {
	return r.application
}

func (r *rootResolver) RejectApplication(ctx context.Context, args struct{ Input rejectApplicationInput }) (rejectApplicationResponse, error) {

	input := args.Input

	reviewer := am.GetAuthenticateMember(ctx)

	switch reviewer {
	case nil:
		return rejectApplicationResponse{error: errors.New("Unauthenticated")}, nil
	default:

		err := r.community.RejectApplication(input.ApplicationID.UUID, input.RejectionReason, *reviewer)

		switch err {
		case nil:

			application, err := r.community.GetApplication(args.Input.ApplicationID.UUID)
			if err != nil {
				return rejectApplicationResponse{}, err
			}

			gqlApplication, err := newApplicationTypeFromEntity(r.community, application)
			if err != nil {
				return rejectApplicationResponse{}, err
			}

			return rejectApplicationResponse{application: &gqlApplication}, nil

		default:

			switch err.Error() {
			case "InsufficientPermissions":
				return rejectApplicationResponse{error: err}, nil
			case "ApplicationDoesNotExist":
				return rejectApplicationResponse{error: err}, nil
			case "ApplicationReviewed":
				return rejectApplicationResponse{error: err}, nil
			default:
				return rejectApplicationResponse{}, err
			}

		}

	}

}
