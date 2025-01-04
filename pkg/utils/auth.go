package utils

import (
	"context"
	"errors"
)

func ValidateIsOwner(ctx context.Context, ownerId string) error {
	user, err := GetUserFromCtx(ctx)
	if err != nil {
		return err
	}
	if user.UserID.String() != ownerId {
		return errors.New("unauthorized")
	}
	return nil
}
