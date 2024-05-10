package utils

 import (
     "backend/model"

     "github.com/dgrijalva/jwt-go"
 )

 func ParseToken(tokenString string) (claims *model.Claims, err error) {
     token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
         return []byte("my_secret_key"), nil
     })

     if err != nil {
         return nil, err
     }

     claims, ok := token.Claims.(*model.Claims)

     if !ok {
         return nil, err
     }

     return claims, nil
 }