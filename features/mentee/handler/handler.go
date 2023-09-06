package handler

import (
	"kelompok1/immersive-dash/features/mentee"
	"kelompok1/immersive-dash/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MenteeHandler struct {
	menteeService mentee.MenteeServiceInterface
}

func New(service mentee.MenteeServiceInterface) *MenteeHandler {
	return &MenteeHandler{
		menteeService: service,
	}
}

func (handler *MenteeHandler) CreateMentee(c echo.Context) error {
	newMenteeData := new(MenteeRequest)

	//mendapatkan data yang dikirm oleh FE melalui request
	err := c.Bind(&newMenteeData)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helpers.WebResponse(http.StatusUnprocessableEntity, "error bind data", nil))
	}

	//mapping from mentee req to core mentee
	newMentee := RequestToCore(*newMenteeData)
	_, err = handler.menteeService.Create(newMentee)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data, "+err.Error(), nil))
	}
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success create mentee", nil))
}

func (handler *MenteeHandler) GetAllMentee(c echo.Context) error {
	result, err := handler.menteeService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}
	var menteesResponse []MenteeResponse
	for _, v := range result {
		menteesResponse = append(menteesResponse, CoreToResponse(v))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusFound, "success read data", menteesResponse))
}
func (handler *MenteeHandler) GetMenteeById(c echo.Context) error {
	idParam := c.Param("mentee_id")
	idConv, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "mentee id is not valid", nil))
	}
	result, err := handler.menteeService.GetById(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	resultResponse := CoreToResponse(result)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusFound, "success read data", resultResponse))
}

func (handler *MenteeHandler) UpdateMenteeById(c echo.Context) error {
	idParam := c.Param("mentee_id")
	idConv, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "mentee id is not valid", nil))
	}
	// Mengambil data input dari permintaan
	menteeInput := MenteeRequest{}
	errBind := c.Bind(&menteeInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}
	//Mapping mentee request to core mentee
	coreMentee := RequestToCore(menteeInput)

	// Melakukan pembaruan data proyek di service
	err = handler.menteeService.Update(uint(idConv), coreMentee)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error update data, "+err.Error(), nil))
	}

	// Mendapatkan data proyek yang telah diperbarui untuk response
	updatedMentee, err := handler.menteeService.GetById(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "mentee not found", nil))
	}
	// Mapping updated mentee to MenteeResponse
	resultResponse := CoreToResponse(updatedMentee)
	// Kirim respon JSON
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "mentee updated successfully", resultResponse))
}

func (handler *MenteeHandler) DeleteMenteeById(c echo.Context) error {
	idParam := c.Param("mentee_id")
	idConv, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "mentee id is not valid", nil))
	}
	err = handler.menteeService.Delete(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error delete data, "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success delete mentee", nil))
}

// func (handler *MenteeHandler) CreateMentee(c echo.Context) error {}
