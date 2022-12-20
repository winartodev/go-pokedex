package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/winartodev/go-pokedex/entity"
	"github.com/winartodev/go-pokedex/helper"
	"github.com/winartodev/go-pokedex/usecase"
)

type Server struct {
	Router         *httprouter.Router
	PokemonUsecase usecase.PokemonUsecaseItf
	TypeUsecase    usecase.TypeUsecaseItf
	UserUsecase    usecase.UserUsecaseItf
}

func (s *Server) GetAllPokemon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var pokemons []entity.PokemonList
	var err error
	var ctx = r.Context()

	filter := buildQueryFilter(r.URL.Query())
	if len(filter) > 0 {
		pokemons, err = s.PokemonUsecase.GetAllPokemonByFilter(ctx, filter)
		if err != nil {
			helper.FailedResponse(w, http.StatusBadRequest, err)
			return
		}
	} else {
		pokemons, err = s.PokemonUsecase.GetAllPokemon(ctx)
		if err != nil {
			helper.FailedResponse(w, http.StatusBadRequest, err)
			return
		}
	}

	helper.SuccessResponse(w, "", pokemons)
}

func (s *Server) GetPokemonByID(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id, err := strconv.ParseInt(param.ByName("id"), 10, 64)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	pokemon, err := s.PokemonUsecase.GetPokemonByID(r.Context(), id)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(w, "", pokemon)
}

func (s *Server) CatchPokemon(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id, err := strconv.ParseInt(param.ByName("id"), 10, 64)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	err = s.PokemonUsecase.CatchPokemon(r.Context(), id)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(w, "Pokemon success catched", nil)
}

func (s *Server) CreatePokemon(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	var pokemon entity.Pokemon
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pokemon); err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	id, err := s.PokemonUsecase.CreatePokemon(r.Context(), pokemon)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(w, "success create pokemon", id)
}

func (s *Server) UpdatePokemon(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id, err := strconv.ParseInt(param.ByName("id"), 10, 64)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	var pokemon entity.Pokemon
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pokemon); err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	res, err := s.PokemonUsecase.UpdatePokemon(r.Context(), id, pokemon)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(w, "update pokemon success ", res)
}

func (s *Server) DeletePokemon(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id, err := strconv.ParseInt(param.ByName("id"), 10, 64)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	err = s.PokemonUsecase.DeletePokemon(r.Context(), id)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(w, "delete pokemon success", nil)
}

func (s *Server) GetAllType(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := s.TypeUsecase.GetAllType(r.Context())
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(w, "", res)
}

func (s *Server) CreateType(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var types entity.Type
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&types); err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	res, err := s.TypeUsecase.CreateType(r.Context(), types)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(w, "create type success", res)
}

func (s *Server) GetTypeByID(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id, err := strconv.ParseInt(param.ByName("id"), 10, 64)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	res, err := s.TypeUsecase.GeTypeByID(r.Context(), id)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(w, "", res)
}

func (s *Server) UpdateType(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id, err := strconv.ParseInt(param.ByName("id"), 10, 64)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	var types entity.Type
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&types); err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	err = s.TypeUsecase.UpdateType(r.Context(), id, types)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(w, "update type success", nil)
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request entity.User
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	// validate username & password
	if request.Username == "" {
		helper.FailedResponse(w, http.StatusBadRequest, errors.New("username can't be empty"))
		return
	}
	if request.Password == "" {
		helper.FailedResponse(w, http.StatusBadRequest, errors.New("password can't be empty"))
		return
	}

	id, err := s.UserUsecase.Register(r.Context(), request.Username, request.Email, request.Password, request.Role)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(w, "", id)
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request entity.User
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	// validate username & password
	if request.Username == "" {
		helper.FailedResponse(w, http.StatusBadRequest, errors.New("username can't be empty"))
		return
	}
	if request.Password == "" {
		helper.FailedResponse(w, http.StatusBadRequest, errors.New("password can't be empty"))
		return
	}

	token, err := s.UserUsecase.Login(r.Context(), request.Username, request.Password)
	if err != nil {
		helper.FailedResponse(w, http.StatusBadRequest, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
	})

	helper.SuccessResponse(w, "login success", nil)
}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		MaxAge: -1,
	})

	helper.SuccessResponse(w, "user logout success", nil)
}

func (s *Server) Healthz(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("ok"))
}
