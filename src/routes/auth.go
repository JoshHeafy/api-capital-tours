package routes

import (
	"api-capital-tours/src/auth"
	"api-capital-tours/src/controller"
	"api-capital-tours/src/database/models/tables"
	"encoding/json"
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

func RutasAuth(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/", authentication).Methods("GET")
	s.HandleFunc("/verify", verifyLogin).Methods("GET")
	s.HandleFunc("/login", login).Methods("PUT")
	s.HandleFunc("/first-step", register).Methods("POST")
}

func authentication(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func verifyLogin(w http.ResponseWriter, r *http.Request) {
	response := controller.NewResponseManager()
	token := r.Header.Get("Access-Token")
	_, err := auth.ValidateToken(token)

	if token == "" || err != nil {
		response.Msg = "Inicio de Sesión Inválido"
		response.Status = "error"
		response.StatusCode = 401
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
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

	dataUser, _ := new(go_basic_orm.Querys).NewQuerys("users").Select().Where("email", "=", req_body["email"]).Exec(go_basic_orm.Config_Query{Cloud: true}).One()

	if len(dataUser) <= 0 {
		response.Msg = "Usuario y/o Contraseña Incorrecto"
		response.StatusCode = 300
		response.Status = "Usuario y/o Contraseña Incorrecto"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser["password"].(string)), []byte(req_body["password"].(string)))
	if err != nil {
		response.Msg = "Usuario y/o Contraseña Incorrecto"
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	token, err := generateJWTToken(dataUser["email"].(string))
	if err != nil {
		response.Msg = "Error al generar el token"
		response.StatusCode = http.StatusInternalServerError
		response.Status = "Error"
		json.NewEncoder(w).Encode(response)
		return
	}

	returnData := dataUser
	delete(returnData, "password")
	response.Data["user"] = returnData
	response.Data["token"] = token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	dataUser, err := new(go_basic_orm.Querys).NewQuerys("users").Select("email").Where("email", "=", data_request["email"]).Exec(go_basic_orm.Config_Query{Cloud: true}).One()

	if dataUser != nil || err != nil {
		response.Msg = "Usuario ya registrado"
		response.Status = "error"
		response.StatusCode = 409
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(response)
		return
	}

	var data_insert []map[string]interface{}
	data_insert = append(data_insert, data_request)

	schema, table := tables.Users_GetSchema()
	clientes := go_basic_orm.SqlExec{}
	err = clientes.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = clientes.Exec("Platcont")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	returnData := clientes.Data[0]
	delete(returnData, "id_user")
	delete(returnData, "password")
	delete(returnData, "password")
	response.Data = returnData

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func generateJWTToken(email string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error configuración de variables de entorno")
	}
	key := os.Getenv("ENV_KEY_JWT")

	claims := jwtclaim{
		Email: email,
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

type jwtclaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
