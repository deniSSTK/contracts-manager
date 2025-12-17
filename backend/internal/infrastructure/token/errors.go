package token

import "errors"

var ErrFailedToParseToken = errors.New("failed to parse token")

var errFailedToSignToken = errors.New("failed to sign token")

func ErrFailedToSignToken(tokenType string) error {
	return errors.New(errFailedToSignToken.Error() + ": " + tokenType)
}
