package main

import "fmt"

func callbackExplore(cfg *config, location string) error {
	fmt.Printf("Exploring %s...", location)
	fmt.Println("")
	resp, err := cfg.pokeapiClient.ListPokemonsInArea(location)

	if err != nil {
		return err
	}

	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
