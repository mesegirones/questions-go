package storage

import (
	"context"
	"questions-go/domain"
	"strconv"
)

type QuestionRepository struct {
	Logger    LoggerProxy
	Questions []*domain.Question
	Answers   []*domain.UserAnswer
}

func NewQuestionRepository(config *StorageConfig) *QuestionRepository {
	return &QuestionRepository{
		Logger:    config.Logger,
		Questions: config.Questions,
		Answers:   config.Answer,
	}
}

func (r *QuestionRepository) List(ctx context.Context) ([]*domain.Question, error) {
	//TODO: handle pagination??
	return r.Questions, nil
}

func (r *QuestionRepository) SaveAnswers(ctx context.Context, userAnswer domain.UserAnswer) error {
	lastUser := r.Answers[len(r.Answers)]
	userId, err := strconv.Atoi(lastUser.UserId)
	if err != nil {
		return domain.ErrInternalServerError
	}
	userId++
	userAnswer.UserId = strconv.Itoa(userId)
	r.Answers = append(r.Answers, &userAnswer)
	return nil
}

func (r *QuestionRepository) AnswersList(ctx context.Context) ([]*domain.UserAnswer, error) {
	//TODO: handle pagination??
	return r.Answers, nil
}
