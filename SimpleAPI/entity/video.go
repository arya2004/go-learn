package entity

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}

// xml:"title" form:"title" "validate":"email" binding:"required"
type Video struct {
	Title       string `json:"title" biding:"min=2,max=10"`
	Description string `json:"description" biding:"max=20"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}