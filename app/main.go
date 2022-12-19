package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/winartodev/go-pokedex/config"
	"github.com/winartodev/go-pokedex/middleware"
	pokemonrepository "github.com/winartodev/go-pokedex/repository/pokemon"
	pokemontypserepository "github.com/winartodev/go-pokedex/repository/pokemontypes"
	typserepository "github.com/winartodev/go-pokedex/repository/types"
	userrepository "github.com/winartodev/go-pokedex/repository/user"
	"github.com/winartodev/go-pokedex/server"
	"github.com/winartodev/go-pokedex/usecase"
)

func main() {
	// initialize config
	cfg := config.NewConfig()

	// make connection to database
	db, err := config.NewDatabase(cfg)
	if err != nil {
		panic(err)
	}

	// generate database
	// err = config.MakeMigrate(db)
	// if err != nil {
	// 	panic(err)
	// }

	defer db.Close()

	// initialize repository
	pokemonRepository := pokemonrepository.NewPokemonRepository(db)
	pokemonTypeRepository := pokemontypserepository.NewPokemonTypeRepository(db)
	typeRepository := typserepository.NewTypeRepository(db)
	userrepository := userrepository.NewUserRepository(db)

	// initialize usecase
	pokemonUsecase := usecase.NewPokemonUsecase(usecase.PokemonUsecase{PokemonRepository: pokemonRepository, PokemonTypeRepository: pokemonTypeRepository})
	typeUsecase := usecase.NewTypeUsecase(usecase.TypeUsecase{TypesRepository: typeRepository})
	userUsecsae := usecase.NewUserUsecase(usecase.UserUsecase{UserRepository: userrepository})

	s := server.Server{
		Router:         httprouter.New(),
		PokemonUsecase: pokemonUsecase,
		TypeUsecase:    typeUsecase,
		UserUsecase:    userUsecsae,
	}

	// internal
	s.Router.GET("/internal/pokedex/pokemons", middleware.Auth(s.GetAllPokemon))
	s.Router.POST("/internal/pokedex/pokemons", middleware.Auth(s.CreatePokemon))
	s.Router.GET("/internal/pokedex/pokemons/:id", middleware.Auth(s.GetPokemonByID))
	s.Router.PUT("/internal/pokedex/pokemons/:id", middleware.Auth(s.UpdatePokemon))
	s.Router.DELETE("/internal/pokedex/pokemons/:id", middleware.Auth(s.DeletePokemon))

	s.Router.GET("/internal/pokedex/types", middleware.Auth(s.GetAllType))
	s.Router.POST("/internal/pokedex/types", middleware.Auth(s.CreateType))
	s.Router.GET("/internal/pokedex/types/:id", middleware.Auth(s.GetTypeByID))
	s.Router.PUT("/internal/pokedex/types/:id", middleware.Auth(s.UpdateType))

	// user
	s.Router.POST("/user/pokedex/pokemons/:id/catch", middleware.Auth(s.CatchPokemon))

	// public
	s.Router.GET("/pokedex/pokemons", s.GetAllPokemon)
	s.Router.GET("/pokedex/pokemons/:id", s.GetPokemonByID)
	s.Router.GET("/pokedex/types", s.GetAllType)

	s.Router.POST("/login", s.Login)
	s.Router.POST("/register", s.Register)
	s.Router.POST("/logout", s.Logout)

	s.Router.GET("/healthz", s.Healthz)

	fmt.Printf("http listen and serve at :%d\n", cfg.Application.Port)
	if err := http.ListenAndServe(":8080", s.Router); err != nil {
		log.Fatal(err)
	}
}
