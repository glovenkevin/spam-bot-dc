package repo

import "context"

type AmythestRepoInterface interface {
	GenerateWanted(ctx context.Context, imageLink string) ([]byte, error)
	GenerateCircle(ctx context.Context, imageLink string) ([]byte, error)
}
