package main

import (
	"encoding/json"
	"fmt"
	"game-app/entity"
	"game-app/repository/mysql"
	"game-app/service/user"
	"io"
	"log"
	"net/http"
)

const (
	jwtSignKey = "jwt_secret"
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

	mysqlRepo := mysql.New()
	userSvc := user.New(mysqlRepo, jwtSignKey)

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

	mysqlRepo := mysql.New()
	userSvc := user.New(mysqlRepo, jwtSignKey)

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
	var data []byte
	var rData []byte
	if req.Method != http.MethodGet {
		fmt.Fprintf(writer, `{"error": "invalid method"}`)
	}
	var pReq user.ProfileRequest
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error"": "%s"}`, err.Error())))
		return
	}
	err = json.Unmarshal(data, &pReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error"": "%s"}`, err.Error())))
		return
	}

	mysqlRepo := mysql.New()
	userSvc := user.New(mysqlRepo, jwtSignKey)

	resp, err := userSvc.Profile(pReq)
	if err != nil {
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
