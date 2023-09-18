package handlers

import (
	"auth_api/auth"
	"auth_api/config"
	"auth_api/entity"
	"auth_api/errorConst"
	"auth_api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	Usecase *usecase.UserUsecase
	config  config.Config
}

func NewUsersHandler(usecase *usecase.UserUsecase, config config.Config) *UsersHandler {
	return &UsersHandler{Usecase: usecase, config: config}
}

func (u *UsersHandler) Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	passwordConfirm := ctx.PostForm("password_confirm")

	if name == "" || email == "" || password == "" || passwordConfirm == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "incomplete form"})
		return
	}
	if password != passwordConfirm {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "password and confirmation password don't match"})
		return
	}
	if !entity.IsPasswordValid(password) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errorConst.PASSWORD_INVALID})
		return
	}

	hashedPassword := auth.HashPassword(password)

	input := usecase.UserInput{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}
	output, err := u.Usecase.CreateUser(input)
	if err != nil {
		if err.Error() == errorConst.EMAIL_REGISTERED {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, output)
}

func (u *UsersHandler) SignInUser(ctx *gin.Context) {
	email, ok := ctx.GetPostForm("email")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email"})
		return
	}
	password, ok := ctx.GetPostForm("password")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid password"})
		return
	}

	user, err := u.Usecase.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or password"})
		return
	}

	if err := auth.ComparePassword(user.Password, password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or password"})
		return
	}

	accessToken, err := auth.CreateToken(u.config.AccessTokenExpiresIn, user.ID, u.config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refreshToken, err := auth.CreateToken(u.config.RefreshTokenExpiresIn, user.ID, u.config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken, "refresh_token": refreshToken})
}

func (u *UsersHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	// the app should destroy the jwt tokens in storage
	// if needed, we should create a blacklist in redis to store de invalidated tokens

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (u *UsersHandler) RefreshToken(ctx *gin.Context) {

}

func (u *UsersHandler) GetMe(ctx *gin.Context) {
	user := ctx.MustGet("currentUser").(*entity.UserOutput)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": user})
}
