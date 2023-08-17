package routes

import (
	"api-capital-tours/src/auth"
	"api-capital-tours/src/controller"
	"api-capital-tours/src/database/models/tables"
	"api-capital-tours/src/database/orm"
	"api-capital-tours/src/libraries/library"
	"api-capital-tours/src/middleware"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func RutasAuthAdmin(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/verify", verifyLogin).Methods("GET")
	s.Handle("/info-user", middleware.Autentication(http.HandlerFunc(getDatauser))).Methods("GET")
	s.HandleFunc("/login", login).Methods("PUT")
	s.Handle("/update-user", middleware.Autentication(http.HandlerFunc(updateUserAdmin))).Methods("PUT")
	s.Handle("/change-pass-user", middleware.Autentication(http.HandlerFunc(changePassUser))).Methods("POST")
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

	_data_user := orm.NewQuerys("users_admin").Select().Where("username", "=", req_body["username"]).Exec(orm.Config_Query{Cloud: true}).One()

	if len(_data_user) <= 0 {
		controller.ErrorsWaning(w, errors.New("usuario y/o contraseña incorrecto"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(_data_user["password"].(string)), []byte(req_body["password"].(string)))
	if err != nil {
		controller.ErrorsWaning(w, errors.New("usuario y/o contraseña incorrecto"))
		return
	}

	token, err := generateJWTToken(_data_user["id_user_admin"].(string), _data_user["email"].(string), _data_user["username"].(string), _data_user["nombre"].(string), _data_user["apellidos"].(string))
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

func generateJWTToken(idUser string, email string, username string, nombres string, apellidos string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error configuración de variables de entorno")
	}
	key := os.Getenv("ENV_KEY_JWT")

	claims := auth.JWTClaim{
		IdUser:    idUser,
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

func updateUserAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	id_user := library.GetTokenKey(r, "us")

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}
	data_request["where"] = map[string]interface{}{"id_user_admin": id_user}

	var data_update []map[string]interface{}
	data_update = append(data_update, data_request)

	schema, table := tables.UserAdmin_GetSchema()
	propietarios := orm.SqlExec{}
	err = propietarios.New(data_update, table).Update(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = propietarios.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = propietarios.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func changePassUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	if data_request["password_old"] == nil {
		controller.ErrorsWaning(w, errors.New("se requiere contraseña antigua"))
		return
	}
	if data_request["password"] == nil {
		controller.ErrorsWaning(w, errors.New("se requiere contraseña nueva"))
		return
	}

	id_user := library.GetTokenKey(r, "us")

	_data_user := orm.NewQuerys("users_admin").Select().Where("id_user_admin", "=", id_user).Exec(orm.Config_Query{Cloud: true}).One()

	err = bcrypt.CompareHashAndPassword([]byte(_data_user["password"].(string)), []byte(data_request["password_old"].(string)))
	if err != nil {
		controller.ErrorsWaning(w, errors.New("la contraseña anterior no es válida"))
		return
	}

	UpdatePassword := append([]map[string]interface{}{}, map[string]interface{}{
		"password": data_request["password"],
		"where": map[string]interface{}{
			"id_user_admin": id_user,
		},
	})

	if len(UpdatePassword[0]["password"].(string)) < 6 {
		controller.ErrorsError(w, errors.New("la contraseña debe ser mayor a 6 caracteres"))
		return
	}

	schema, table := tables.UserAdmin_GetSchema()
	users := orm.SqlExec{}
	errUpd := users.New(UpdatePassword, table).Update(schema)
	if errUpd != nil {
		controller.ErrorsWaning(w, errUpd)
		return
	}

	errUpd = users.Exec()
	if errUpd != nil {
		controller.ErrorsWaning(w, errUpd)
		return
	}
	response.Data = users.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getDatauser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	id_user := library.GetTokenKey(r, "us")

	_data_user := orm.NewQuerys("users_admin").Select().Where("id_user_admin", "=", id_user).Exec(orm.Config_Query{Cloud: true}).One()

	if len(_data_user) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	delete(_data_user, "id_user_admin")
	delete(_data_user, "password")
	delete(_data_user, "cargo")

	response.Data["user"] = _data_user

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
