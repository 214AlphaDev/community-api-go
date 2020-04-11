package community_api_go

import (
	"context"
	am "github.com/214alphadev/community-authentication-middleware"
)

func (r *rootResolver) MyCurrentApplication(ctx context.Context) (*application, error) {

	memberId := am.GetAuthenticateMember(ctx)

	switch memberId {
	case nil:
		return nil, nil
	default:
		a, err := r.community.GetLastApplication(*memberId, *memberId)
		switch err {
		case nil:
			a, err := newApplicationTypeFromEntity(r.community, a)
			return &a, err
		default:
			if isEqualError(err, "ApplicationNotFound") {
				return nil, nil
			}
			return nil, err
		}
	}

}
