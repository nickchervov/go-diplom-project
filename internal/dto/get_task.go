package dto

type GetTaskInput struct {
	Id string `json:"id"`
}

type GetTaskOutput struct {
	Id      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}
