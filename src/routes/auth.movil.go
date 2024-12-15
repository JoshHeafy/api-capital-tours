package routes

import (
	"api-capital-tours/src/auth"
	"api-capital-tours/src/controller"
	"api-capital-tours/src/database/models/tables"
	"api-capital-tours/src/database/orm"
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

func RutasAuthMovil(r *mux.Router) {
	s := r.PathPrefix("/auth-movil").Subrouter()
	s.HandleFunc("/login", loginMovil).Methods("PUT")
	s.HandleFunc("/signup", signup).Methods("POST")
}

func loginMovil(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	req_body, err := controller.CheckBody(w, r)
	if err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("error al leer el cuerpo de la solicitud"))
		return
	}

	_data_user := orm.NewQuerys("users_mobile").Select().Where("email", "=", req_body["email"]).Exec(orm.Config_Query{Cloud: true}).One()

	if len(_data_user) <= 0 {
		controller.ErrorsWaning(w, errors.New("usuario y/o contraseña incorrecto"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(_data_user["password"].(string)), []byte(req_body["password"].(string)))
	if err != nil {
		controller.ErrorsWaning(w, errors.New("usuario y/o contraseña incorrecto"))
		return
	}

	token, err := generateJWTTokenMovil(_data_user["id_user"].(string), _data_user["email"].(string))
	if err != nil {
		controller.ErrorServer(w, errors.New("error al generar token"))
		return
	}

	_data_propietario := orm.NewQuerys("propietarios p").Select(
		"p.numero_documento, p.nombre_propietario, p.direccion, p.referencia, p.tipo_documento, p.telefono, p.email, i.numero_flota",
	).InnerJoin("vehiculos v", "p.numero_documento = v.numero_documento").InnerJoin("inscripciones i", "v.numero_placa = i.numero_placa").Where("v.numero_placa", "=", _data_user["numero_placa"]).Exec(orm.Config_Query{Cloud: true}).One()

	returnData := _data_user
	returnData["numero_flota"] = _data_propietario["numero_flota"]

	delete(returnData, "password")
	delete(_data_propietario, "numero_flota")
	response.Data["user"] = returnData
	response.Data["propietario"] = _data_propietario
	response.Data["token"] = token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func generateJWTTokenMovil(idUser string, email string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error configuración de variables de entorno")
	}
	key := os.Getenv("ENV_KEY_JWT")

	claims := auth.JWTClaim{
		IdUser: idUser,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * (60 * 24)).Unix(),
			Issuer:    "pdt",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		log.Println(err)
		return "", errors.New("error al generar token")
	}

	return tokenString, nil
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("error al leer el cuerpo de la solicitud"))
		return
	}

	_data_vehiculo := orm.NewQuerys("vehiculos").Select("numero_placa").Where("numero_placa", "=", data_request["numero_placa"]).Exec(orm.Config_Query{Cloud: true}).One()

	if _data_vehiculo["numero_placa"] == nil {
		controller.ErrorConflict(w, errors.New("esta placa no es válida"))
		return
	}

	_data_inscripciones := orm.NewQuerys("inscripciones").Select().Where("numero_placa", "=", _data_vehiculo["numero_placa"]).Exec(orm.Config_Query{Cloud: true}).One()

	if len(_data_inscripciones) <= 0 {
		controller.ErrorConflict(w, errors.New("esta placa no cuenta con una suscripción"))
		return
	}

	_data_user := orm.NewQuerys("users_mobile").Select("email").Where("email", "=", data_request["email"]).Exec(orm.Config_Query{Cloud: true}).One()

	if len(_data_user) >= 1 {
		controller.ErrorConflict(w, errors.New("email ya registrado"))
		return
	}

	var data_insert []map[string]interface{}
	data_insert = append(data_insert, data_request)

	schema, table := tables.UsersMovil_GetSchema()
	clientes := orm.SqlExec{}

	if err := clientes.New(data_insert, table).Insert(schema); err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	if err := clientes.Exec(); err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("hubo un error al registrar al usuario, por favor intente nuevamente o comuniquese con el administrador"))
		return
	}

	returnData := clientes.Data[0]
	delete(returnData, "id_user")
	delete(returnData, "password")
	response.Data = returnData

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
