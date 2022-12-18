run: 
	go run ./app/main.go

.PHONY: build
build: 
	@ echo "Build Binary..."
	@ go build -o ./build/go-pokedex app/main.go
	@ echo "\nBuild Binary Success"

start: 
	./build/go-pokedex

build-image:
	@ echo "Docker Build Image"
	@ docker build . -t go_pokedex_app -f ./deployments/Dockerfile
	@ echo "\nDocker Build Fished"

up: 
	@ docker-compose up -d --build mysql
	@ $(MAKE) build-image
	@ docker-compose up app

down:
	docker compose down

destroy:
	@ docker-compose down --remove-orphans -v
	@ printf "Cleanning artifacts... "
	@ rm -f *.out
	@ echo "done."
	@ docker rmi go_pokedex_app

test: 
	go test -v --race ./... 

coverage: 
	go test -v -coverprofile coverage.out --race ./...

show_coverage:
	go tool cover -func coverage.out

html:
	go tool cover -html=coverage.out

generate_mock: 
	@ mockery --dir=repository/pokemon --name=PokemonRepositoryItf --filename=pokemon_mock.go --output=repository/pokemon/mocks --outpkg=pokemonrepositorymock
	@ mockery --dir=repository/pokemontypes --name=PokemonTypeRepositoryItf --filename=pokemon_type_mock.go --output=repository/pokemontypes/mocks --outpkg=pokemontyperepositorymock
	@ mockery --dir=repository/types --name=TypeRepositoryItf --filename=types_mock.go --output=repository/types/mocks --outpkg=typesrepositorymock
	@ mockery --dir=repository/user --name=UserRepositoryItf --filename=user_mock.go --output=repository/user/mocks --outpkg=userrepositorymock
	@ mockery --dir=usecase --name=PokemonUsecaseItf --filename=pokemon_mock.go --output=usecase/mocks --outpkg=usecasemock
	@ mockery --dir=usecase --name=TypeUsecaseItf --filename=type_mock.go --output=usecase/mocks --outpkg=usecasemock
	@ mockery --dir=usecase --name=UserUsecaseItf --filename=user_mock.go --output=usecase/mocks --outpkg=usecasemock