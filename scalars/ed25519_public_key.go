package scalars

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/ed25519"
	"reflect"
)

type Ed25519PublicKey struct {
	ed25519.PublicKey
}

func (Ed25519PublicKey) ImplementsGraphQLType(name string) bool {
	return name == "Ed25519PublicKey"
}

func (k *Ed25519PublicKey) UnmarshalGraphQL(input interface{}) error {

	var hexToEdPubKey = func(hexKey string) (ed25519.PublicKey, error) {
		pk, err := hex.DecodeString(hexKey)
		if err != nil {
			return nil, err
		}

		if len(pk) != 32 {
			return nil, errors.New("ed25519 public key is too short")
		}

		nilPubKey := make([]byte, 32)

		if reflect.DeepEqual(pk, nilPubKey) {
			return nil, errors.New("ed25519 public key is a slice of empty bytes")
		}

		return ed25519.PublicKey(pk), nil

	}

	switch v := input.(type) {
	case *string:
		pk, err := hexToEdPubKey(*v)
		if err != nil {
			return err
		}
		k.PublicKey = pk
		return nil
	case string:
		pk, err := hexToEdPubKey(v)
		if err != nil {
			return err
		}
		k.PublicKey = pk
		return nil
	default:
		return errors.New("failed to unmarshal ed25519 public key")
	}

}

func (k Ed25519PublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(k.PublicKey))
}
