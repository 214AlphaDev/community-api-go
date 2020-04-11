package community_api_go

import (
	gql "github.com/graph-gophers/graphql-go"
	cd "github.com/214alphadev/community-bl"
	gqlh "github.com/214alphadev/graphql-handler"
)

func NewCommunityApi(community cd.CommunityInterface, logger gqlh.Logger) (*gqlh.Handler, error) {

	resolver := newResolver(community)

	schema, err := gql.ParseSchema(schema, resolver)
	if err != nil {
		return nil, err
	}

	return gqlh.NewHandler(schema, logger)

}
