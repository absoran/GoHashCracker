package models

type Output struct {
	WordlistSize int
	FoundAt      int
	Word         string
	Hash         string
	IsCracked    bool
	Rules        string
}
