package community_api_go

func isEqualError(error error, errors ...string) bool {

	if error == nil {
		return false
	}

	for _, err := range errors {
		if error.Error() == err {
			return true
		}
	}

	return false

}
