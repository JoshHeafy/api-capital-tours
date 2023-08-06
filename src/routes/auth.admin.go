package routes

import (
	"api-capital-tours/src/auth"
	"api-capital-tours/src/controller"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/deybin/go_basic_orm"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func RutasAuthAdmin(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/verify", verifyLogin).Methods("GET")
	s.HandleFunc("/login", login).Methods("PUT")
}

func verifyLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	token := r.Header.Get("Access-Token")
	_, err := auth.ValidateToken(token)

	if token == "" || err != nil {
		controller.ErrorUnauthorized(w, errors.New("inicio de sesión inválido"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	req_body, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	_data_user, _ := new(go_basic_orm.Querys).NewQuerys("users_admin").Select().Where("username", "=", req_body["username"]).Exec(go_basic_orm.Config_Query{Cloud: true}).One()

	if len(_data_user) <= 0 {
		controller.ErrorsWaning(w, errors.New("usuario y/o contraseña incorrecto"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(_data_user["password"].(string)), []byte(req_body["password"].(string)))
	if err != nil {
		controller.ErrorsWaning(w, errors.New("usuario y/o contraseña incorrecto"))
		return
	}

	token, err := generateJWTToken(_data_user["email"].(string), _data_user["username"].(string), _data_user["nombre"].(string), _data_user["apellidos"].(string))
	if err != nil {
		controller.ErrorServer(w, errors.New("error al generar token"))
		return
	}

	returnData := _data_user
	delete(returnData, "password")
	response.Data["user"] = returnData
	response.Data["token"] = token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func generateJWTToken(email string, username string, nombres string, apellidos string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error configuración de variables de entorno")
	}
	key := os.Getenv("ENV_KEY_JWT")

	claims := auth.JWTClaim{
		Email:     email,
		Username:  username,
		Nombres:   nombres,
		Apellidos: apellidos,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * (60 * 24)).Unix(),
			Issuer:    "pdt",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
