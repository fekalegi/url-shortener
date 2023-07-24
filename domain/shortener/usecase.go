package shortener

import (
	"sort"
)

type shortenerImplementation struct {
	repo Repository
}

func NewShortenerService(repo Repository) Service {
	return &shortenerImplementation{
		repo: repo,
	}
}

type Service interface {
	CreateShortenedURL(req *Link) error
	GetByShortenedURL(req string) (*Link, error)
	GetAll(sort string) ([]Link, error)
}

func (s *shortenerImplementation) CreateShortenedURL(req *Link) error {
	err := s.repo.Store(req)
	if err != nil {
		return err
	}

	return s.repo.StoreKey(req)
}

func (s *shortenerImplementation) GetByShortenedURL(req string) (*Link, error) {
	link, err := s.repo.GetByShortenedURL(req)
	if err != nil {
		return nil, err
	}

	link.Clicks++
	err = s.repo.Store(link)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (s *shortenerImplementation) GetAll(sortBy string) ([]Link, error) {
	keys, err := s.repo.GetAllKeys()
	if err != nil {
		return nil, err
	}

	var newKeys []string
	var links []Link
	for _, v := range keys {
		l, err := s.repo.GetByShortenedURL(v)
		if err != nil {
			continue
		}

		links = append(links, *l)
		newKeys = append(newKeys, l.ShortURL)
	}

	err = s.repo.SetKeys(newKeys)
	if err != nil {
		return nil, err
	}

	switch sortBy {
	case "asc":
		sort.SliceStable(links, func(i, j int) bool {
			return links[i].Clicks < links[j].Clicks
		})
	default:
		sort.SliceStable(links, func(i, j int) bool {
			return links[i].Clicks > links[j].Clicks
		})
	}

	return links, nil
}
