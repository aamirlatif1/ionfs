package store

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultRootFolderName = "ionfs-store"
)

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])
	blockSize := 5
	sliceLen := len(hashStr) / blockSize
	paths := make([]string, 0, blockSize)
	for i := 0; i < sliceLen; i++ {
		paths = append(paths, hashStr[i*blockSize:(i+1)*blockSize])
	}
	return PathKey{
		Pathname: strings.Join(paths, "/"),
		Filename: hashStr,
	}
}

type PathKey struct {
	Pathname string
	Filename string
}

func (p PathKey) FullPath() string {
	return fmt.Sprintf("%s/%s", p.Pathname, p.Filename)
}

type PathTransformFunc func(string) PathKey

type StoreOpts struct {
	// Root is the folder name of the root, containing all the folder/files of the system.
	Root              string
	PathTransformFunc PathTransformFunc
}

var DefaultPathTransformFunc = func(key string) PathKey {
	return PathKey{
		Pathname: key,
		Filename: key,
	}
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	if opts.PathTransformFunc == nil {
		opts.PathTransformFunc = DefaultPathTransformFunc
	}
	if len(opts.Root) == 0 {
		opts.Root = defaultRootFolderName
	}
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) Write(key string, r io.Reader) error {
	pathKey := s.PathTransformFunc(key)
	if err := os.MkdirAll(s.Root+"/"+pathKey.Pathname, os.ModePerm); err != nil {
		return err
	}
	fullpathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())

	f, err := os.Create(fullpathWithRoot)
	if err != nil {
		return err
	}
	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}
	log.Printf("written (%d) byte to disk %s", n, pathKey)
	return nil
}

func (s *Store) Read(key string) (io.Reader, error) {
	f, err := s.readStream(key)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)
	return buf, nil
}

func (s *Store) readStream(key string) (io.ReadCloser, error) {
	pathKey := s.PathTransformFunc(key)
	fullpathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())
	f, err := os.Open(fullpathWithRoot)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (s *Store) Delete(key string) error {
	pathKey := s.PathTransformFunc(key)
	fullpathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())
	defer func() {
		log.Printf("deleted [%s from disk]", pathKey.Filename)
	}()
	err := os.RemoveAll(fullpathWithRoot)
	if err != nil {
		return err
	}

	dir := filepath.Dir(fullpathWithRoot)
	for {
		parent := filepath.Dir(dir)
		if dir == parent {
			break
		}
		err := os.Remove(dir)
		if err != nil {
			break
		}
		dir = parent
	}
	return nil
}

func (s *Store) Has(key string) bool {
	pathKey := s.PathTransformFunc(key)
	fullpathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())
	_, err := os.Stat(fullpathWithRoot)
	if err != nil && err == fs.ErrNotExist {
		return false
	}
	return true
}
