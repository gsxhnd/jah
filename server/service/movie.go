package service

import (
	"github.com/gsxhnd/jaha/server/db/database"
	"github.com/gsxhnd/jaha/server/model"
	"github.com/gsxhnd/jaha/server/storage"
	"github.com/gsxhnd/jaha/utils"
)

type MovieService interface {
	CreateMovies([]model.Movie) error
	DeleteMovies([]uint) error
	UpdateMovie(*model.Movie) error
	GetMovies(*database.Pagination) ([]model.Movie, error)
	GetMovieInfo(string) (*model.MovieInfo, error)
	SearchMoviesByCode(string) ([]model.Movie, error)
	UploadMovieCover(code, filename string, file []byte) error
}

type movieService struct {
	logger  utils.Logger
	db      database.Driver
	storage storage.Storage
}

func NewMovieService(l utils.Logger, db database.Driver, s storage.Storage) MovieService {
	return movieService{
		logger:  l,
		db:      db,
		storage: s,
	}
}

// CreateMovies implements MovieService.
func (s movieService) CreateMovies(movies []model.Movie) error {
	return s.db.CreateMovies(movies)
}

// DeleteMovies implements MovieService.
func (s movieService) DeleteMovies(ids []uint) error {
	return s.db.DeleteMovies(ids)
}

// UpdateMovie implements MovieService.
func (s movieService) UpdateMovie(m *model.Movie) error {
	return s.db.UpdateMovie(m)
}

// GetMovies implements MovieService.
func (s movieService) GetMovies(p *database.Pagination) ([]model.Movie, error) {
	return s.db.GetMovies(p)
}

// GetMovies implements MovieService.
func (s movieService) SearchMoviesByCode(code string) ([]model.Movie, error) {
	return s.db.GetMovies(nil, "code", code)
}

func (s movieService) GetMovieInfo(code string) (*model.MovieInfo, error) {
	var data model.MovieInfo
	movie, err := s.db.GetMovieByCode(code)
	if err != nil {
		return nil, err
	}
	data.Movie = *movie

	movieTags, err := s.db.GetMovieTagByMovieId(movie.Id)
	if err != nil {
		return nil, err
	}

	data.Tags = movieTags

	movieActor, err := s.db.GetMovieActorsByMovieId(movie.Id)
	if err != nil {
		return nil, err
	}
	data.Actors = movieActor

	return &data, nil
}

func (s movieService) UploadMovieCover(code, filename string, file []byte) error {
	movie, err := s.db.GetMovieByCode(code)
	if err != nil {
		return err
	}

	if err := s.storage.SaveImage(file, "movie", movie.Id, filename); err != nil {
		return err
	}
	return nil
}
