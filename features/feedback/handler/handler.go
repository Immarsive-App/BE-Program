package handler

import (
	"kelompok1/immersive-dash/app/middlewares"
	"kelompok1/immersive-dash/features/feedback"
	"kelompok1/immersive-dash/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type FeedbackHandler struct {
	feedbackService feedback.FeedbackServiceInterface
}

func New(service feedback.FeedbackServiceInterface) *FeedbackHandler {
	return &FeedbackHandler{
		feedbackService: service,
	}
}

func (handler *FeedbackHandler) GetAllFeedback(c echo.Context) error {
	result, err := handler.feedbackService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	// mapping dari struct core to struct response
	var feedbackResponse []FeedbackResponse
	for _, value := range result {
		feedbackResponse = append(feedbackResponse, FeedbackResponse{
			ID:     value.ID,
			Note:   value.Note,
			UserId: value.UserId,
		})
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", feedbackResponse))

}

func (handler *FeedbackHandler) CreateFeedback(c echo.Context) error {
	// Ambil ID pengguna dari token JWT yang terkait dengan permintaan
	userInput := new(FeedbackRequest)
	userId := middlewares.ExtractTokenUserId(c)

	// Mendapatkan data yang dikirim oleh FE melalui permintaan body

	if err := c.Bind(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}

	// Mapping data dari FeedbackRequest ke CoreFeedback
	feedbackCore := RequestToCore(*userInput)
	feedbackCore.UserId = uint(userId)
	// Panggil fungsi service untuk membuat feedback
	// Panggil fungsi service untuk membuat feedback
	result, err := handler.feedbackService.Create(feedbackCore, uint(userId))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}

	// Mapping data hasil operasi ke dalam response
	response := CoreToResponse(result)
	// 	ID:       result.ID,
	// 	Note:     result.Note,
	// 	UserId:   result.UserId,
	// 	FullName: result.FullName, // Sesuaikan dengan nama properti yang sesuai dengan data user.
	// 	Status:   result.Status,   // Sesuaikan dengan nama properti yang sesuai dengan data status.
	// }

	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success insert data", response))
}

func (handler *FeedbackHandler) UpdateFeedback(c echo.Context) error {
	// Ambil `user_id` dari token
	userID := middlewares.ExtractTokenUserId(c)

	// Ambil `feedback_id` dari parameter URL
	feedbackID, err := strconv.Atoi(c.Param("feedback_id"))
	if err != nil || feedbackID <= 0 {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid feedback_id", nil))
	}

	// Ambil data feedback yang akan diperbarui dari permintaan JSON
	feedbackInput := new(FeedbackRequest)
	errBind := c.Bind(&feedbackInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Error binding data", nil))
	}

	// Mapping data dari FeedbackRequest ke struct Core
	feedbackCore := RequestToCore(*feedbackInput)
	feedbackCore.UserId = uint(userID)

	// Panggil fungsi service untuk memperbarui feedback
	err = handler.feedbackService.Update(uint(feedbackID), feedbackCore)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "Feedback not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error updating feedback", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Feedback updated successfully", nil))
}

func (handler *FeedbackHandler) DeleteFeedback(c echo.Context) error {
	// Ambil `user_id` dari parameter URL
	feedbackID, err := strconv.Atoi(c.Param("feedback_id"))
	if err != nil || feedbackID <= 0 {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid user_id", nil))
	}

	// Panggil fungsi service untuk menghapus pengguna berdasarkan `user_id`
	err = handler.feedbackService.Deletes(uint(feedbackID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "Feedback not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error deleting Feedback", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Feedback deleted successfully", nil))
}
