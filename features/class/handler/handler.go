package handler

import (
	"kelompok1/immersive-dash/app/middlewares"
	"kelompok1/immersive-dash/features/class"
	"kelompok1/immersive-dash/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ClassHandler struct {
	classService class.ClassServiceInterface
}

func New(service class.ClassServiceInterface) *ClassHandler {
	return &ClassHandler{
		classService: service,
	}
}

func (handler *ClassHandler) GetAllClass(c echo.Context) error {
	result, err := handler.classService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	// mapping dari struct core to struct response
	var classResponse []ClassResponse
	for _, value := range result {
		classResponse = append(classResponse, ClassResponse{
			ID:        value.ID,
			ClassName: value.ClassName,
		})
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", classResponse))

}

func (handler *ClassHandler) CreateClass(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	classInput := new(ClassRequest)
	errBind := c.Bind(&classInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	//mapping dari struct request to struct core
	classCore := RequestToCore(*classInput)
	classCore.UserId = uint(id)
	err := handler.classService.Create(classCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success insert data", nil))
}

func (handler *ClassHandler) GetClassById(c echo.Context) error {
	// Ambil ID proyek dari parameter URL
	id := c.Param("class_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil || idConv <= 0 {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid class ID", nil))
	}

	// Panggil fungsi service untuk mendapatkan detail proyek
	result, err := handler.classService.GetById(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error reading data", nil))
		}
	}

	resultResponse := ClassResponse{
		ID:        result.ID,
		ClassName: result.ClassName,
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success reading data", resultResponse))
}

func (handler *ClassHandler) UpdateClass(c echo.Context) error {
	// Ambil `class_id` dari parameter URL
	id := middlewares.ExtractTokenUserId(c)
	classID, err := strconv.Atoi(c.Param("class_id"))
	if err != nil || classID <= 0 {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid class_id", nil))
	}

	// Ambil data pengguna yang akan diperbarui dari permintaan JSON
	classInput := new(ClassRequest)
	errBind := c.Bind(&classInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Error binding data", nil))
	}

	// Mapping data dari UserRequest ke struct Core
	classCore := RequestToCore(*classInput)
	classCore.UserId = uint(id)

	// Panggil fungsi service untuk memperbarui pengguna
	err = handler.classService.Update(uint(classID), classCore)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "class not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error updating class", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Class updated successfully", nil))
}

func (handler *ClassHandler) DeleteClass(c echo.Context) error {
	// Ambil `user_id` dari parameter URL
	classID, err := strconv.Atoi(c.Param("class_id"))
	if err != nil || classID <= 0 {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid user_id", nil))
	}

	// Panggil fungsi service untuk menghapus pengguna berdasarkan `user_id`
	err = handler.classService.Deletes(uint(classID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "Class not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error deleting Class", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Class deleted successfully", nil))
}
