package repository

import (
	"errors"
	"sync"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
)

type URLRepository struct {
	mu   sync.Mutex
	Data map[string]string
}

func NewMapRepository() URLRepository {
	data := make(map[string]string, 0)
	return URLRepository{
		Data: data,
	}
}

func (r *URLRepository) Get(path string) (*entity.URL, error) {
	// &entity.URL{} , nil // TODO: replace this
	r.mu.Lock()
	defer r.mu.Unlock()
	if url, ok := r.Data[path]; ok {
		return &entity.URL{
			ShortURL: path,
			LongURL:  url,
		}, nil
	}
	return nil, errors.New("url not found")

	/*
		if _, isExist := r.Data[path]; isExist {
			return nil, entity.ErrorNotFound
		}
		url := &entity.URL{
			LongURL:  r.Data[path],
			ShortURL: path,
		}
		return url, nil
	*/

}

func (r *URLRepository) Create(longURL string) (*entity.URL, error) {
	// &entity.URL{} , nil // TODO: replace this
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Data[longURL] = entity.GetRandomShortURL(longURL)
	return &entity.URL{
		LongURL:  longURL,
		ShortURL: r.Data[longURL],
	}, nil

	/*
		rndm := entity.GetRandomShortURL(longURL)
		r.Data[rndm] = longURL
		url := &entity.URL{
			LongURL: longURL,
			ShortURL: rndm,
		}
	*/
}

func (r *URLRepository) CreateCustom(longURL, customPath string) (*entity.URL, error) {
	// &entity.URL{} , nil // TODO: replace this
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Data[customPath] = longURL
	return &entity.URL{
		LongURL:  longURL,
		ShortURL: customPath,
	}, nil
}
