package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateToken(ttl time.Duration, payload interface{}) (string, error) {
	privateRsaPath, ok := os.LookupEnv("PRIVATE_RSA_PATH")
	if !ok {
		panic(fmt.Sprintf("No private rsa path specified for %s", privateRsaPath))
	}

	privateRsa, err := os.ReadFile(privateRsaPath)
	if err != nil {
		fmt.Println("base64.StdEncoding.DecodeString - err", err)
		// TODO: Improve error
		return "", nil
	}
	// TODO: move to ENV file
	// decodedPrivateKey, err := base64.StdEncoding.DecodeString("LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlCUEFJQkFBSkJBTzVIKytVM0xrWC91SlRvRHhWN01CUURXSTdGU0l0VXNjbGFFKzlaUUg5Q2VpOGIxcUVmCnJxR0hSVDVWUis4c3UxVWtCUVpZTER3MnN3RTVWbjg5c0ZVQ0F3RUFBUUpCQUw4ZjRBMUlDSWEvQ2ZmdWR3TGMKNzRCdCtwOXg0TEZaZXMwdHdtV3Vha3hub3NaV0w4eVpSTUJpRmI4a25VL0hwb3piTnNxMmN1ZU9wKzVWdGRXNApiTlVDSVFENm9JdWxqcHdrZTFGY1VPaldnaXRQSjNnbFBma3NHVFBhdFYwYnJJVVI5d0loQVBOanJ1enB4ckhsCkUxRmJxeGtUNFZ5bWhCOU1HazU0Wk1jWnVjSmZOcjBUQWlFQWhML3UxOVZPdlVBWVd6Wjc3Y3JxMTdWSFBTcXoKUlhsZjd2TnJpdEg1ZGdjQ0lRRHR5QmFPdUxuNDlIOFIvZ2ZEZ1V1cjg3YWl5UHZ1YStxeEpXMzQrb0tFNXdJZwpQbG1KYXZsbW9jUG4rTkVRdGhLcTZuZFVYRGpXTTlTbktQQTVlUDZSUEs0PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==")
	// if err != nil {
	// 	fmt.Println("base64.StdEncoding.DecodeString - err", err)
	// 	// TODO: Improve error
	// 	return "", nil
	// }

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateRsa)
	if err != nil {
		fmt.Println("ParseRSAPrivateKeyFromPEM - err", err)
		// TODO: Improve error
		return "", nil
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(15 * time.Hour).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		fmt.Println("jwt.NewWithClaims - err", err)
		// TODO: Improve error
		return "", nil
	}

	fmt.Println("jwt.NewWithClaims - err", err)
	// TODO: Improve error
	return token, nil
}

func ValidateToken(token string, publicKey string) (interface{}, error) {
	if token == "" {
		fmt.Println("Invalid token received")
		return nil, status.Error(codes.Unauthenticated, "Invalid token received")
	}

	publicRsaPath, ok := os.LookupEnv("PUBLIC_RSA_PATH")
	if !ok {
		panic(fmt.Sprintf("No public rsa path specified for %s", publicRsaPath))
	}

	publicRsa, err := os.ReadFile(publicRsaPath)
	if err != nil {
		fmt.Println(err)
		// TODO: Improve error
		return "", err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(publicRsa)
	if err != nil {
		fmt.Println(err)
		// TODO: Improve error
		return "", err
	}

	// bearer := "Bearer "
	// token = token[len(bearer):]

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			// TODO: Improve error
			return nil, fmt.Errorf("error on jwt parse method")
		}
		return key, nil
	})
	// Check if has some error on token validation
	if err != nil {
		fmt.Println("Error: ", err)
		// TODO: Improve error
		return "", err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		// TODO: Improve error
		return "", fmt.Errorf("error on parsed claims tokens")
	}

	return claims["sub"], nil
}
