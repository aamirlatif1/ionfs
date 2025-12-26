package store_test

import (
	"bytes"
	"testing"

	"github.com/aamirlatif1/ionfs/internal/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StoreTestSuite struct {
	suite.Suite
	key   string
	store *store.Store
}

func (suite *StoreTestSuite) SetupTest() {
	suite.key = "momsbestpicture"
	opts := store.StoreOpts{
		PathTransformFunc: store.CASPathTransformFunc,
	}
	suite.store = store.NewStore(opts)
}

func (suite *StoreTestSuite) TearDownTest() {
	err := suite.store.Delete(suite.key)
	assert.NoError(suite.T(), err)
}

func (suite *StoreTestSuite) TestPathTransformFunc() {
	expectedOriginalPath := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPath := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	actualPath := store.CASPathTransformFunc(suite.key)
	assert.Equal(suite.T(), expectedOriginalPath, actualPath.Filename)
	assert.Equal(suite.T(), expectedPath, actualPath.Pathname)
}

func (suite *StoreTestSuite) TestWrite() {

	data := bytes.NewReader([]byte("hello world"))

	err := suite.store.Write(suite.key, data)

	assert.NoError(suite.T(), err)
	assert.True(suite.T(), suite.store.Has(suite.key))
}

func (suite *StoreTestSuite) TestRead() {

	data := bytes.NewReader([]byte("hello world"))
	err := suite.store.Write(suite.key, data)
	assert.NoError(suite.T(), err)

	r, err := suite.store.Read(suite.key)
	assert.NoError(suite.T(), err)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(r)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "hello world", buf.String())
}

func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}
