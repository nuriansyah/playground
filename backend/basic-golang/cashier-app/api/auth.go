package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type LoginSuccessResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type AuthErrorResponse struct {
	Error string `json:"error"`
}

// Jwt key yang akan dipakai untuk membuat signature
var jwtKey = []byte("key")

// Struct claim digunakan sebagai object yang akan di encode oleh jwt
// jwt.StandardClaims ditambahkan sebagai embedded type untuk provide standart claim yang biasanya ada pada JWT
type Claims struct {
	Username string
	jwt.StandardClaims
}

func (api *API) login(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	username := req.URL.Query().Get("username")
	password := req.URL.Query().Get("password")
	res, err := api.usersRepo.Login(username, password)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	// Task: 1. Deklarasi expiry time untuk token jwt
	//       2. Buat claim menggunakan variable yang sudah didefinisikan diatas
	//       3. expiry time menggunakan time millisecond

	// TODO: answer here

	expiryTime := time.Now().Add(time.Millisecond * 100000).Unix()
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "golang-jwt",
			ExpiresAt: expiryTime,
		},
		Username: username,
	}

	// Task: Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai

	// TODO: answer here
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Task: 1. Buat jwt string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan
	//       2. return internal error ketika ada kesalahan ketika pembuatan JWT string

	// TODO: answer here
	jwtString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	// Task: Set token string kedalam cookie response

	// TODO: answer here
	cookie := http.Cookie{
		Name:    "token",
		Value:   jwtString,
		Expires: time.Unix(expiryTime, 0),
	}
	http.SetCookie(w, &cookie)

	// Task: Return response berupa username dan token JWT yang sudah login
	encoder.Encode(LoginSuccessResponse{
		Username: *res,
	})
	// json.NewEncoder(w).Encode(LoginSuccessResponse{Username: "", Token: ""}) // TODO: replace this
}

func (api *API) logout(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	username := req.URL.Query().Get("username")
	err := api.usersRepo.Logout(username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder := json.NewEncoder(w)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	cookie := http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
	}
	http.SetCookie(w, &cookie)

}
