package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockPass struct {
	Name string
}

type MockFail struct{}

var (
	pass = MockPass{Name: "v0.0.0"}
	fail = MockFail{}
)

func (pass *MockPass) GetVersion(getURL string) error {
	return nil
}

func (fail *MockFail) GetVersion(getURL string) error {
	return errors.New("error test")
}

func Test_GetLatestRelease(t *testing.T) {
	err := GetLatestRelease("url", &pass)
	assert.Nil(t, err)
	assert.Equal(t, pass.Name, "v0.0.0")

	err = GetLatestRelease("url", &fail)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "error test")
}
