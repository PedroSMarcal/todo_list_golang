package todo

type (
	RequestTask struct {
		Email       string `json:"email"`
		Description string `json:"description"`
		Author      string `json:"author"`
		LastDay     string `json:"lastDay"`
	}
)
