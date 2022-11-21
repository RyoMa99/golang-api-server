package rest

import (
	"encoding/json"
	"net/http"

	"github.com/RyoMa99/go_ddd/usecase"
	"github.com/julienschmidt/httprouter"
)

type UserHandler interface {
	Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params)
	Create(w http.ResponseWriter, r *http.Request, pr httprouter.Params)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

type RequestUser struct {
	Name string `json:"name"`
}

// Userデータに関するHandlerを生成
func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

// UserIndex : GET /users -> 検索結果を返す
func (uh userHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	// GETパラメータ
	name := r.FormValue("name")

	user, err := uh.userUseCase.Search(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// クライアントにレスポンスを返却
	if err = json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (uh *userHandler) Create(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	var requestUser RequestUser
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := json.NewDecoder(r.Body).Decode(&requestUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	createdUser, err := uh.userUseCase.Create(requestUser.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(createdUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
