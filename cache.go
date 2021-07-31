package mp3lify

import (
	"encoding/base64"
	"errors"
	"io"
	"os"
	"path"
)

const cacheDir = "/tmp/audio"

func GetCache(src string) io.ReadCloser {
	key := base64.StdEncoding.EncodeToString([]byte(src))
	p := path.Join(cacheDir, key)
	f, err := os.Open(p)
	if err != nil {
		return nil
	}

	return f
}

func SetCache(src string) (io.WriteCloser, error) {
	key := base64.StdEncoding.EncodeToString([]byte(src))
	p := path.Join(cacheDir, key)

	if _, err := os.Stat(p); err == nil {
		return nil, errors.New("file exists")
	}

	f, err := os.Create(p)
	if err != nil {
		return nil, err
	}

	return f, nil
}
