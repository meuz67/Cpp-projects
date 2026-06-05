package models

type Track struct {
	ID     uint64 `json:"id"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
}
type CreateTrack struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
}
