package handler

import (
	"kelompok1/immersive-dash/features/mentee"
	"kelompok1/immersive-dash/helpers"
	"net/http"
	"strconv"
	"strings"

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
	//var mentee mentee.CoreMentee
	NewUser := new(MenteeRequest)
	//mendapatkan data yang dikirm oleh FE melalui request
	err := c.Bind(&NewUser)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helpers.WebResponse(http.StatusUnprocessableEntity, "error bind data", nil))
	}
	input := RequestToCore(*NewUser)
	menteeId, err := handler.menteeService.Create(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data, "+err.Error(), nil))
	}
	input.ID = menteeId
	newMenteeResponse := CoreToCreateDeleteResponse(input)
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success create mentee", newMenteeResponse))
}

func (handler *MenteeHandler) GetAllMentee(c echo.Context) error {
	// Buat query map dengan kriteria pencarian
	nameQuery := map[string]any{}
	// 	"class_name":     className,
	// 	"status_name":    statusName,
	// 	"education_type": educationType,
	// }
	className := c.QueryParam("class_name")
	if className != "" {
		nameQuery["class_name"] = className
	}
	statusName := c.QueryParam("status_name")
	if statusName != "" {
		nameQuery["status_name"] = statusName
	}
	educationType := c.QueryParam("education_type")
	if educationType != "" {
		nameQuery["education_type"] = educationType
	}
	result, err := handler.menteeService.GetAll(nameQuery)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}
	var menteesResponse []GetAllMenteeResponse
	for _, v := range result {
		menteesResponse = append(menteesResponse, CoreToGetAllResponse(v))
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

	resultResponse := CoreToGetByIdResponse(result)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusFound, "success read data", resultResponse))
}
func (handler *MenteeHandler) GetMenteeFeedback(c echo.Context) error {
	idParam := c.Param("mentee_id")
	idConv, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "mentee id is not valid", nil))
	}
	result, err := handler.menteeService.GetById(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	resultResponse := CoreToGetMenteeFeedbackResponse(result)
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
	resultResponse := CoreToUpdateResponse(updatedMentee)
	// Kirim respon JSON
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "mentee updated successfully", resultResponse))
}

func (handler *MenteeHandler) DeleteMenteeById(c echo.Context) error {
	idParam := c.Param("mentee_id")
	idConv, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "mentee id is not valid", nil))
	}
	mentee, err := handler.menteeService.GetById(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "Mentee not found", nil))
	}
	err = handler.menteeService.Delete(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "Mentee not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "Mentee not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error deactivating data, "+err.Error(), nil))
	}
	menteeResponse := CoreToCreateDeleteResponse(mentee)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Deactivate specific mentee", menteeResponse))
}
