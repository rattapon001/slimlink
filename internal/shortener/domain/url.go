package domain

import (
	"math/rand"
	"time"
)

type (
	UrlId string
	Url   struct {
		Id          UrlId     `json:"id"`
		Url         string    `json:"url"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		AccessCount int64     `json:"access_count"`
	}
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// randomString generates a random string of length n
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

// NewUrl creates a new Url object
func NewUrl(url string) (*Url, error) {

	return &Url{
		Id:        UrlId(RandomString(8)),
		Url:       url,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil

}

// UpdateUrl updates the Url object
func (u *Url) UpdateUrl(url string) {
	u.Url = url
	u.UpdatedAt = time.Now()
}

// IncrementAccessCount increments the access count of the Url object
func (u *Url) IncrementAccessCount() {
	u.AccessCount++
}
