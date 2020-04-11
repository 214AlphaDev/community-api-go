package community_api_go

import (
	"github.com/214alphadev/community-api-go/scalars"
	cd "github.com/214alphadev/community-bl"
	vo "github.com/214alphadev/community-bl/value_objects"
)

type ProperName struct {
	vo.ProperName
}

func (pn ProperName) FirstName() string {
	return pn.ProperName.FirstName()
}

func (pn ProperName) LastName() string {
	return pn.ProperName.LastName()
}

type Member struct {
	cd.MemberEntity
	requester cd.MemberEntity
}

func newMemberType(community cd.CommunityInterface, memberID cd.MemberIdentifier, requesterID cd.MemberIdentifier) (Member, error) {

	member, err := community.GetMember(memberID)
	if err != nil {
		return Member{}, err
	}

	requester, err := community.GetMember(requesterID)
	if err != nil {
		return Member{}, err
	}

	return Member{
		MemberEntity: member,
		requester:    requester,
	}, nil

}

func newMemberTypeFromEntity(member cd.MemberEntity, requester cd.MemberEntity) Member {
	return Member{
		MemberEntity: member,
		requester:    requester,
	}
}

func (m Member) ID() scalars.UUIDV4 {
	return scalars.UUIDV4{
		UUID: m.MemberEntity.ID,
	}
}

func (m Member) Name() ProperName {
	return ProperName{
		ProperName: m.MemberEntity.Metadata.ProperName,
	}
}

func (m Member) Username() scalars.Username {
	return scalars.Username{
		Username: m.MemberEntity.Username,
	}
}

func (m Member) ProfilePicture() *scalars.ProfilePicture {
	if m.Metadata.ProfileImage == nil {
		return nil
	}
	return &scalars.ProfilePicture{
		Base64String: *m.MemberEntity.Metadata.ProfileImage,
	}
}

func (m Member) Admin() bool {
	return m.MemberEntity.Admin
}

func (m Member) EmailAddress() *scalars.EmailAddress {

	if m.MemberEntity.ID == m.requester.ID {
		return &scalars.EmailAddress{
			EmailAddress: m.MemberEntity.EmailAddress,
		}
	}

	if m.requester.Admin {
		return &scalars.EmailAddress{
			EmailAddress: m.MemberEntity.EmailAddress,
		}
	}

	return nil

}
