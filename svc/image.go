package svc

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"

	"github.com/bolg-developers/MikanMusic-API/infra"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/image/draw"
)

func CreateImage(r io.Reader, w, h int) (string, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return "", errors.WithStack(err)
	}

	resizedImg, err := resize(img, w, h)
	if err != nil {
		return "", errors.WithStack(err)
	}

	buf := bytes.NewBuffer(make([]byte, 0))
	if err := jpeg.Encode(buf, resizedImg, &jpeg.Options{Quality: 100}); err != nil {
		return "", errors.WithStack(err)
	}

	key := uuid.New().String() + ".jpg"
	reader := bytes.NewReader(buf.Bytes())
	if err = infra.Bucket().Create(reader, key); err != nil {
		return "", err
	}

	return infra.Bucket().URL(key), nil
}

func resize(img image.Image, w, h int) (image.Image, error) {
	ret := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.CatmullRom.Scale(ret, ret.Bounds(), img, img.Bounds(), draw.Over, nil)
	return ret, nil
}
