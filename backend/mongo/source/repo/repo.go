package repo

import "context"

type Repo interface {
	Find(ctx *context.Context, key string, value interface{}) interface{}
	FindOne(ctx *context.Context, key string, value interface{}) interface{}
}

func Find(repoImpl Repo, ctx *context.Context, key string, value interface{}) interface{} {
	return repoImpl.Find(ctx, key, value)
}

func FindOne(repoImpl Repo, ctx *context.Context, key string, value interface{}) interface{} {
	return repoImpl.FindOne(ctx, key, value)
}
