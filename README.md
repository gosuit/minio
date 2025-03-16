# Minio

This library is a simple wrapper around the <a href="https://github.com/minio/minio-go">MinIO Go</a> that facilitates the creation and configuration of a MinIO client.

## Installation

```zsh
go get github.com/gosuit/minio
```

## Features

- Easy configuration using environment variables or YAML files.
- Automatic bucket creation if the specified bucket does not exist.

## Usage

```golang
import (
    "context"
    "log"

    "github.com/gosuit/minio" 
)

func main() {
    ctx := context.Background()

    // Create a new configuration
    cfg := &minio.Config{
        Address:  "localhost:9000",
        User:     "your-access-key",
        Password: "your-secret-key",
        Bucket:   "your-bucket-name",
        UseSSL:   false,
    }

    // Create a new MinIO client
    client, err := minio.New(ctx, cfg)
    if err != nil {
        log.Fatalf("Failed to create MinIO client: %v", err)
    }

    log.Println("MinIO client created successfully!")
}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
