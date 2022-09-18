package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/swildz/go-jwt-siddiq/config"
	"github.com/swildz/go-jwt-siddiq/helper"
	"github.com/swildz/go-jwt-siddiq/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {

	//mengambil inputan json
	var userInput models.User

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"messsage": err.Error()}
		helper.ResponJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	//ambil data user berdasarkan username

	var user models.User

	if err := models.DB.Where("user_name = ?", userInput.UserName).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"messsage": "Username atau password salah"}
			helper.ResponJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"messsage": err.Error()}
			helper.ResponJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	//cek apakah password valid

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"messsage": "Username atau password salah"}
		helper.ResponJSON(w, http.StatusUnauthorized, response)
		return
	}

	//proses pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-siddiq",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	//Mendeklarasikan algoritma yang akan digunakan untuk signin
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//sign Token

	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"messsage": "Error"}
		helper.ResponJSON(w, http.StatusInternalServerError, response)
		return
	}

	// set token yang ke cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"messsage": "login succes"}
	helper.ResponJSON(w, http.StatusOK, response)

}

func Register(w http.ResponseWriter, r *http.Request) {

	//mengambil inputan json
	var userInput models.User

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"messsage": err.Error()}
		helper.ResponJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	//hash password menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	//insert ke databases
	if err := models.DB.Create(&userInput).Error; err != nil {

		response := map[string]string{"messsage": err.Error()}
		helper.ResponJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]string{"messsage": "success"}
	helper.ResponJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {

	// hapus token yang ada cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"messsage": "logout berhasil"}
	helper.ResponJSON(w, http.StatusOK, response)

}
