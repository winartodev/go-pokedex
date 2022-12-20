package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/mock"
	"github.com/winartodev/go-pokedex/entity"
	"github.com/winartodev/go-pokedex/usecase"
	usecasemock "github.com/winartodev/go-pokedex/usecase/mocks"
)

type mockServerProvider struct {
	Router         *httprouter.Router
	PokemonUsecase *usecasemock.PokemonUsecaseItf
	TypeUsecase    *usecasemock.TypeUsecaseItf
	UserUsecase    *usecasemock.UserUsecaseItf
}

func serverPorvider() mockServerProvider {
	return mockServerProvider{
		Router:         httprouter.New(),
		PokemonUsecase: new(usecasemock.PokemonUsecaseItf),
		TypeUsecase:    new(usecasemock.TypeUsecaseItf),
		UserUsecase:    new(usecasemock.UserUsecaseItf),
	}
}

var (
	pokemon = entity.Pokemon{
		Name:        "Bulbasour",
		Species:     "Pokemon",
		Types:       []int64{1, 2, 3},
		Catched:     1,
		ImageURL:    "https://image.com/image/1",
		Description: "asdf",
		Weight:      0.3,
		Height:      1,
		Stats: entity.Stats{
			HP:     100,
			Attack: 100,
			Def:    100,
			Speed:  100,
		},
	}
)

func TestServer_GetAllPokemon(t *testing.T) {
	prov := serverPorvider()

	type fields struct {
		Router         *httprouter.Router
		PokemonUsecase usecase.PokemonUsecaseItf
		TypeUsecase    usecase.TypeUsecaseItf
		UserUsecase    usecase.UserUsecaseItf
	}
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mock   func()
	}{
		{
			name: "success without query param",
			fields: fields{
				Router:         prov.Router,
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest("GET", "/pokedex/pokemons", nil),
				in2: httprouter.Params{},
			},
			mock: func() {
				prov.PokemonUsecase.On("GetAllPokemon", mock.Anything).
					Return([]entity.PokemonList{{ID: 1, Name: "Bulbasour"}}, nil).Times(1)
			},
		},
		{
			name: "success using query param",
			fields: fields{
				Router:         prov.Router,
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest("GET", "/pokedex/pokemons?name=bulbasour", nil),
				in2: httprouter.Params{},
			},
			mock: func() {
				prov.PokemonUsecase.On("GetAllPokemonByFilter", mock.Anything, mock.Anything).
					Return([]entity.PokemonList{{ID: 1, Name: "Bulbasour"}}, nil).Times(1)
			},
		},
		{
			name: "failed get pokemon without query param",
			fields: fields{
				Router:         prov.Router,
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest("GET", "/pokedex/pokemons", nil),
				in2: httprouter.Params{},
			},
			mock: func() {
				prov.PokemonUsecase.On("GetAllPokemon", mock.Anything).
					Return(nil, errors.New("error")).Times(1)
			},
		},
		{
			name: "failed get pokemon using query param",
			fields: fields{
				Router:         prov.Router,
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest("GET", "/pokedex/pokemons?name=bulbasour", nil),
				in2: httprouter.Params{},
			},
			mock: func() {
				prov.PokemonUsecase.On("GetAllPokemonByFilter", mock.Anything, mock.Anything).
					Return(nil, errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Router:         tt.fields.Router,
				PokemonUsecase: tt.fields.PokemonUsecase,
				TypeUsecase:    tt.fields.TypeUsecase,
				UserUsecase:    tt.fields.UserUsecase,
			}
			s.GetAllPokemon(tt.args.w, tt.args.r, tt.args.in2)
		})
	}
}

func TestServer_GetPokemonByID(t *testing.T) {
	prov := serverPorvider()

	type fields struct {
		Router         *httprouter.Router
		PokemonUsecase usecase.PokemonUsecaseItf
		TypeUsecase    usecase.TypeUsecaseItf
		UserUsecase    usecase.UserUsecaseItf
	}
	type args struct {
		w     http.ResponseWriter
		r     *http.Request
		param httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mock   func()
	}{
		{
			name: "success",
			fields: fields{
				Router:         prov.Router,
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("GET", "/pokedex/pokemons/:id", nil),
				param: httprouter.Params{{Key: "id", Value: "1"}},
			},
			mock: func() {
				prov.PokemonUsecase.On("GetPokemonByID", mock.Anything, mock.Anything).
					Return(&entity.PokemonDetail{ID: 1, Name: "Bulbasour"}, nil).Times(1)
			},
		},
		{
			name: "failed parsing param",
			fields: fields{
				Router:         prov.Router,
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("GET", "/pokedex/pokemons/:id", nil),
				param: httprouter.Params{{Key: "id", Value: "asdf"}},
			},
			mock: func() {},
		},
		{
			name: "failed get pokemon by id",
			fields: fields{
				Router:         prov.Router,
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("GET", "/pokedex/pokemons/:id", nil),
				param: httprouter.Params{{Key: "id", Value: "1"}},
			},
			mock: func() {
				prov.PokemonUsecase.On("GetPokemonByID", mock.Anything, mock.Anything).
					Return(nil, errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Router:         tt.fields.Router,
				PokemonUsecase: tt.fields.PokemonUsecase,
				TypeUsecase:    tt.fields.TypeUsecase,
				UserUsecase:    tt.fields.UserUsecase,
			}
			s.GetPokemonByID(tt.args.w, tt.args.r, tt.args.param)
		})
	}
}

func TestServer_CatchPokemon(t *testing.T) {
	prov := serverPorvider()

	type fields struct {
		Router         *httprouter.Router
		PokemonUsecase usecase.PokemonUsecaseItf
		TypeUsecase    usecase.TypeUsecaseItf
		UserUsecase    usecase.UserUsecaseItf
	}
	type args struct {
		w     http.ResponseWriter
		r     *http.Request
		param httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mock   func()
	}{
		{
			name: "success",
			fields: fields{
				Router:         prov.Router,
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("POST", "/user/pokedex/pokemons/:id/catch", nil),
				param: httprouter.Params{{Key: "id", Value: "1"}},
			},
			mock: func() {
				prov.PokemonUsecase.On("CatchPokemon", mock.Anything, mock.Anything).
					Return(nil).Times(1)
			},
		},
		{
			name: "failed parse integer",
			fields: fields{
				Router:         prov.Router,
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("POST", "/user/pokedex/pokemons/:id/catch", nil),
				param: httprouter.Params{{Key: "id", Value: "abc"}},
			},
			mock: func() {},
		},
		{
			name: "failed to catch pokemon",
			fields: fields{
				Router:         prov.Router,
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("POST", "/user/pokedex/pokemons/:id/catch", nil),
				param: httprouter.Params{{Key: "id", Value: "1"}},
			},
			mock: func() {
				prov.PokemonUsecase.On("CatchPokemon", mock.Anything, mock.Anything).
					Return(errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Router:         tt.fields.Router,
				PokemonUsecase: tt.fields.PokemonUsecase,
				TypeUsecase:    tt.fields.TypeUsecase,
				UserUsecase:    tt.fields.UserUsecase,
			}
			s.CatchPokemon(tt.args.w, tt.args.r, tt.args.param)
		})
	}
}

func TestServer_CreatePokemon(t *testing.T) {
	prov := serverPorvider()
	body, _ := json.Marshal(pokemon)

	type fields struct {
		Router         *httprouter.Router
		PokemonUsecase usecase.PokemonUsecaseItf
		TypeUsecase    usecase.TypeUsecaseItf
		UserUsecase    usecase.UserUsecaseItf
	}
	type args struct {
		w     http.ResponseWriter
		r     *http.Request
		param httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mock   func()
	}{
		{
			name: "success",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("POST", "/pokedex/pokemons", bytes.NewBuffer(body)),
				param: httprouter.Params{},
			},
			mock: func() {
				prov.PokemonUsecase.On("CreatePokemon", mock.Anything, mock.Anything).
					Return(int64(1), nil).Times(1)
			},
		},
		{
			name: "failed create pokemon",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("POST", "/pokedex/pokemons", bytes.NewBuffer(body)),
				param: httprouter.Params{},
			},
			mock: func() {
				prov.PokemonUsecase.On("CreatePokemon", mock.Anything, mock.Anything).
					Return(int64(0), errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Router:         tt.fields.Router,
				PokemonUsecase: tt.fields.PokemonUsecase,
				TypeUsecase:    tt.fields.TypeUsecase,
				UserUsecase:    tt.fields.UserUsecase,
			}
			s.CreatePokemon(tt.args.w, tt.args.r, tt.args.param)
		})
	}
}

func TestServer_UpdatePokemon(t *testing.T) {
	prov := serverPorvider()
	body, _ := json.Marshal(pokemon)

	type fields struct {
		Router         *httprouter.Router
		PokemonUsecase usecase.PokemonUsecaseItf
		TypeUsecase    usecase.TypeUsecaseItf
		UserUsecase    usecase.UserUsecaseItf
	}
	type args struct {
		w     http.ResponseWriter
		r     *http.Request
		param httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mock   func()
	}{
		{
			name: "success",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("PUT", "/pokedex/pokemons/:id", bytes.NewBuffer(body)),
				param: httprouter.Params{{Key: "id", Value: "1"}},
			},
			mock: func() {
				prov.PokemonUsecase.On("UpdatePokemon", mock.Anything, mock.Anything, mock.Anything).
					Return(&entity.PokemonDetail{}, nil).Times(1)
			},
		},
		{
			name: "failed parse int",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("PUT", "/pokedex/pokemons/:id", bytes.NewBuffer(body)),
				param: httprouter.Params{{Key: "id", Value: "abc"}},
			},
			mock: func() {},
		},
		{
			name: "failed update pokemon",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("PUT", "/pokedex/pokemons/:id", bytes.NewBuffer(body)),
				param: httprouter.Params{{Key: "id", Value: "1"}},
			},
			mock: func() {
				prov.PokemonUsecase.On("UpdatePokemon", mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Router:         tt.fields.Router,
				PokemonUsecase: tt.fields.PokemonUsecase,
				TypeUsecase:    tt.fields.TypeUsecase,
				UserUsecase:    tt.fields.UserUsecase,
			}
			s.UpdatePokemon(tt.args.w, tt.args.r, tt.args.param)
		})
	}
}

func TestServer_DeletePokemon(t *testing.T) {
	prov := serverPorvider()

	type fields struct {
		Router         *httprouter.Router
		PokemonUsecase usecase.PokemonUsecaseItf
		TypeUsecase    usecase.TypeUsecaseItf
		UserUsecase    usecase.UserUsecaseItf
	}
	type args struct {
		w     http.ResponseWriter
		r     *http.Request
		param httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mock   func()
	}{
		{
			name: "success",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("DELETE", "/internal/pokedex/pokemons/:id", nil),
				param: httprouter.Params{{Key: "id", Value: "1"}},
			},
			mock: func() {
				prov.PokemonUsecase.On("DeletePokemon", mock.Anything, mock.Anything).
					Return(nil).Times(1)
			},
		},
		{
			name: "failed parse int",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("DELETE", "/internal/pokedex/pokemons/:id", nil),
				param: httprouter.Params{{Key: "id", Value: "abc"}},
			},
			mock: func() {},
		},
		{
			name: "failed delete pokemon",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("DELETE", "/internal/pokedex/pokemons/:id", nil),
				param: httprouter.Params{{Key: "id", Value: "1"}},
			},
			mock: func() {
				prov.PokemonUsecase.On("DeletePokemon", mock.Anything, mock.Anything).
					Return(errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Router:         tt.fields.Router,
				PokemonUsecase: tt.fields.PokemonUsecase,
				TypeUsecase:    tt.fields.TypeUsecase,
				UserUsecase:    tt.fields.UserUsecase,
			}
			s.DeletePokemon(tt.args.w, tt.args.r, tt.args.param)
		})
	}
}

func TestServer_Register(t *testing.T) {
	prov := serverPorvider()

	usernameEmpty := entity.User{
		Username: "",
		Password: "123",
	}

	passwordEmpty := entity.User{
		Username: "winarto",
		Password: "",
	}

	correctUser := entity.User{
		Username: "winarto",
		Password: "123",
	}

	bodyUsernameEmtpy, _ := json.Marshal(usernameEmpty)
	bodyPasswordEmpty, _ := json.Marshal(passwordEmpty)
	bodyCorrectUser, _ := json.Marshal(correctUser)

	type fields struct {
		Router         *httprouter.Router
		PokemonUsecase usecase.PokemonUsecaseItf
		TypeUsecase    usecase.TypeUsecaseItf
		UserUsecase    usecase.UserUsecaseItf
	}
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mock   func()
	}{
		{
			name: "success",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest("POST", "/register", bytes.NewBuffer(bodyCorrectUser)),
				in2: httprouter.Params{},
			},
			mock: func() {
				prov.UserUsecase.On("Register", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(int64(1), nil).Times(1)
			},
		},
		{
			name: "failed username empty",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest("POST", "/register", bytes.NewBuffer(bodyUsernameEmtpy)),
				in2: httprouter.Params{},
			},
			mock: func() {},
		},
		{
			name: "failed password empty",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest("POST", "/register", bytes.NewBuffer(bodyPasswordEmpty)),
				in2: httprouter.Params{},
			},
			mock: func() {},
		},
		{
			name: "failed register user",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest("POST", "/register", bytes.NewBuffer(bodyCorrectUser)),
				in2: httprouter.Params{},
			},
			mock: func() {
				prov.UserUsecase.On("Register", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(int64(0), errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Router:         tt.fields.Router,
				PokemonUsecase: tt.fields.PokemonUsecase,
				TypeUsecase:    tt.fields.TypeUsecase,
				UserUsecase:    tt.fields.UserUsecase,
			}
			s.Register(tt.args.w, tt.args.r, tt.args.in2)
		})
	}
}

func TestServer_Login(t *testing.T) {
	prov := serverPorvider()

	usernameEmpty := entity.User{
		Username: "",
		Password: "123",
	}

	passwordEmpty := entity.User{
		Username: "winarto",
		Password: "",
	}

	correctUser := entity.User{
		Username: "winarto",
		Password: "123",
	}

	bodyUsernameEmtpy, _ := json.Marshal(usernameEmpty)
	bodyPasswordEmpty, _ := json.Marshal(passwordEmpty)
	bodyCorrectUser, _ := json.Marshal(correctUser)

	type fields struct {
		Router         *httprouter.Router
		PokemonUsecase usecase.PokemonUsecaseItf
		TypeUsecase    usecase.TypeUsecaseItf
		UserUsecase    usecase.UserUsecaseItf
	}
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mock   func()
	}{
		{
			name: "success",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest("POST", "/login", bytes.NewBuffer(bodyCorrectUser)),
				in2: httprouter.Params{},
			},
			mock: func() {
				prov.UserUsecase.On("Login", mock.Anything, mock.Anything, mock.Anything).
					Return("token", nil).Times(1)
			},
		},
		{
			name: "failed username empty",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest("POST", "/login", bytes.NewBuffer(bodyUsernameEmtpy)),
				in2: httprouter.Params{},
			},
			mock: func() {},
		},
		{
			name: "failed password empty",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest("POST", "/login", bytes.NewBuffer(bodyPasswordEmpty)),
				in2: httprouter.Params{},
			},
			mock: func() {},
		},
		{
			name: "failed login user",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest("POST", "/login", bytes.NewBuffer(bodyCorrectUser)),
				in2: httprouter.Params{},
			},
			mock: func() {
				prov.UserUsecase.On("Login", mock.Anything, mock.Anything, mock.Anything).
					Return("", errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Router:         tt.fields.Router,
				PokemonUsecase: tt.fields.PokemonUsecase,
				TypeUsecase:    tt.fields.TypeUsecase,
				UserUsecase:    tt.fields.UserUsecase,
			}
			s.Login(tt.args.w, tt.args.r, tt.args.in2)
		})
	}
}

func TestServer_Logout(t *testing.T) {
	prov := serverPorvider()
	type fields struct {
		Router         *httprouter.Router
		PokemonUsecase usecase.PokemonUsecaseItf
		TypeUsecase    usecase.TypeUsecaseItf
		UserUsecase    usecase.UserUsecaseItf
	}
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest("POST", "/logout", nil),
				in2: httprouter.Params{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Router:         tt.fields.Router,
				PokemonUsecase: tt.fields.PokemonUsecase,
				TypeUsecase:    tt.fields.TypeUsecase,
				UserUsecase:    tt.fields.UserUsecase,
			}
			s.Logout(tt.args.w, tt.args.r, tt.args.in2)
		})
	}
}

func TestServer_Healthz(t *testing.T) {
	prov := serverPorvider()

	type fields struct {
		Router         *httprouter.Router
		PokemonUsecase usecase.PokemonUsecaseItf
		TypeUsecase    usecase.TypeUsecaseItf
		UserUsecase    usecase.UserUsecaseItf
	}
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				Router:         httprouter.New(),
				PokemonUsecase: prov.PokemonUsecase,
				TypeUsecase:    prov.TypeUsecase,
				UserUsecase:    prov.UserUsecase,
			},
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest("POST", "/healthz", nil),
				in2: httprouter.Params{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Router:         tt.fields.Router,
				PokemonUsecase: tt.fields.PokemonUsecase,
				TypeUsecase:    tt.fields.TypeUsecase,
				UserUsecase:    tt.fields.UserUsecase,
			}
			s.Healthz(tt.args.w, tt.args.r, tt.args.in2)
		})
	}
}
