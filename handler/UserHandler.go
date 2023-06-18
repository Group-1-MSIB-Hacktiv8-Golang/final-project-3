package handler

import (
	"agolang/project-3/dto"
	"agolang/project-3/entity"
	"agolang/project-3/pkg/errs"
	"agolang/project-3/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

// NewUserHandler membuat instance baru dari userHandler.
// Menerima parameter userService yang merupakan implementasi dari service.UserService.
// Mengembalikan userHandler yang baru dibuat.
func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

// Register digunakan untuk menangani permintaan pendaftaran pengguna baru.
func (uh *userHandler) Register(ctx *gin.Context) {
	var requestBody dto.NewUserRequest

	// Membaca dan mengikat JSON permintaan ke struct requestBody.
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		// Membuat kesalahan UnprocessableEntityError dengan pesan "invalid request body".
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")

		// Mengirim respons JSON dengan kode status dan pesan kesalahan yang sesuai.
		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	// Memanggil userService untuk membuat pengguna baru.
	result, err := uh.userService.CreateNewUser(requestBody)

	if err != nil {
		// Mengirim detail kesalahan ke klien
		ctx.JSON(err.Status(), err)
		return
	}

	// Mengirim respons JSON dengan kode status dan hasil yang diterima dari userService.
	ctx.JSON(http.StatusCreated, result)
}

// Login digunakan untuk menangani permintaan login pengguna.
func (uh *userHandler) Login(ctx *gin.Context) {
	var requestBody dto.LoginUserRequest

	// Membaca dan mengikat JSON permintaan ke struct requestBody.
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		// Membuat kesalahan UnprocessableEntityError dengan pesan "invalid request body".
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")

		// Mengirim respons JSON dengan kode status dan pesan kesalahan yang sesuai.
		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	// Memanggil userService untuk melakukan login pengguna.
	result, err := uh.userService.Login(requestBody)

	if err != nil {
		// Mengirim detail kesalahan ke klien
		ctx.JSON(err.Status(), err)
		return
	}

	// Mengirim respons JSON dengan kode status dan hasil yang diterima dari userService.
	ctx.JSON(http.StatusOK, result)
}

func (uh *userHandler) UpdateUser(ctx *gin.Context) {
	var requestBody dto.UpdateUserRequest

	//Mengambil data dario contex gin
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.Status(), newError)
		return
	}

	// Membaca dan mengikat JSON permintaan ke struct requestBody.
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		// Membuat kesalahan UnprocessableEntityError dengan pesan "invalid request body".
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")

		// Mengirim respons JSON dengan kode status dan pesan kesalahan yang sesuai.
		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	result, err := uh.userService.UpdateUser(userData, &requestBody)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	// Mengirim respons JSON dengan kode status dan hasil yang diterima dari userService.
	ctx.JSON(http.StatusOK, result)

}

func (uh *userHandler) DeleteUser(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.Status(), newError)
		return
	}

	result, err := uh.userService.DeleteUser(userData)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	// Mengirim respons JSON dengan kode status dan hasil yang diterima dari userService.
	ctx.JSON(http.StatusOK, result)

}
