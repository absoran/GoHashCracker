package internal

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/absoran/goproject/models"
	"github.com/absoran/goproject/shared"
)

// this function takes user and checks if there is user with same credentials at database. If there is user with same credentials function returns true, if there isnt false.
func CheckUserExist(user models.User) bool {
	var s bool = false
	querry := shared.SqlCheckExist + user.Username + "')"
	result := db.QueryRow(querry)
	err := result.Scan(&s)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if s {
		return true
	} else {
		return false
	}
}

// GetInputData function takes flags from CLI and parse them into input object. if flag count below 2 then function returns error.
func GetInputData() (input models.Input, err error) {
	if len(os.Args) < 2 {
		return models.Input{}, errors.New("missing argument")
	}
	mode := flag.String("mode", "enc", "app mode -enc,-dec,-web")
	method := flag.String("method", "md5", "Hash type for enc-dec")
	word := flag.String("word", "default", "word for hash single word")
	hash := flag.String("hash", "default", "a hash for crack")
	haswordlist := flag.Bool("haswordlist", false, "wordlist to hash all elements")
	rule := flag.String("rules", "", "rules for improve match chance")
	filepath := flag.String("filepath", "../rockyou.txt", "wordlist location")
	flag.Parse()
	data := models.Input{
		Method:      *method,
		Hash:        *hash,
		Filepath:    *filepath,
		HaswordList: *haswordlist,
		Mode:        *mode,
		Word:        *word,
		Rules:       *rule,
	}
	return data, nil
}

// ProcessInputFromWEB function takes input and defines which function will be used then calls the functions
func ProcessInputFromWEB(input models.Input) string {
	switch input.Mode {
	case "dec":
		{
			if input.HaswordList {
				result := Crack(input)
				if result.IsCracked {
					return result.Word
				} else {
					return ("cannot crack hash in wordlist")
				}
			}
		}
	case "enc":
		{
			switch input.Method {
			case "sha1":
				{

					return HashSHA1(input.Word)
				}
			case "sha256":
				{
					return HashSHA256(input.Word)
				}
			case "sha512":
				{
					return HashSHA512(input.Word)
				}
			case "md5":
				{
					return HashMD5(input.Word)
				}
			}
		}

	}
	return ""
}

// ProcessFlag function takes entered flags in CLI. Basically this application has 3 mode. First encrpyt mode, second decrypt mode and third web mode.
// Encrypt mode appplication waits for 2 flags. These are word to be hashed and method for hasting method.
// Decrypt mode takes 5 flags. This flags are method, hash, rules, haswordlist boolean and wordlist location. Function takes flags and define which function will be used.
// Web mode takes no flags. When web mode activates app will be listening specified port and endpoints.
func ProcessFlag(input models.Input) {
	switch input.Mode {
	case "dec":
		{
			if input.HaswordList {
				result := Crack(input)
				if result.IsCracked {
					PrintResult(result)
				} else {
					fmt.Println("cannot crack hash in wordlist")
				}
			}
		}
	case "enc":
		{
			switch input.Method {
			case "sha1":
				{
					fmt.Println(HashSHA1(input.Word))
				}
			case "sha256":
				{
					fmt.Println(HashSHA256(input.Word))
				}
			case "sha512":
				{
					fmt.Println(HashSHA512(input.Word))
				}
			case "md5":
				{
					fmt.Println(HashMD5(input.Word))
				}
			}
		}
	case "web":
		{
			HandleRequests()
		}
	}
}

// this function takes filepath as parameter and reads all data in parameter indicates. Returns String slice
func ReadWordlist(filepath string) []string {
	wordlist := make([]string, 0)
	file, err := os.Open(filepath)
	CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return wordlist
}

// hash functions hash given string and returns hashed string.
func HashSHA256(data string) string {
	hasher := sha256.New()
	inputToHash := []byte(data)
	_, err := hasher.Write(inputToHash)
	CheckError(err)
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}
func HashSHA1(data string) string {
	hasher := sha1.New()
	inputToHash := []byte(data)
	_, err := hasher.Write(inputToHash)
	CheckError(err)
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}
func HashMD5(data string) string {
	hasher := md5.New()
	inputToHash := []byte(data)
	_, err := hasher.Write(inputToHash)
	CheckError(err)
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}
func HashSHA512(data string) string {
	hasher := sha512.New()
	inputToHash := []byte(data)
	_, err := hasher.Write(inputToHash)
	CheckError(err)
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

// rule functions called via processinput and processinputfromWEB functions.
//rule : %u
func MakeFirstLetterUpper(word string) string {
	upper := strings.Title(word)
	return upper
}

//rule : 1
func Add1ToWord(word string) string {
	return word + "1"
}

//rule : 12
func Add12ToWord(word string) string {
	return word + "12"
}

//rule : 123
func Add123ToWord(word string) string {
	return word + "123"
}

// Crack function take input as parameter and return output. Function basically tries to crack hash in input object with specified flags. For example if hash method is sha1 then checks if there is a rule to apply.
// then iterate through readed wordlist and hash these words. If there is match between hash taken from input and generated hash function takes that word and write it into output's word property.
// If there is a rule append rule to the and of the word for specify rule that used for cracking. Output has wordlist length property and foundAt property, if hash cracked these 2 values will be set to founded location and wordlist's length.
func Crack(input models.Input) models.Output {
	readedwords := ReadWordlist(input.Filepath)
	size := len(readedwords)
	output := models.Output{IsCracked: false}

	switch input.Method {
	case "sha1":
		{
			if input.Rules == "" {
				for index, word := range readedwords {
					hash := HashSHA1(word)
					if hash == input.Hash {
						output.FoundAt = index
						output.Hash = hash
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hash != input.Hash {
						return output
					}
				}
			}
			if input.Rules == "%u" {
				for index, word := range readedwords {
					wordwithrule := MakeFirstLetterUpper(word)
					hashwithrule := HashSHA1(wordwithrule)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
			if input.Rules == "1" {
				for index, word := range readedwords {
					wordwithrule := Add1ToWord(word)
					hashwithrule := HashSHA1(wordwithrule)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
			if input.Rules == "12" {
				for index, word := range readedwords {
					wordwithrule := Add12ToWord(word)
					hashwithrule := HashSHA1(wordwithrule)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
			if input.Rules == "123" {
				for index, word := range readedwords {
					wordwithrule := Add123ToWord(word)
					hashwithrule := HashSHA1(wordwithrule)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
		}
	case "sha256":
		{
			if input.Rules == "" {
				for index, word := range readedwords {
					hash := HashSHA256(word)
					if hash == input.Hash {
						output.FoundAt = index
						output.Hash = hash
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hash != input.Hash {
						return output
					}
				}
			}

			if input.Rules == "%u" {
				for index, word := range readedwords {
					wordwithrule := MakeFirstLetterUpper(word)
					hashwithrule := HashSHA256(wordwithrule)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
			if input.Rules == "1" {
				for index, word := range readedwords {
					wordwithrule := Add1ToWord(word)
					hashwithrule := HashSHA256(wordwithrule)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
			if input.Rules == "12" {
				for index, word := range readedwords {
					wordwithrule := Add12ToWord(word)
					hashwithrule := HashSHA256(wordwithrule)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
			if input.Rules == "123" {
				for index, word := range readedwords {
					wordwithrule := Add123ToWord(word)
					hashwithrule := HashSHA256(wordwithrule)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
		}
	case "sha512":
		{
			if input.Rules == "" {
				for index, word := range readedwords {
					hash := HashSHA512(word)
					if hash == input.Hash {
						output.FoundAt = index
						output.Hash = hash
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hash != input.Hash {
						return output
					}
				}
			}

			if input.Rules == "u" {
				for index, word := range readedwords {
					wordwithrule := MakeFirstLetterUpper(word)
					hashwithrule := HashSHA512(wordwithrule)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
			if input.Rules == "1" {
				for index, word := range readedwords {
					wordwithrule := Add1ToWord(word)
					hashwithrule := HashSHA512(wordwithrule)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
			if input.Rules == "12" {
				for index, word := range readedwords {
					wordwithrule := Add12ToWord(word)
					hashwithrule := HashSHA512(wordwithrule)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
			if input.Rules == "123" {
				for index, word := range readedwords {
					wordwithrule := Add123ToWord(word)
					hashwithrule := HashSHA512(wordwithrule)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
		}
	case "md5":
		{
			if input.Rules == "" {
				for index, word := range readedwords {
					hash := HashMD5(word)
					if hash == input.Hash {
						output.FoundAt = index
						output.Hash = hash
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hash != input.Hash {
						return output
					}
				}
			}

			if input.Rules == "%u" {
				for index, word := range readedwords {
					wordwithrule := MakeFirstLetterUpper(word)
					hashwithrule := HashMD5(wordwithrule)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
			if input.Rules == "1" {
				for index, word := range readedwords {
					wordwith1 := Add1ToWord(word)
					hashwithrule := HashMD5(wordwith1)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
			if input.Rules == "12" {
				for index, word := range readedwords {
					wordwith12 := Add12ToWord(word)
					hashwithrule := HashMD5(wordwith12)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
			if input.Rules == "123" {
				for index, word := range readedwords {
					wordwith123 := Add123ToWord(word)
					hashwithrule := HashMD5(wordwith123)
					if hashwithrule == input.Hash {
						output.FoundAt = index
						output.Hash = hashwithrule
						output.Word = word
						output.WordlistSize = size
						output.Rules = input.Rules
						output.IsCracked = true
						return output
					}
					if index == len(readedwords)-1 && hashwithrule != input.Hash {
						return output
					}
				}
			}
		}
	}
	return output
}
