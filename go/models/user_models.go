package models

type UserCreate struct {
	Name string `uri:"name" binding:"required"`
}

type UserDetail struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}
