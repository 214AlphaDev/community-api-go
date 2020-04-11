package scalars

import (
	"encoding/json"
	"errors"
	vo "github.com/214alphadev/community-bl/value_objects"
)

type Username struct {
	vo.Username
}

func (Username) ImplementsGraphQLType(name string) bool {
	return name == "Username"
}

func (u *Username) UnmarshalGraphQL(input interface{}) error {

	switch v := input.(type) {
	case *string:
		username, err := vo.NewUsername(*v)
		if err != nil {
			return err
		}
		u.Username = username
		return nil
	case string:
		username, err := vo.NewUsername(v)
		if err != nil {
			return err
		}
		u.Username = username
		return nil
	default:
		return errors.New("failed to unmarshal username")
	}

}

func (u Username) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Username.String())
}
