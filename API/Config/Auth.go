package Config

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const SECRET = "qwertyuiopasdfghjklzxcvbnm123456"


func GenrateJWT(email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"verified": true,
		"user" : email,
	})
	hmacSampleSecret := []byte(SECRET)
	jwtstring,_ := token.SignedString(hmacSampleSecret)
	return jwtstring
}


func CheckJWT(mytoken string) (bool){
	hmacSampleSecret := []byte(SECRET)
	jtoken,_ := jwt.Parse(mytoken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if _, ok := jtoken.Claims.(jwt.MapClaims); ok && jtoken.Valid {
		return true
	} else {
		return false
	}
}

func ClaimJWT(mytoken string) string{
	hmacSampleSecret := []byte(SECRET)
	jtoken,_ := jwt.Parse(mytoken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claim, ok := jtoken.Claims.(jwt.MapClaims); ok && jtoken.Valid {
		user := claim["user"]
		return user.(string)
	} else {
		return ""
	}
}