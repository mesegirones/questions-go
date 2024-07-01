package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"questions-go/domain"

	"github.com/gin-gonic/gin"
)

type QuestionService interface {
	GetQuestionList(ctx context.Context) ([]domain.Question, error)
	SubmitQuestionAnswers(ctx context.Context, input []*domain.AnswersInput) (*domain.UserAnswer, error)
	GetStatistics(ctx context.Context, userId string) (*domain.AnswerStatistics, error)
}

type QuestionHandler struct {
	Service QuestionService
}

func NewQuestionHandler(r *gin.Engine, service QuestionService) {
	handler := &QuestionHandler{
		Service: service,
	}

	g := r.Group("/question")

	g.GET("/list", handler.GetQuestionList)
	g.GET("/statistics/:userId", handler.GetStatistics)

	g.POST("/answers", handler.SubmitQuestionAnswers)

}

func (h *QuestionHandler) GetQuestionList(c *gin.Context) {
	response, err := h.Service.GetQuestionList(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageResponse{Message: err.Error()})
	}
	c.JSON(http.StatusAccepted, response)
}

func (h *QuestionHandler) SubmitQuestionAnswers(c *gin.Context) {
	var input []*domain.AnswersInput
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrBadParamInput)
		return
	}
	response, err := h.Service.SubmitQuestionAnswers(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageResponse{Message: err.Error()})
	}
	c.JSON(http.StatusAccepted, response)
}

func (h *QuestionHandler) GetStatistics(c *gin.Context) {
	userId := c.Param("userId")
	response, err := h.Service.GetStatistics(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageResponse{Message: err.Error()})
	}
	c.JSON(http.StatusAccepted, response)
}
