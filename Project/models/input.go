package models

type Input struct {
	Mode        string `json:"mode"`
	Word        string `json:"word"`
	Filepath    string `json:"filepath"`
	Method      string `json:"method"`
	Hash        string `json:"hash"`
	HaswordList bool   `json:"haswordlist"`
	Rules       string `json:"rules"`
}
