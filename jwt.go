package util

import (
	//The go-jwt package is used for signing the encoded security token.
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
)

type JWTServer struct {
	method  jwt.SigningMethod
	signKey *rsa.PrivateKey
}

func NewJWTServer(method jwt.SigningMethod, signKey *rsa.PrivateKey) *JWTServer {
	return &JWTServer{method: method, signKey: signKey}
}

//Generate JWT token
func (s *JWTServer) GenerateWithClaims(claims jwt.Claims) (string, error) {
	return jwt.NewWithClaims(s.method, claims).SignedString(s.signKey)
}

type JWTClient struct {
	verifyKey *rsa.PublicKey
}

func NewJWTClient(verifyKey *rsa.PublicKey) *JWTClient {
	return &JWTClient{verifyKey: verifyKey}
}

func (c *JWTClient) ParseWithClaims(tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	//validata the token
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		//verify the token with public key , which is the counter part of private key
		return c.verifyKey, nil
	})
}
