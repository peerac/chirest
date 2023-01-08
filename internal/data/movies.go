package data

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   int32     `json:"runtime,omitempty"`
}

type MovieModel struct {
	DB *gorm.DB
}

func (m *MovieModel) AddMovie(movie *Movie) error {
	err := m.DB.Create(&movie).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *MovieModel) FetchMovies() ([]Movie, error) {
	var movies []Movie
	err := m.DB.Find(&movies).Error
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (m *MovieModel) FetchMovieByID(id int64) (*Movie, error) {
	var movie *Movie
	result := m.DB.Where("id = ?", id).Find(&movie)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return movie, nil
}

func (m *MovieModel) UpdateMovie(movie *Movie) error {
	result := m.DB.Updates(&movie)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (m *MovieModel) DeleteMovie(id int64) error {
	var movie *Movie
	result := m.DB.Where("id = ?", id).Delete(&movie)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
