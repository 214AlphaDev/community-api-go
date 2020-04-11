package scalars

import (
	"encoding/json"
	"errors"
	vo "github.com/214alphadev/community-bl/value_objects"
)

type EmailAddress struct {
	vo.EmailAddress
}

func (EmailAddress) ImplementsGraphQLType(name string) bool {
	return name == "EmailAddress"
}

func (e *EmailAddress) UnmarshalGraphQL(input interface{}) error {

	switch v := input.(type) {
	case *string:
		emailAddress, err := vo.NewEmailAddress(*v)
		if err != nil {
			return err
		}
		e.EmailAddress = emailAddress
		return nil
	case string:
		emailAddress, err := vo.NewEmailAddress(v)
		if err != nil {
			return err
		}
		e.EmailAddress = emailAddress
		return nil
	default:
		return errors.New("failed to unmarshal email address")
	}

}

func (e EmailAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.EmailAddress.String())
}
