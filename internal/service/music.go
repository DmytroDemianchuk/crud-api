package service

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/dmytrodemianchuk/crud-api/internal/domain"
	"github.com/gin-gonic/gin"
)

type MusicService interface {
	List(ctx context.Context) (domain.ListMusic, error)
	Get(ctx context.Context, id int) (domain.Music, error)
	Create(ctx context.Context, movie domain.Music) (domain.Music, error)
	Update(ctx context.Context, id int, movie domain.Music) (domain.Music, error)
	Delete(ctx context.Context, id int) error
}

type Music struct {
	musicService MusicService
}

func NewMusic(musicService MusicService) *Music {
	return &Music{musicService: musicService}
}

func (m Music) List(ctx *gin.Context) {
	movies, err := m.musicService.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		return
	}

	ctx.JSON(http.StatusOK, movies)
}

func (m Movie) Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fields := map[string]string{"id": "should be an integer"}
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("validation error", fields))
		return
	}

	movie, err := m.movieService.Get(ctx, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			ctx.JSON(http.StatusNotFound, restErrors.NewNotFoundErr("movie not found"))
		default:
			ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		}

		return
	}

	ctx.JSON(http.StatusOK, movie)
}

func (m Music) Create(ctx *gin.Context) {
	var movie domain.Movie
	if err := ctx.BindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("cannot parse body", nil))
		return
	}

	createdMusic, err := m.movieService.Create(ctx, movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		return
	}

	ctx.JSON(http.StatusCreated, createdMusic)
}

func (m Music) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fields := map[string]string{"id": "should be an integer"}
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("validation error", fields))
		return
	}

	var music domain.Music
	if err := ctx.BindJSON(&music); err != nil {
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("cannot parse body", nil))
		return
	}

	updatedMusic, err := m.movieService.Update(ctx, id, movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		return
	}

	ctx.JSON(http.StatusOK, updatedMusic)
}

func (m Music) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fields := map[string]string{"id": "should be an integer"}
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("validation error", fields))
		return
	}

	if err := m.musicService.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
