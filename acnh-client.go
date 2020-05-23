package acnh

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Villager struct {
	ID          string `json:"id"`
	Name        Name   `json:"name"`
	Personality string `json:"personality"`
	Birthday    string `json:"birthday"`
}

type Name struct {
	English string `json:"name-en"`
}

type VillagerInfo struct {
	villagerData map[string]Villager
}

func (v VillagerInfo) Villager(name string) (Villager, error) {
	for _, v := range v.villagerData {
		if strings.ToLower(v.Name.English) == strings.ToLower(name) {
			return v, nil
		}
	}
	return Villager{}, fmt.Errorf("could not find villager %s", name)
}

func villagers() VillagerInfo {
	resp, err := http.Get("https://acnhapi.com/villagers")
	if err != nil {
		panic("did not make it")
	}
	var Villagers map[string]Villager

	defer resp.Body.Close()
	// f := make(Villagers)
	json.NewDecoder(resp.Body).Decode(&Villagers)
	vi := VillagerInfo{Villagers}
	return vi
}
