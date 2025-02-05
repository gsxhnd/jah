package service

import (
	"github.com/gsxhnd/jaha/server/db/database"
	"github.com/gsxhnd/jaha/server/model"
	"github.com/gsxhnd/jaha/utils"
)

type MovieTagService interface {
	CreateMovieTags(data []model.MovieTag) error
	DeleteMovieTags(movieIds []uint) error
	GetMovieTags(movieId uint) ([]model.MovieTag, error)
}

type movieTagService struct {
	logger utils.Logger
	db     database.Driver
}

func NewMovieTagService(l utils.Logger, db database.Driver) MovieTagService {
	return movieTagService{
		logger: l,
		db:     db,
	}
}

// CreateMovieTags implements MovieTagService.
func (s movieTagService) CreateMovieTags(movieTags []model.MovieTag) error {
	return s.db.CreateMovieTags(movieTags)
}

// DeleteMovieTags implements MovieTagService.
func (s movieTagService) DeleteMovieTags(ids []uint) error {
	return s.db.DeleteMovieTags(ids)
}

// GetMovieTags implements MovieTagService.
func (s movieTagService) GetMovieTags(movieId uint) ([]model.MovieTag, error) {
	return s.db.GetMovieTagByMovieId(movieId)
}
