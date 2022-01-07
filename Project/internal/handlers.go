package internal

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/absoran/goproject/models"
	"github.com/absoran/goproject/shared"
	jwt "github.com/dgrijalva/jwt-go"
)

var key = []byte("super_secret_key_999")
var view *template.Template

// init function runs at build initialize view object for template operations
func init() {
	view = template.Must(template.ParseGlob("views/*"))
}

// homepage handler waits for valid jwt token to open. When user login via signin page valid jwt token being added to cookies. Homepage handler checks is there a cookie,
// if there is cookie checks if its valid. If there is token and its valid than user can see homepage.
func Homepage(w http.ResponseWriter, r *http.Request) {
	var result models.HtmlVar
	if r.Method == "POST" {
		encryptMethod := strings.ToLower(r.FormValue("method_id1"))
		decryptMethod := strings.ToLower(r.FormValue("method_id2"))
		wordToHash := r.FormValue("Word")
		HashToCrack := r.FormValue("Hash")
		rules := r.FormValue("rules_id")
		rulesToCrack := ParseRule(rules)

		//encrypt
		if encryptMethod != "" {
			input := models.Input{
				Mode:   "enc",
				Word:   wordToHash,
				Method: encryptMethod,
			}
			result.Result2 = ProcessInputFromWEB(input)
			view.ExecuteTemplate(w, "homepage.gohtml", result)

		} else { //decrypt
			input := models.Input{
				Mode:        "dec",
				Method:      decryptMethod,
				Hash:        HashToCrack,
				Rules:       rulesToCrack,
				HaswordList: true,
				Filepath:    "rockyou.txt",
			}
			result.Result = ProcessInputFromWEB(input)
			view.ExecuteTemplate(w, "homepage.gohtml", result)
		}
	} else {
		view.ExecuteTemplate(w, "homepage.gohtml", nil)
	}

	// Signup handler takes user credentials and checks if there user already exist with entered username. If there is not user with entered username handler takes credendials,
	// and build user object with entered credentials. When user object append to the database user redirects to SignIn page.
}
func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit SignUp")
	if r.Method == "POST" {

		usrname := r.FormValue("username")
		email := r.FormValue("mail")
		passwd := r.FormValue("password")
		user := models.User{
			Username: usrname,
			Email:    email,
			Password: passwd,
		}
		if !CheckUserExist(user) {
			passwdhash := HashSHA512(user.Password)
			_, err := db.Exec(shared.SqlInsert_user, user.Username, user.Email, passwdhash)
			if err != nil {
				log.Fatal(err)
				http.Redirect(w, r, "/signup.html", http.StatusInternalServerError)
			}
		}
		http.Redirect(w, r, "/api/signin", http.StatusSeeOther)
	} else {
		view.ExecuteTemplate(w, "signup.html", nil)
	}
}

// SignIn handler takes entered credentials and checks database if there is user exist with entered credentials. If user exists handler sends querry to database and take user's credentials
// After then handler hash entered password and compare it with taken password from database. If 2 hashed passwords equal, jwt token will be created and append to the cookies.
// If user signin without any error tahn user redirected to homepage.
func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit SignIn")
	var Result models.HtmlVar
	if r.Method == "POST" {
		var userFromdb models.User
		var userFromWeb models.User
		username := r.FormValue("username")
		password := r.FormValue("password")
		userFromWeb.Username = username
		userFromWeb.Password = password
		hashedpasswdFromWeb := HashSHA512(userFromWeb.Password)

		if !CheckUserExist(userFromWeb) {
			Result.Result = "User Does Not Exist"
			view.ExecuteTemplate(w, "signin.html", Result)
			return
		}
		row := db.QueryRow(fmt.Sprintf(shared.SqlGetUserByUserName+"'%s'", userFromWeb.Username))
		err := row.Scan(&userFromdb.ID, &userFromdb.Username, &userFromdb.Email, &userFromdb.Password)
		if err != nil {
			log.Fatal(err)
			http.Redirect(w, r, "/signin.html", http.StatusInternalServerError)
			return
		}
		if userFromdb.Username != userFromWeb.Username || userFromdb.Password != hashedpasswdFromWeb {
			Result.Result = "Wrong Name or Password"
			view.ExecuteTemplate(w, "signin.html", Result)
			return
		}
		claims := jwt.StandardClaims{
			Issuer:    userFromdb.ID,
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := token.SignedString(key)
		if err != nil {
			log.Fatal(err)
			return
		}
		jwtcookie := http.Cookie{
			Name:    "jwtToken",
			Value:   signedToken,
			Expires: time.Now().Add(time.Hour * 1),
		}
		http.SetCookie(w, &jwtcookie)
		http.Redirect(w, r, "/api/homepage", http.StatusSeeOther)

	}
	view.ExecuteTemplate(w, "signin.html", nil)
}
