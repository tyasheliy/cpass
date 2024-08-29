package passcl

import (
	"context"
)

type Client interface {
	Init(ctx context.Context, subFolder *string, key string) error
	Insert(ctx context.Context, passName string, data []string, options InsertOptions) error
	InsertOtp(ctx context.Context, passName string, uri string, options InsertOtpOptions) error
	Show(ctx context.Context, passName string) (string, error)
	ShowOtp(ctx context.Context, passName string) (string, error)
	Generate(ctx context.Context, passName string, options GenerateOptions) error
}

type GenerateOptions struct {
	Force     bool
	NoSymbols bool
	Length    int
}

type InsertOptions struct {
	Force     bool
	MultiLine bool
}

type InsertOtpOptions struct {
	Force bool
}
