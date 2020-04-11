package community_api_go

import (
	"context"
	"errors"
	am "github.com/214alphadev/community-authentication-middleware"
)

type applyForVerificationResponse struct {
	error       error
	application *application
}

func (r applyForVerificationResponse) Error() *string {
	switch r.error {
	case nil:
		return nil
	default:
		e := r.error.Error()
		return &e
	}
}

func (r applyForVerificationResponse) Application() *application {
	return r.application
}

func (r *rootResolver) ApplyForVerification(ctx context.Context, args struct{ ApplicationText string }) (applyForVerificationResponse, error) {

	memberID := am.GetAuthenticateMember(ctx)

	switch memberID {
	case nil:
		return applyForVerificationResponse{error: errors.New("Unauthenticated")}, nil
	default:

		applicationEntity, err := r.community.ApplyForVerification(args.ApplicationText, *memberID)

		switch err {
		case nil:

			application, err := newApplicationTypeFromEntity(r.community, applicationEntity)
			if err != nil {
				return applyForVerificationResponse{}, err
			}

			return applyForVerificationResponse{application: &application}, nil

		default:

			switch err.Error() {
			case "PendingApplication":
				return applyForVerificationResponse{error: err}, nil
			case "AlreadyVerified":
				return applyForVerificationResponse{error: err}, nil
			default:
				return applyForVerificationResponse{}, err
			}

		}

	}

}
