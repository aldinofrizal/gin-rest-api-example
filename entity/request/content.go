package request

type Content struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	ImageUrl    string `form:"image_url" json:"image_url" binding:"required"`
}
