package todo

// Todo is a to-do item
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Note  string `json:"note"`
	Done  bool   `json:"done"`
}
