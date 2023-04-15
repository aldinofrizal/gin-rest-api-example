package request

type Bookmark struct {
	Name       string `form:"name" json:"name" binding:"required"`
	Overview   string `form:"overview" json:"overview" binding:"required"`
	TmdbId     int    `form:"tmdb_id" json:"tmdb_id" binding:"required"`
	PosterPath string `form:"poster_path" json:"poster_path" binding:"required"`
	UserId     int
}
