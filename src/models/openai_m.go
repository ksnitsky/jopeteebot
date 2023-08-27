package models

// TODO: Move it to another package (same as telegram_m)

type GPTResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Choices []GPTChoice `json:"choices"`
}

type GPTChoice struct {
	Index        int      `json:"index"`
	Message      GPTMessage  `json:"message"`
	FinishReason string   `json:"finish_reason"`
}

type GPTMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
