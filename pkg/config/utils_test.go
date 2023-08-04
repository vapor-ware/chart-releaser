package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigPath(t *testing.T) {
	path, err := GetConfigPath("")
	assert.NoError(t, err)

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	expected := fmt.Sprintf("%s/%s", cwd, DefaultFile)
	assert.Equal(t, expected, path)
}

func TestGetConfigPath2(t *testing.T) {
	path, err := GetConfigPath("foobar")
	assert.Error(t, err)
	assert.Equal(t, "", path)
}

func TestGetConfigPath3(t *testing.T) {
	path, err := GetConfigPath(".")
	assert.NoError(t, err)
	assert.Equal(t, DefaultFile, path)
}

func TestGetConfigPath4(t *testing.T) {
	dir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	path, err := GetConfigPath(dir)
	assert.NoError(t, err)
	assert.Equal(t, filepath.Join(dir, DefaultFile), path)
}
