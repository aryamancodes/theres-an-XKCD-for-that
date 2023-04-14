package model

type Search struct {
	Query       string `form:"q" binding:"required"`
	Autocorrect bool   `form:"autocorrect"`
}

type Suggest struct {
	Query string `form:"q" binding:"required"`
}
