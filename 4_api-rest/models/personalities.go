package models

type Personality struct {
	Id      int    `json:"id"`      // json:"id" is a tag
	Name    string `json:"name"`    // json:"name" is a tag
	History string `json:"history"` // json:"history" is a tag
}

var Personalities []Personality
