package ggst_api

type PlayerFilter int

const (
	AllPlayers PlayerFilter = 0
	Self       PlayerFilter = 1
	Follow     PlayerFilter = 2
	Rival      PlayerFilter = 3
	Favorite   PlayerFilter = 4
)
