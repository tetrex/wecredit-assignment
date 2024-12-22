package jwt

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AccessTokens struct {
	AccessCode string `json:"access_code"`
}

type JwtCustomClaims struct {
	CreatedAt int `json:"created_at"`
	ValidFor  int `json:"valid_for"`
	UserId    int `json:"user_id"`
	jwt.RegisteredClaims
}

var (
	jwtSec = "parent_concern_secret"
)

func verifyToken(token string, jwtSec string) (*JwtCustomClaims, error) {
	secretKey := jwtSec

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("token is invalid")
		}
		return []byte(secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}

	payload, ok := jwtToken.Claims.(*JwtCustomClaims)
	if !ok {
		return nil, errors.New("token is invalid")
	}

	return payload, nil
}

func GenerateTokens(userId int) (*AccessTokens, error) {
	accessToken, _, err := generateAccessToken(userId, jwtSec)
	if err != nil {
		return nil, err
	}

	cacheJSON := AccessTokens{
		AccessCode: accessToken,
	}

	return &cacheJSON, nil
}

func getTokenFromBearer(c echo.Context) string {
	var vals = c.Request().Header
	var m = make(map[string]string)
	for key := range vals {
		m[key] = vals.Get(key)
	}

	reqToken := m["Authorization"]

	if reqToken == "" {
		return ""
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}
	return splitToken[1]
}

func generateToken(userId int, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	// Create the JWT claims, which includes the userId and expiry time
	claims := &JwtCustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

func generateAccessToken(userId int, jwtSec string) (string, time.Time, error) {
	// Declare the expiration time of the token
	expirationTime := time.Now().Add(1 * time.Hour)

	return generateToken(userId, expirationTime, []byte(jwtSec))
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := getTokenFromBearer(c)
		if token == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "token not found"})
		}
		claims, err := verifyToken(token, jwtSec)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": err.Error()})
		}
		user := JwtCustomClaims{
			CreatedAt: claims.CreatedAt,
			ValidFor:  claims.ValidFor,
			UserId:    claims.UserId,
		}
		c.Set("user", user)

		return next(c)
	}
}

func GetJwt(c echo.Context) (JwtCustomClaims, bool) {
	user, ok := c.Get("user").(JwtCustomClaims)
	if !ok {
		return JwtCustomClaims{}, false
	}

	res := JwtCustomClaims{
		CreatedAt:        user.CreatedAt,
		ValidFor:         user.ValidFor,
		UserId:           user.UserId,
		RegisteredClaims: user.RegisteredClaims,
	}

	return res, true
}

func OptionalJWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := getTokenFromBearer(c)
		if token != "" {
			claims, err := verifyToken(token, jwtSec)
			if err == nil {
				user := JwtCustomClaims{
					CreatedAt: claims.CreatedAt,
					ValidFor:  claims.ValidFor,
					UserId:    claims.UserId, //  to string
				}
				c.Set("user", user)
			}
		}
		return next(c)
	}
}
