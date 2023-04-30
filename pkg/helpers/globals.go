package helpers

var Url = "https://api.openai.com/v1/chat/completions"

type ApiReq struct {
	Model string `json:"model"`
	Messages []Message `json:"messages"`
	Max_tokens int `json:"max_tokens"`
}

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type ApiRes struct {
	Id string `json:"id"`
	Object string `json:"object"`
	Created int `json:"created"`
	Choices []Choice `json:"choices"`
	Usage Usage `json:"usage"`
}

type Choice struct {
	Index int `json:"index"`
	Message Message `json:"message"`
	Finish_reason string `json:"finish_reason"`
}

type Usage struct {
	Prompt_tokens int `json:"prompt_tokens"`
	Completion_tokens int `json:"completion_tokens"`
	Total_tokens int `json:"total_tokens"`
}

