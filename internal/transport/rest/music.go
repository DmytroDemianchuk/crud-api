package rest

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	restErrors "github.com/dmytrodemianchuk/crud-api/pkg/rest/errors"

	"github.com/dmytrodemianchuk/crud-api/internal/domain"
	"github.com/gin-gonic/gin"
)

type MusicService interface {
	List(ctx context.Context) (domain.ListMusic, error)
	Get(ctx context.Context, id int) (domain.Music, error)
	Create(ctx context.Context, music domain.Music) (domain.Music, error)
	Update(ctx context.Context, id int, movie domain.Music) (domain.Music, error)
	Delete(ctx context.Context, id int) error
}

type Music struct {
	musicService MusicService
}

func NewMovie(musicService MusicService) *Music {
	return &Music{musicService: musicService}
}

func (m *Music) List(ctx *gin.Context) {
	movies, err := m.musicService.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		return
	}

	ctx.JSON(http.StatusOK, movies)
}

func (m *Music) Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fields := map[string]string{"id": "should be an integer"}
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("validation error", fields))
		return
	}

	movie, err := m.musicService.Get(ctx, id)
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

func (m *Music) Create(ctx *gin.Context) {
	var movie domain.Movie
	if err := ctx.BindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("cannot parse body", nil))
		return
	}

	createdMusic, err := m.musicService.Create(ctx, movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		return
	}

	ctx.JSON(http.StatusCreated, createdMusic)
}

func (m *Music) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fields := map[string]string{"id": "should be an integer"}
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("validation error", fields))
		return
	}

	var movie domain.Music
	if err := ctx.BindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("cannot parse body", nil))
		return
	}

	updatedMovie, err := m.musicService.Update(ctx, id, movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		return
	}

	ctx.JSON(http.StatusOK, updatedMovie)
}

func (m *Music) Delete(ctx *gin.Context) {
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
