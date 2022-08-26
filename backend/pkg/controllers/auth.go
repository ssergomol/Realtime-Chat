package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ssergomol/RealtimeChat/pkg/database"
	"github.com/ssergomol/RealtimeChat/pkg/models"
	"github.com/ssergomol/RealtimeChat/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func SignUp(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	utils.ParseBody(r, &data)

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Username: data["username"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	res, _ := json.Marshal(user)
	w.Write(res)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	utils.ParseBody(r, &data)

	var user models.User

	database.DB.Where("username = ?", data["username"]).First(&user)

	if user.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		utils.SetBodyInfoMessage(w, "User not found")
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		utils.SetBodyInfoMessage(w, "Incorrect password")
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.SetBodyInfoMessage(w, "Couldn't sign in")
		return
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	utils.SetBodyInfoMessage(w, "Successfully signed in")
	return
}

func getUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		fmt.Printf("No jwt token cookie was found\n")
		utils.SetBodyInfoMessage(w, "Couldn't get jwt token")
		return
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.SetBodyInfoMessage(w, "Unauthenticated")
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	res, _ := json.Marshal(user)
	w.Write(res)
}

func SignOut(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	utils.SetBodyInfoMessage(w, "Successfully signed out")
}
