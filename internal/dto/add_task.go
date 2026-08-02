package dto

type CreateTaskInput struct {
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type CreateTaskOutput struct {
	Id string `json:"id"`
}
