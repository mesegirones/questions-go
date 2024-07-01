package domain

type Question struct {
	Id       string            `json:"id"`
	Question string            `json:"question"`
	Options  []QuestionOptions `json:"options"`
}

type QuestionOptions struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	IsCorrect bool   `json:"isCorrect"`
}

type QuestionOutput struct {
	Id       string                  `json:"id"`
	Question string                  `json:"question"`
	Options  []QuestionOptionsOutput `json:"options"`
}

type QuestionOptionsOutput struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type QuestionResult struct {
	QuestionId        string `json:"questionId"`
	SubmittedAnswerId string `json:"submittedAnswerId"`
	CorrectAnswerId   string `json:"correctAnswerId"`
}

type AnswersInput struct {
	QuestionId        string `json:"questionId"`
	SubmittedAnswerId string `json:"submittedAnswerId"`
}

type UserAnswer struct {
	UserId       string   `json:"userId"`
	TotalCorrect int      `json:"toalCorrect"`
	Answers      []Answer `json:"results"`
}

type Answer struct {
	QuestionId        string `json:"questionId"`
	SubmittedAnswerId string `json:"submittedAnswerId"`
	IsCorrect         bool   `json:"isCorrect"`
}

type AnswerStatistics struct {
	SuperiorPercent int `json:"superiorPercent"`
	EqualPercent    int `json:"equalPercent"`
	InferiorPercent int `json:"inferiorPercent"`
}
