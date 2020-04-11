package scalars

import (
	"encoding/json"
	"errors"
	vo "github.com/214alphadev/community-bl/value_objects"
)

type ProfilePicture struct {
	vo.Base64String
}

func (ProfilePicture) ImplementsGraphQLType(name string) bool {
	return name == "ProfilePicture"
}

func (p *ProfilePicture) UnmarshalGraphQL(input interface{}) error {

	switch v := input.(type) {
	case *string:
		s, err := vo.NewBase64String(*v)
		switch err {
		case nil:
			p.Base64String = s
			return nil
		default:
			return err
		}
	case string:
		s, err := vo.NewBase64String(v)
		switch err {
		case nil:
			p.Base64String = s
			return nil
		default:
			return err
		}
	default:
		return errors.New("failed to unmarshal profile picture")
	}

}

func (p ProfilePicture) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Base64String.String())
}
