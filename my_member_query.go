package community_api_go

import (
	"context"
	am "github.com/214alphadev/community-authentication-middleware"
)

func (r *rootResolver) MyMember(ctx context.Context) (*Member, error) {

	memberID := am.GetAuthenticateMember(ctx)

	switch memberID {
	case nil:
		return nil, nil
	default:
		m, err := newMemberType(r.community, *memberID, *memberID)
		if err != nil {
			return nil, err
		}
		return &m, nil
	}

}
