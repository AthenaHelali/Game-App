package main

import (
	"encoding/json"
	"fmt"
	"game-app/entity"
	"game-app/repository/mysql"
	"game-app/service/authservice"
	"game-app/service/user"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	jwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users/register", userRegisterHandler)
	mux.HandleFunc("/users/login", userLoginHandler)
	mux.HandleFunc("/users/profile", userProfileHandler)
	mux.HandleFunc("/health-check", healthCheckHandler)
	server := http.Server{Addr: ":8080", Handler: mux}
	log.Fatalln(server.ListenAndServe())

}
func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error": "invalid method"}`)
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error"": "%s"}`, err.Error())))
		return
	}

	var uReq user.RegisterRequest
	err = json.Unmarshal(data, &uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error"": "%s"}`, err.Error())))
		return
	}
	authSvc := authservice.New(jwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)
	mysqlRepo := mysql.New()
	userSvc := user.New(authSvc, mysqlRepo)

	_, rErr := userSvc.Register(uReq)
	if rErr != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, rErr.Error())))
		return
	}
	writer.Write([]byte(`{"message:" : "user created"}`))
}

func healthCheckHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, `{"message": "everything is good!"}`)
}
func TestUserMysqlRepo() {
	mysqlRepo := mysql.New()

	createdUser, err := mysqlRepo.RegisterUser(entity.User{
		ID:          0,
		PhoneNumber: "091988855",
		Name:        "Athena Helali",
	})

	if err != nil {
		fmt.Println("register user", err)
	} else {
		fmt.Println("created user:", createdUser)
	}

	isUnique, err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber + "2")
	if err != nil {
		fmt.Println("unique error", err)
	}

	fmt.Println("is unique", isUnique)

}

func userLoginHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error": "invalid method"}`)
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error"": "%s"}`, err.Error())))
		return
	}

	var lReq user.LoginRequest
	err = json.Unmarshal(data, &lReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error"": "%s"}`, err.Error())))
		return
	}

	authSvc := authservice.New(jwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)
	mysqlRepo := mysql.New()
	userSvc := user.New(authSvc, mysqlRepo)

	resp, rErr := userSvc.Login(lReq)
	if rErr != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, rErr.Error())))

		return
	}
	var rData []byte
	rData, err = json.Marshal(resp)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error"": "%s"}`, err.Error())))

		return
	}

	writer.Write(rData)
}

func userProfileHandler(writer http.ResponseWriter, req *http.Request) {
	var rData []byte
	if req.Method != http.MethodGet {
		fmt.Fprintf(writer, `{"error": "invalid method"}`)
	}
	authSvc := authservice.New(jwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)
	Btoken := req.Header.Get("Authorization")
	claims, err := authSvc.ParseToken(Btoken)
	if err != nil {
		fmt.Fprintf(writer, `{"error": "token is not valid"}`)
	}

	mysqlRepo := mysql.New()
	userSvc := user.New(authSvc, mysqlRepo)
	resp, sErr := userSvc.Profile(user.ProfileRequest{UserID: claims.UserID})
	if sErr != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error"": "%s"}`, err.Error())))

		return
	}

	rData, err = json.Marshal(resp)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error"": "%s"}`, err.Error())))
		return
	}

	writer.Write(rData)

}
