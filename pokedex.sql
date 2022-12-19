CREATE DATABASE IF NOT EXISTS `pokedex` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

-- pokedex.pokemons definition

CREATE TABLE IF NOT EXISTS `pokemons` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `species` varchar(255) NOT NULL,
  `catched` int NOT NULL,
  `metadata` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `pokemons` WRITE;
INSERT INTO pokedex.pokemons (id,name,species,catched,metadata) VALUES
	 (1,'Wigglytuff','Balloon Pokemon',0,'{"image_url":"https://img.pokemondb.net/artwork/large/wigglytuff.jpg","description":"Wigglytuff is a Normal/Fairy type Pokémon introduced in Generation 1. It is known as the Balloon Pokemon.","weight":12,"height":1,"stats":{"hp":140,"attack":70,"def":45,"speed":45}}'),
	 (2,'Bulbasaur','Seed Pokemon',1,'{"image_url":"https://img.pokemondb.net/artwork/avif/bulbasaur.avif","description":"Bulbasaur is a Grass/Poison type Pokémon introduced in Generation 1. It is known as the Seed Pokemon.","weight":6.9,"height":0.7,"stats":{"hp":45,"attack":49,"def":49,"speed":45}}'),
	 (3,'Charmander','Lizard Pokemon',0,'{"image_url":"https://img.pokemondb.net/artwork/avif/charmander.avif","description":"Charmander is a Fire type Pokémon introduced in Generation 1. It is known as the Lizard Pokemon.","weight":8.5,"height":0.6,"stats":{"hp":39,"attack":52,"def":43,"speed":65}}');
UNLOCK TABLES;


-- pokedex.types definition

CREATE TABLE IF NOT EXISTS `types` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `types` WRITE;
INSERT INTO pokedex.types (id,name) VALUES
	 (1,'NORMAL'),
	 (2,'GRASS'),
	 (3,'PSYCHIC'),
	 (4,'FLYING'),
	 (5,'FIRE'),
	 (6,'WATER'),
	 (7,'ELECTRIC'),
	 (8,'BUG'),
	 (9,'POISON'),
	 (10,'GROUND');
UNLOCK TABLES;

-- pokedex.pokemon_types definition

CREATE TABLE IF NOT EXISTS `pokemon_types` (
  `id` int NOT NULL AUTO_INCREMENT,
  `pokemon_id` int NOT NULL,
  `types_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `pokemon_types` WRITE;
INSERT INTO pokedex.pokemon_types (id,pokemon_id,types_id) VALUES
	 (1,1,1),
	 (2,2,1),
	 (3,1,0),
	 (4,2,9),
	 (5,3,1),
	 (6,3,5),
	 (7,3,0),
	 (8,3,0),
	 (9,3,0),
	 (10,3,0);
UNLOCK TABLES;

-- pokedex.users definition

CREATE TABLE IF NOT EXISTS `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` text NOT NULL,
  `role` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `users` WRITE;
INSERT INTO pokedex.users (id,username,email,password,`role`) VALUES
	 (1,'admin','admin@mail','$2a$14$nKK/x8BuCSunEa/hGFvLw.Bou4I.chXde4gWwS6L9/X25wQsDXyCC',2),
	 (2,'user','user@mail','$2a$14$.McC4pQLD49wo3Oq7i3sV.xqWGOkfZ/lbVn9dYwBkjng0HXhWLcMi',1),
UNLOCK TABLES;