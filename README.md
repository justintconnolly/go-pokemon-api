# Gotta GO Query 'Em All - The Pokemon API 🥳
Welcome to the Pokemon API! This API is your one-stop shop for getting data on all your favorite Pokemon.

---

***Features:***
- Get Pokemon info by name - Just pass a name like "Pikachu" and get back data like number, types, stats, and more!
- Lightning fast queries - Our API is optimized to fetch your Pokemon data as quick as a Thunderbolt ⚡
- Intuitive responses - Responses are formatted in easy to understand JSON so you can get querying ASAP!
- Search is case-insensitive - No need to worry about capitalization when searching for your fave mons.

---

***Getting Started***
`https://pokeapi.com/api/v1/pokemon?name=pikachu`

It's super easy to query the API. Just make a GET request and pass the pokemon name as a query parameter.

The API will handle capitalizing the name properly so you can type names in any case like "pikachu" or "PIKACHU".

---

***Examples***
`GET` `/api/v1/pokemon?name=bulbasaur`


```
{
  "name": "Bulbasaur",
  "pokedex_number": 1,
  "types": ["grass", "poison"],
  "stats": {
    "hp": 45,
    "attack": 49,
    // etc
  }
}
```



Try it out with your favorite Pokemon! The API returns key info like name, Pokedex number, types, stats, and more.

---

### Local Development
To run the API locally:

Clone the repo
Run go build
Start the server with ./pokemon-api
Make requests to http://localhost:8080/api/v1/pokemon
Now you're ready to catch 'em all!

___

### Contributing
Pull requests are welcome! We'd love help expanding the API and adding more Pokemon data.

---

## Gotta query 'em all!