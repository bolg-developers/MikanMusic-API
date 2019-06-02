package svc

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/bolg-developers/MikanMusic-API/infra"
	"github.com/google/uuid"
)

func CreateAudio(r io.Reader) (string, error) {
	buffer, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	key := uuid.New().String() + ".mp3"
	if err = infra.Bucket().Create(bytes.NewReader(buffer), key); err != nil {
		return "", err
	}

	return infra.Bucket().URL(key), nil
}
