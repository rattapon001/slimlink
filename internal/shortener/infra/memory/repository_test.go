package infra_memory

import (
	"github.com/rattapon001/slimlink/internal/shortener/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlMemoryRepository(t *testing.T) {
	repo := NewUrlMemoryRepository()
	url, err := domain.NewUrl("https://example.com")
	assert.Nil(t, err)

	err = repo.Save(url)
	assert.Nil(t, err)

	u, err := repo.Find(url.Id)
	assert.Nil(t, err)
	assert.Equal(t, url, u)

	urls, err := repo.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(urls))

	err = repo.Delete(url.Id)
	assert.Nil(t, err)

	_, err = repo.Find(url.Id)
	assert.Equal(t, domain.ErrUrlNotFound, err)
}
