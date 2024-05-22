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

//signup 

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

//Home 

func Home(w http.ResponseWriter, r *http.Request) {
  // Check for presence of "token" cookie
  cookie, err := r.Cookie("token")
  if err != nil {
    if err == http.ErrNoCookie {
      // Handle missing cookie case (e.g., redirect to login)
      fmt.Fprintf(w, "Missing authentication token")
      return
    }
    // Handle other cookie errors
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Error retrieving cookie: %v", err)
    return
  }

  // Parse the token from the cookie
  claims, err := utils.ParseToken(cookie.Value)
  if err != nil {
    w.WriteHeader(http.StatusUnauthorized)
    fmt.Fprintf(w, "Invalid token: %v", err)
    return
  }

  // Validate user role
  if !(claims.Role == "user" || claims.Role == "admin") {
    w.WriteHeader(http.StatusUnauthorized)
    fmt.Fprintf(w, "Unauthorized access: Invalid role")
    return
  }

  // Access granted, send response
  w.WriteHeader(http.StatusOK)
  response := map[string]string{
    "success": "home page",
    "role":    claims.Role,
  }
  json.NewEncoder(w).Encode(response)
}

// func Edit(w http.ResponseWriter, r *http.Request) {
   
//   cookie, err := r.Cookie("token")
//   if err != nil{
//     if err == http.ErrNoCookie{
//         fmt.Fprintf(w, "Missing authentication token")
//       return
//     }
//   }
//    w.WriteHeader(http.StatusInternalServerError)
//     fmt.Fprintf(w, "Error retrieving cookie: %v", err)
//     return

//       // Parse the token from the cookie
//   claims, err := utils.ParseToken(cookie.Value)
//   if err != nil {
//     w.WriteHeader(http.StatusUnauthorized)
//     fmt.Fprintf(w, "Invalid token: %v", err)
//     return
//   }

//   // Validate user role (modify based on your authorization needs)
//   if claims.Role != "user" && claims.Role != "admin" {
//     w.WriteHeader(http.StatusUnauthorized)
//     fmt.Fprintf(w, "Unauthorized access: Invalid role")
//     return
//   }

//    // If method is not POST, return error
//   if r.Method != http.MethodPost {
//     w.WriteHeader(http.StatusMethodNotAllowed)
//     fmt.Fprintf(w, "Method not allowed: %s", r.Method)
//     return
//   }
//    userID := claims.ID // Replace with actual logic to get user ID from claims

//   // Decode request body into user update data
//   var userUpdate model.UpdatedFeedBack // Assuming you have a model.UserUpdate struct for updates
//   decoder := json.NewDecoder(r.Body)
//   if err := decoder.Decode(&userUpdate); err != nil {
//     w.WriteHeader(http.StatusBadRequest)
//     fmt.Fprintf(w, "Error decoding request body: %v", err)
//     return
//   }
//   // Find existing user by ID
//   var existingUser model.User
//   if err := model.DB.First(&existingUser, userID).Error; err != nil {
//     w.WriteHeader(http.StatusInternalServerError)
//     fmt.Fprintf(w, "Error finding user: %v", err)
//     return
//   }

//   // Update user fields (based on userUpdate data)
//   if userUpdate.Email != "" {
//     existingUser.Email = userUpdate.Email
//   }

//   // Update password if provided (use utils.UpdateHashPassword for secure update)
//   if userUpdate.Password != "" {
//     hashedPassword, err := utils.CompareHashPassword(existingUser.Password, userUpdate.Password)
//     if err != nil {
//       w.WriteHeader(http.StatusInternalServerError)
//       fmt.Fprintf(w, "Error updating password: %v", err)
//       return
//     }
//     existingUser.Password = hashedPassword
//   }

//   // Save updated user to database
//   if err := model.DB.Save(&existingUser).Error; err != nil {
//     w.WriteHeader(http.StatusInternalServerError)
//     fmt.Fprintf(w, "Error saving user: %v", err)
//     return
//   }
//    // User edited successfully
//   w.WriteHeader(http.StatusOK)
//   fmt.Fprintf(w, "User edited successfully")

// }