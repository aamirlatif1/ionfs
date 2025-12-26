package store_test

import (
	"bytes"
	"testing"

	"github.com/aamirlatif1/ionfs/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	expectedOriginalPath := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPath := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	actualPath := store.CASPathTransformFunc(key)
	assert.Equal(t, expectedOriginalPath, actualPath.Filename)
	assert.Equal(t, expectedPath, actualPath.Pathname)
}

func TestWrite(t *testing.T) {
	opts := store.StoreOpts{
		PathTransformFunc: store.CASPathTransformFunc,
	}
	key := "myspecialpicture"
	s := store.NewStore(opts)
	data := bytes.NewReader([]byte("hello world"))
	err := s.Write(key, data)
	assert.NoError(t, err)

	err = s.Delete(key)
	assert.NoError(t, err)

}

func TestRead(t *testing.T) {
	opts := store.StoreOpts{
		PathTransformFunc: store.CASPathTransformFunc,
	}
	key := "myspecialpicture"
	s := store.NewStore(opts)
	data := bytes.NewReader([]byte("hello world"))
	err := s.Write(key, data)
	assert.NoError(t, err)

	r, err := s.Read(key)
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(r)
	assert.NoError(t, err)
	assert.Equal(t, "hello world", buf.String())

	err = s.Delete(key)
	assert.NoError(t, err)
}
