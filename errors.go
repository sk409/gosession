package gosession

import (
	"errors"
	"fmt"
)

func errorSpecifiedKeyDoesNotExsistInThisSession(key string) error {
	errorMessage := fmt.Sprintf("%s does not exists in this session", key)
	return errors.New(errorMessage)
}

func errorValueCorrespondingToKeyIsNotSpecifiedType(key string, specifiedType string) error {
	errorMessage := fmt.Sprintf("The value corresponding to the %s is not %s", key, specifiedType)
	return errors.New(errorMessage)
}
