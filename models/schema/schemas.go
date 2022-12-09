package schema

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"Author" binding:"required"`
}

type BookUpdateInput struct {
	Title  string `json:"title"`
	Author string `json:"Author"`
}
