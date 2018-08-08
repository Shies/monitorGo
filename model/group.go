package model

type Group struct {
	Id           int64
	Name         string
	IsUserAdmin  int
	IsGroupAdmin int
	IsConfAdmin  int
}
