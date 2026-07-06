package auth

import (
	"errors"
	"net/http"

	apperr "bookmarks/internal/errors"
	"bookmarks/internal/handler"
	"bookmarks/internal/jwt"
	"bookmarks/internal/models"
	"bookmarks/internal/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginHandler(c *gin.Context, db *gorm.DB) {
	var ur models.UserRequest
	if err := c.ShouldBindJSON(&ur); err != nil {
		handler.RespondError(c, apperr.RecordInvalidError())
		return
	}

	if ur.Email == "" || ur.Password == "" {
		handler.RespondError(c, apperr.ForbiddenError())
		return
	}

	repo := repository.NewUserRepository(db)
	user, err := repo.GetUserByEmail(ur.Email)
	if err != nil {
		var appError *apperr.AppError
		if errors.As(err, &appError) && appError.HTTPStatus == http.StatusNotFound {
			handler.RespondError(c, apperr.UnauthorizedError())
			return
		}
		handler.RespondError(c, err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(ur.Password)); err != nil {
		handler.RespondError(c, apperr.UnauthorizedError())
		return
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		handler.RespondError(c, apperr.InternalError())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
