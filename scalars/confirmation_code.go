package scalars

import (
	"encoding/json"
	"errors"
	vo "github.com/214alphadev/community-bl/value_objects"
)

type ConfirmationCode struct {
	Code vo.ConfirmationCode
}

func (ConfirmationCode) ImplementsGraphQLType(name string) bool {
	return name == "ConfirmationCode"
}

func (c *ConfirmationCode) UnmarshalGraphQL(input interface{}) error {

	switch v := input.(type) {
	case *string:
		cc, err := vo.NewConfirmationCode(*v)
		if err != nil {
			return err
		}
		c.Code = cc
		return nil
	case string:
		cc, err := vo.NewConfirmationCode(v)
		if err != nil {
			return err
		}
		c.Code = cc
		return nil
	default:
		return errors.New("failed to unmarshal confirmation code")
	}

}

func (c ConfirmationCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Code.String())
}
