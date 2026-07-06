package handler

import (
	"errors"
	"net/http"

	apperr "bookmarks/internal/errors"

	"github.com/gin-gonic/gin"
)

func RespondError(c *gin.Context, err error) {
	var appError *apperr.AppError
	if errors.As(err, &appError) {
		c.JSON(appError.HTTPStatus, apperr.Response{
			Error: apperr.ErrorBody{
				Code:    appError.Code,
				Message: appError.Message,
			},
		})
		return
	}

	c.JSON(http.StatusInternalServerError, apperr.Response{
		Error: apperr.ErrorBody{
			Code:    apperr.InternalError().Code,
			Message: apperr.InternalError().Message,
		},
	})
}
