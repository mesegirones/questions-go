package question

import (
	"context"
	"questions-go/domain"
)

type QuestionRepository interface {
	List(ctx context.Context) ([]domain.Question, error)
	SaveAnswers(ctx context.Context, input domain.UserAnswer) (*domain.UserAnswer, error)
	AnswersList(ctx context.Context) ([]domain.UserAnswer, error)
}

type LoggerProxy interface {
	Error(v ...interface{})
}

type Service struct {
	QuestionRepository QuestionRepository
	Logger             LoggerProxy
}

func NewService(loggerProxy LoggerProxy, questionRepository QuestionRepository) *Service {
	return &Service{
		Logger:             loggerProxy,
		QuestionRepository: questionRepository,
	}
}

func (s *Service) GetQuestionList(ctx context.Context) ([]domain.Question, error) {
	list, err := s.QuestionRepository.List(ctx)
	if err != nil {
		s.Logger.Error(err)
		return nil, domain.ErrInternalServerError
	}
	return list, nil
}

func (s *Service) SubmitQuestionAnswers(ctx context.Context, input []*domain.AnswersInput) (*domain.UserAnswer, error) {
	questionList, err := s.QuestionRepository.List(ctx)
	if err != nil {
		s.Logger.Error(err)
		return nil, domain.ErrInternalServerError
	}

	helperQuestions := make(map[string]*domain.QuestionOptions)
	for _, question := range questionList {
		for _, questionOption := range question.Options {
			if questionOption.IsCorrect {
				helperQuestions[question.Id] = &questionOption
			}
		}
	}

	var userAnswers []domain.Answer
	totalCorrect := 0
	for _, answer := range input {
		correctAnswer := helperQuestions[answer.QuestionId]
		if correctAnswer == nil {
			return nil, domain.ErrBadParamInput
		}
		userAnswer := domain.Answer{
			QuestionId:        answer.QuestionId,
			SubmittedAnswerId: answer.SubmittedAnswerId,
			IsCorrect:         answer.SubmittedAnswerId == correctAnswer.Id,
		}
		userAnswers = append(userAnswers, userAnswer)

		if userAnswer.IsCorrect {
			totalCorrect++
		}
	}

	res, err := s.QuestionRepository.SaveAnswers(ctx, domain.UserAnswer{
		UserId:       "-1",
		Answers:      userAnswers,
		TotalCorrect: totalCorrect,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) GetStatistics(ctx context.Context, userId string) (*domain.AnswerStatistics, error) {
	answersList, err := s.QuestionRepository.AnswersList(ctx)
	if err != nil {
		s.Logger.Error(err)
		return nil, domain.ErrInternalServerError
	}

	helperAnswersPerUser := make(map[string]*domain.UserAnswer)
	for _, answers := range answersList {
		helperAnswersPerUser[answers.UserId] = &answers
	}

	var supTotal, equalTotal, infTotal int

	userAnswer := helperAnswersPerUser[userId]
	if userAnswer == nil {
		return nil, domain.ErrNotFound
	}

	for id, data := range helperAnswersPerUser {
		if id == userId || data.TotalCorrect == userAnswer.TotalCorrect {
			equalTotal++
		}
		if data.TotalCorrect > userAnswer.TotalCorrect {
			supTotal++
		}
		if data.TotalCorrect < userAnswer.TotalCorrect {
			infTotal++
		}
	}

	response := &domain.AnswerStatistics{
		SuperiorPercent: getPercent(len(answersList), supTotal),
		EqualPercent:    getPercent(len(answersList), equalTotal),
		InferiorPercent: getPercent(len(answersList), infTotal),
	}

	return response, nil
}

func getPercent(total int, part int) int {
	return part * 100 / total
}
