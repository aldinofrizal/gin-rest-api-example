package request

type Restaurant struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Address  string `form:"address" json:"address"  binding:"required"`
	ImageUrl string `form:"image_url" json:"image_url" binding:"required"`
}
