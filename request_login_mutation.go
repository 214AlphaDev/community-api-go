package community_api_go

import (
	"context"
	"github.com/214alphadev/community-api-go/scalars"
	cd "github.com/214alphadev/community-bl"
)

func (r *rootResolver) RequestLogin(ctx context.Context, args struct{ EmailAddress scalars.EmailAddress }) (int32, error) {

	err := r.community.RequestLogin(args.EmailAddress.EmailAddress)

	switch v := err.(type) {
	case cd.RequestLoginCoolDownError:
		return int32(v.TryAgainAt), nil
	case nil:
		return 0, nil
	default:
		return 0, err
	}

}
