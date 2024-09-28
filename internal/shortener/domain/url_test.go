package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUrl(t *testing.T) {
	url := "http://example.com"
	u, err := NewUrl(url)
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, url, u.Url)
	assert.NotEqual(t, "", u.Id)
	assert.Equal(t, int64(0), u.AccessCount)
}

func TestUpdateUrl(t *testing.T) {
	url := "http://example.com"
	u, _ := NewUrl(url)
	newUrl := "http://example.org"
	u.UpdateUrl(newUrl)
	assert.Equal(t, newUrl, u.Url)
	assert.NotEqual(t, u.CreatedAt, u.UpdatedAt)
}

func TestIncrementAccessCount(t *testing.T) {
	url := "http://example.com"
	u, _ := NewUrl(url)
	u.IncrementAccessCount()
	assert.Equal(t, int64(1), u.AccessCount)
	u.IncrementAccessCount()
	assert.Equal(t, int64(2), u.AccessCount)
}
