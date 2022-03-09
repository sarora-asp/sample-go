package jwt

import (
	"errors"
	"fmt"
	user "sample/twirp/model/user"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func New(secret string, duration time.Duration, algo string) (*JWT, error) {
	if secret == "" {
		return nil, errors.New("Secret is required")
	}

	if duration == 0 {
		return nil, errors.New("Duration is required")
	}

	if algo == "" {
		return nil, errors.New("Algo is required")
	}

	signingMethod := jwt.GetSigningMethod(algo)

	return &JWT{
		key:      []byte(secret),
		algo:     signingMethod,
		duration: duration,
	}, nil
}

type JWT struct {
	// Secret key used for signing.
	key []byte

	// Duration for which the jwt token is valid.
	duration time.Duration

	// JWT signing algorithm
	algo jwt.SigningMethod
}

type ClaimObject struct {
	Id   int
	Name string
	jwt.StandardClaims
}

func (j *JWT) GenerateToken(u *user.AuthUser) (string, error) {
	claims := ClaimObject{
		Id:   u.Id,
		Name: u.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.duration).Unix(),
			Issuer:    "owner",
		},
	}
	token := jwt.NewWithClaims(j.algo, claims)

	return token.SignedString(j.key)
}

func (j *JWT) ParseToken(token string) (*user.AuthUser, error) {
	claims, error := j.verifyToken(token)
	fmt.Println(claims)
	if error != nil {
		fmt.Println(error)
	}

	Id, ok := claims["Id"]
	if !ok {
		fmt.Println("Unable to parse")
	}

	Name, ok := claims["Name"]
	if !ok {
		fmt.Println("Unable to parse")
	}

	return &user.AuthUser{
		Id:   Id.(int),
		Name: Name.(string),
	}, nil
}

func (j *JWT) verifyToken(token string) (map[string]interface{}, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if j.algo != token.Method {
			return nil, fmt.Errorf("Unable to verify gtoken\n")
		}

		return j.key, nil
	})

	if err != nil {
		fmt.Println(err)
		fmt.Println("Unable to verify token")
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Unable to verify the token")
}

// func (j *JWT) ParseToken(token string) (*twisk.AuthUser, error) {
// 	claims, err := j.verifyToken(token)
// 	if err != nil {
// 		return nil, err
// 	}

// 	id, ok := claims["id"]
// 	if !ok {
// 		return nil, fmt.Errorf("unauthorized: no id claim present")
// 	}

// 	tenantID, ok := claims["t"]
// 	if !ok {
// 		return nil, fmt.Errorf("unauthorized: no tenant_id claim present")
// 	}

// 	username, ok := claims["u"]
// 	if !ok {
// 		return nil, fmt.Errorf("unauthorized: no username claim present")
// 	}

// 	email, ok := claims["e"]
// 	if !ok {
// 		return nil, fmt.Errorf("unauthorized: no email claim present")
// 	}

// 	role, ok := claims["r"]
// 	if !ok {
// 		return nil, fmt.Errorf("unauthorized: no role claim present")
// 	}

// 	return &twisk.AuthUser{
// 		ID:       int64(id.(float64)),
// 		TenantID: int32(tenantID.(float64)),
// 		Username: username.(string),
// 		Email:    email.(string),
// 		Role:     twisk.AccessRole(role.(float64)),
// 	}, nil

// }

// func (j *JWT) verifyToken(token string) (map[string]interface{}, error) {
// 	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 		if token.Method != j.algo {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return j.key, nil
// 	})

// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, fmt.Errorf("could not parse provided token")
// 	}

// 	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
// 		return claims, nil
// 	}

// 	return nil, fmt.Errorf("jwt token could not be verified")
// }
