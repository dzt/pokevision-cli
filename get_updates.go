package main

type Updates struct {
    New []Pokemon
}

func getPokemonIds(pokemons []Pokemon) []string {
    var ids []string
    for _, pokemon := range pokemons {
        ids = append(ids, pokemon.ID)
    }
    return ids
}

func contains(arr []string, item string) bool {
    for _, value := range arr {
        if value == item {
            return true
        }
    }
    return false
}

func GetUpdates(previous, current []Pokemon) Updates {
    previousIds := getPokemonIds(previous)
    var new []Pokemon
    for _, pokemon := range current {
        if !contains(previousIds, pokemon.ID) {
            new = append(new, pokemon)
        }
    }
    return Updates{new}
}
