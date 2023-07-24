package entity

type Task struct {
	ID   int
	Text string
	Done bool
}

func NewTask (id int, text string, done bool) Task {
	return Task {
		ID:   id,
		Text: text,
		Done: done,
	}
}
