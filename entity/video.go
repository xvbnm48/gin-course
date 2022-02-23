package entity

type Person struct {
	FirtsName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Age       int8   `json:"age" binding:"required"`
	Email     string `json:"email" validate:"required, email"`
}

type Video struct {
	Title       string `json:"title" binding:"min=1,max=100"`
	Description string `json:"description" binding:"min=1,max=100"`
	URL         string `json:"url" binding:"required"`
	Author      Person `json:"author" binding:"required"`
}
