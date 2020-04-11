package community_api_go

import (
	cd "github.com/214alphadev/community-bl"
)

type rootResolver struct {
	community cd.CommunityInterface
}

func newResolver(community cd.CommunityInterface) *rootResolver {
	return &rootResolver{
		community: community,
	}
}
