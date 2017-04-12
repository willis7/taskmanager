package common

import (
	"io/ioutil"
	"log"
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
)

//using asymmetric crypto/RSA keys
const (
	// openssl genrsa -out app.rsa 1024
	privKeyPath = "keys/app.rsa"
	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "keys/app.rsa.pub"
)

// private key for signing and public key for verification
var (
	verifyKey, signKey []byte
)

// Read the key files before starting http handlers
func initKeys() {
	var err error

	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
		panic(err)
	}
}

// Generate JWT token
func GenerateJWT(name, role string) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RSA256"))

	// set claims for JWT token
	t.Claims["iss"] = "admin"
	t.Claims["UserInfo"] = struct {
		Name string
		Role string
	}{name, role}

	// set the expire time for the JWT token
	t.Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Middleware for validating JWT tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// validate the token
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		// Verify the tokekn with the public key, which is the counter part of the private key
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {

		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired: // JWT expired
				DisplayAppError(
					w,
					err,
					"Access Token has expired, get a new Token",
					401,
				)
				return

			default:
				DisplayAppError(
					w,
					err,
					"Error while parsing the Token!",
					500,
				)
				return
			}
		default:
			DisplayAppError(
				w,
				err,
				"Error while parsing the Token!",
				500,
			)
			return
		}
	}
	if token.Valid {
		next(w,r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		DisplayAppError(
			w,
			err,
			"Invalid Access Token",
			401,
		)
	}
}
