package encrypt

import "context"

//go:generate mockgen -source encryptor.go -package encryptormock -destination mock/mock.go

type Encryptor interface {
	Encrypt(ctx context.Context, raw string) (string, error)
	Check(ctx context.Context, realString, encryptedString string) (bool, error)
}
