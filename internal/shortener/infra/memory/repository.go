package infra_memory

import "github.com/rattapon001/slimlink/internal/shortener/domain"

type (
	UrlMemoryRepository struct {
		urls []*domain.Url
	}
)

func NewUrlMemoryRepository() *UrlMemoryRepository {
	return &UrlMemoryRepository{
		urls: []*domain.Url{},
	}
}

func (r *UrlMemoryRepository) Save(urls ...*domain.Url) error {
	r.urls = append(r.urls, urls...)
	return nil
}

func (r *UrlMemoryRepository) Find(id domain.UrlId) (*domain.Url, error) {
	for _, u := range r.urls {
		if u.Id == id {
			return u, nil
		}
	}
	return nil, domain.ErrUrlNotFound
}

func (r *UrlMemoryRepository) FindAll() ([]*domain.Url, error) {
	return r.urls, nil
}

func (r *UrlMemoryRepository) Delete(id domain.UrlId) error {
	for i, u := range r.urls {
		if u.Id == id {
			r.urls = append(r.urls[:i], r.urls[i+1:]...)
			return nil
		}
	}
	return domain.ErrUrlNotFound
}
