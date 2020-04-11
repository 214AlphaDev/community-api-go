package community_api_go

import (
	"context"
	gql "github.com/graph-gophers/graphql-go"
	"github.com/214alphadev/community-api-go/scalars"
	cd "github.com/214alphadev/community-bl"
	"strings"
)

type application struct {
	cd.ApplicationEntity
	community cd.CommunityInterface
}

func (a application) Member(ctx context.Context) (Member, error) {
	return newMemberType(a.community, a.ApplicationEntity.MemberID, a.ApplicationEntity.MemberID)
}

func (a application) ApplicationText() string {
	return a.ApplicationEntity.ApplicationText
}

func (a application) State() string {
	return strings.Title(strings.ToLower(string(a.ApplicationEntity.State)))
}

func (a application) CreatedAt() gql.Time {
	return gql.Time{
		Time: a.ApplicationEntity.CreatedAt,
	}
}

func (a application) RejectionReason() *string {
	switch a.ApplicationEntity.RejectionReason {
	case "":
		return nil
	default:
		r := a.ApplicationEntity.RejectionReason
		return &r
	}
}

func (a application) RejectedAt() *gql.Time {
	switch a.ApplicationEntity.RejectedAt {
	case nil:
		return nil
	default:
		return &gql.Time{
			Time: *a.ApplicationEntity.RejectedAt,
		}
	}
}

func (a application) ApprovedAt() *gql.Time {
	switch a.ApplicationEntity.ApprovedAt {
	case nil:
		return nil
	default:
		return &gql.Time{
			Time: *a.ApplicationEntity.ApprovedAt,
		}
	}
}

func (a application) ID() scalars.UUIDV4 {
	return scalars.UUIDV4{
		UUID: a.ApplicationEntity.ID,
	}
}

func newApplicationTypeFromEntity(community cd.CommunityInterface, a cd.ApplicationEntity) (application, error) {
	return application{
		ApplicationEntity: a,
		community:         community,
	}, nil
}
