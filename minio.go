package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Client is an alias for the MinIO client type.
type Client = *minio.Client

// Config holds the configuration parameters required to connect to MinIO.
type Config struct {
	Address  string `confy:"address" yaml:"address" json:"address" toml:"address" env:"MINIO_ADDRESS"`
	User     string `confy:"user"    yaml:"user"    json:"user"    toml:"user"    env:"MINIO_USER"`
	Password string `confy:"password" env:"MINIO_PASSWORD"`
	Bucket   string `confy:"bucket"  yaml:"bucket"  json:"bucket"  toml:"bucket"  env:"MINIO_BUCKET"`
	UseSSL   bool   `confy:"use_ssl" yaml:"use_ssl" json:"use_ssl" toml:"use_ssl" env:"MINIO_USE_SSL"`
}

// New creates a new MinIO client and ensures that the specified bucket exists.
func New(ctx context.Context, cfg *Config) (Client, error) {
	client, err := minio.New(
		cfg.Address,
		&minio.Options{
			Creds: credentials.NewStaticV4(
				cfg.User, cfg.Password, "",
			),
			Secure: cfg.UseSSL,
		},
	)

	if err != nil {
		return nil, err
	}

	exists, err := client.BucketExists(ctx, cfg.Bucket)
	if err != nil {
		return nil, err
	}

	if !exists {
		err := client.MakeBucket(ctx, cfg.Bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}
