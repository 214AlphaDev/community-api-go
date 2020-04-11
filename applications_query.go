package community_api_go

import (
	"context"
	"errors"
	"github.com/214alphadev/community-api-go/scalars"
	am "github.com/214alphadev/community-authentication-middleware"
	cd "github.com/214alphadev/community-bl"
)

type ApplicationsInput struct {
	Position *scalars.UUIDV4
	Next     int32
	State    string
}

type applicationsResponse struct {
	error        error
	applications []application
}

func (ar applicationsResponse) Error() *string {
	switch ar.error {
	case nil:
		return nil
	default:
		e := ar.error.Error()
		return &e
	}
}

func (ar applicationsResponse) Applications() []application {
	return ar.applications
}

func (r *rootResolver) Applications(ctx context.Context, query ApplicationsInput) (applicationsResponse, error) {

	memberID := am.GetAuthenticateMember(ctx)

	switch memberID {
	case nil:
		return applicationsResponse{error: errors.New("Unauthenticated")}, nil
	default:

		q := cd.ApplicationsQuery{
			Next:  uint(query.Next),
			State: cd.ApplicationState(query.State),
		}

		if query.Position != nil {
			q.Position = &query.Position.UUID
		}

		applications, err := r.community.Applications(q, *memberID)

		switch err {
		case nil:

			gqlApplications := []application{}
			for _, a := range applications {
				gqla, err := newApplicationTypeFromEntity(r.community, a)
				if err != nil {
					return applicationsResponse{}, err
				}
				gqlApplications = append(gqlApplications, gqla)
			}
			return applicationsResponse{
				applications: gqlApplications,
			}, nil

		default:

			switch err.Error() {
			case "InsufficientPermissions":
				return applicationsResponse{error: err}, nil
			default:
				return applicationsResponse{}, err
			}

		}

	}

}
