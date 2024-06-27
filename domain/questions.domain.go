package domain

type Question struct {
	Id       string            `json:"id"`
	Question string            `json:"question"`
	Options  []QuestionOptions `json:"options"`
}

type QuestionOptions struct {
	Id        string `json:"id"`
	Answer    string `json:"answer"`
	IsCorrect bool   `json:"isCorrect"`
}
