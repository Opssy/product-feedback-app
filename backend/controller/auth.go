package controller

import (
	"backend/model"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my secret key")


func Login (w http.ResponseWriter, r *http.Request){
 if r.Method != http.MethodPost {
    w.WriteHeader(http.StatusMethodNotAllowed)
    fmt.Fprintf(w, "Method not allowed: %s", r.Method)
    return
  }
  var user model.User

  //Decode request body into user object
  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&user);
  err !=nil{
	 w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "Error decoding request body: %v", err)
    return
  }

  //ehecking user email 

  var existingUser model.User
  if err := model.DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Error finding user: %v", err)
    return
  }

  if existingUser.ID == 0 {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "User does not exist")
    return
  }

  //validate password
  if !utils.CompareHashPassword(user.Password, existingUser.Password){
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "Invalid password")
    return
  }
  // Generate JWT token
  expirationTime := time.Now().Add(5 * time.Minute)
  claims := model.Claims{
    Role: existingUser.Role,
    StandardClaims: jwt.StandardClaims{
      Subject:   existingUser.Email,
      ExpiresAt: expirationTime.Unix(),
    },
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, err := token.SignedString(jwtKey)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Error generating token: %v", err)
    return
  }

  // Set cookie with token
  http.SetCookie(w, &http.Cookie{
    Name:  "token",
    Value: tokenString,
    Path:  "/",
  })

  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "User logged in")

}

func Signup(w http.ResponseWriter, r *http.Request){

      var user model.User

   if r.Method != http.MethodPost {
    w.WriteHeader(http.StatusMethodNotAllowed)
    fmt.Fprintf(w, "Method not allowed: %s", r.Method)
    return
   }

 decoder := json.NewDecoder(r.Body)
 if err := decoder.Decode(&user)
 err != nil{
  w.WriteHeader(http.StatusBadRequest)
   fmt.Fprintf(w, "Error decoding request body: %v", err)
    return
 }  

 hashedPassword, err := utils.GenerateHashPassword(user.Password)
   if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Error hashing password: %v", err)
    return
  }
  user.Password = hashedPassword
 // Save user to database
 if err := model.DB.Create(&user).Error; 
  err != nil{
     w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Error saving user: %v", err)
    return
  }

    if err != nil {
    // Handle signup error
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Error saving user: %v", err)
    return
  }

  w.WriteHeader(http.StatusCreated)
  fmt.Fprintf(w, "User created successfully")
}