package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type GunInfo struct {
	Name         string `json:"Name"`
	Rarity       string `json:"Rarity"`
	Manufacturer string `json:"Manufacturer"`
	Type         string `json:"Type"`
	Elements     string `json:"Elements"`
	Content      string `json:"Content"`
}

type ShieldInfo struct {
	Name         string `json:"Name"`
	Rarity       string `json:"Rarity"`
	Manufacturer string `json:"Manufacturer"`
	Type         string `json:"Type"`
	Elements     string `json:"Elements"`
	Content      string `json:"Content"`
}

type GrenadeInfo struct {
	Name         string `json:"Name"`
	Rarity       string `json:"Rarity"`
	Manufacturer string `json:"Manufacturer"`
	Type         string `json:"Type"`
	Elements     string `json:"Elements"`
	Content      string `json:"Content"`
}

type RelicInfo struct {
	Name    string `json:"Name"`
	Rarity  string `json:"Rarity"`
	Type    string `json:"Type"`
	Content string `json:"Content"`
}

type ClassModInfo struct {
	Name         string `json:"Name"`
	Rarity       string `json:"Rarity"`
	Manufacturer string `json:"Manufacturer"`
	Class        string `json:"Class"`
	Content      string `json:"Content"`
}

func main() {

	getWeapons()
	getShields()
	getRelics()
	getGrenades()
	getClassMods()

}

func getWeapons() {
	// Initialize list of gun information
	allGunsInfo := []GunInfo{}

	url := "https://www.lootlemon.com/db/borderlands-2/weapons"
	domains := []string{"www.lootlemon.com"}

	// Instantiate default collector
	client := colly.NewCollector(colly.AllowedDomains(domains...))

	// Grab gun information and append to the 'allGunsInfo' struct
	client.OnHTML("div.db_item.w-dyn-item", func(element *colly.HTMLElement) {
		weaponName := element.Attr("data-name")
		weaponRarity := element.Attr("data-rarity")
		weaponManufacturer := element.Attr("data-manufacturer")
		weaponType := element.Attr("data-type")
		weaponElements := element.Attr("data-elements")
		weaponContent := element.Attr("data-content")
		i := GunInfo{
			Name:         weaponName,
			Rarity:       weaponRarity,
			Manufacturer: weaponManufacturer,
			Type:         weaponType,
			Elements:     weaponElements,
			Content:      weaponContent,
		}
		allGunsInfo = append(allGunsInfo, i)
	})

	client.Visit(url)
	data, err := json.MarshalIndent(allGunsInfo, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func getShields() {
	// Initialize list of shield information
	allShieldsInfo := []ShieldInfo{}

	url := "https://www.lootlemon.com/db/borderlands-2/shields"
	domains := []string{"www.lootlemon.com"}

	// Instantiate default collector
	client := colly.NewCollector(colly.AllowedDomains(domains...))

	// Grab shield information and append to the 'allShieldsInfo' struct
	client.OnHTML("div.db_item.w-dyn-item", func(element *colly.HTMLElement) {
		weaponName := element.Attr("data-name")
		weaponRarity := element.Attr("data-rarity")
		weaponManufacturer := element.Attr("data-manufacturer")
		weaponType := element.Attr("data-type")
		weaponElements := element.Attr("data-elements")
		weaponContent := element.Attr("data-content")
		i := ShieldInfo{
			Name:         weaponName,
			Rarity:       weaponRarity,
			Manufacturer: weaponManufacturer,
			Type:         weaponType,
			Elements:     weaponElements,
			Content:      weaponContent,
		}
		allShieldsInfo = append(allShieldsInfo, i)
	})

	client.Visit(url)
	data, err := json.MarshalIndent(allShieldsInfo, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func getGrenades() {
	// Initialize list of shield information
	allGrenadesInfo := []GrenadeInfo{}

	url := "https://www.lootlemon.com/db/borderlands-2/grenade-mods"
	domains := []string{"www.lootlemon.com"}

	// Instantiate default collector
	client := colly.NewCollector(colly.AllowedDomains(domains...))

	// Grab shield information and append to the 'allGrenadesInfo' struct
	client.OnHTML("div.db_item.w-dyn-item", func(element *colly.HTMLElement) {
		weaponName := element.Attr("data-name")
		weaponRarity := element.Attr("data-rarity")
		weaponManufacturer := element.Attr("data-manufacturer")
		weaponType := element.Attr("data-type")
		weaponElements := element.Attr("data-elements")
		weaponContent := element.Attr("data-content")
		i := GrenadeInfo{
			Name:         weaponName,
			Rarity:       weaponRarity,
			Manufacturer: weaponManufacturer,
			Type:         weaponType,
			Elements:     weaponElements,
			Content:      weaponContent,
		}
		allGrenadesInfo = append(allGrenadesInfo, i)
	})

	client.Visit(url)
	data, err := json.MarshalIndent(allGrenadesInfo, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func getRelics() {
	// Initialize list of shield information
	allRelicsInfo := []RelicInfo{}

	url := "https://www.lootlemon.com/db/borderlands-2/relics"
	domains := []string{"www.lootlemon.com"}

	// Instantiate default collector
	client := colly.NewCollector(colly.AllowedDomains(domains...))

	// Grab shield information and append to the 'allRelicsInfo' struct
	client.OnHTML("div.db_item.w-dyn-item", func(element *colly.HTMLElement) {
		weaponName := element.Attr("data-name")
		weaponRarity := element.Attr("data-rarity")
		weaponType := element.Attr("data-type")
		weaponContent := element.Attr("data-content")
		i := RelicInfo{
			Name:    weaponName,
			Rarity:  weaponRarity,
			Type:    weaponType,
			Content: weaponContent,
		}
		allRelicsInfo = append(allRelicsInfo, i)
	})

	client.Visit(url)
	data, err := json.MarshalIndent(allRelicsInfo, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func getClassMods() {
	// Initialize list of shield information
	allClassModsInfo := []ClassModInfo{}

	url := "https://www.lootlemon.com/db/borderlands-2/relics"
	domains := []string{"www.lootlemon.com"}

	// Instantiate default collector
	client := colly.NewCollector(colly.AllowedDomains(domains...))

	// Grab shield information and append to the 'allClassModsInfo' struct
	client.OnHTML("div.db_item.w-dyn-item", func(element *colly.HTMLElement) {
		weaponName := element.Attr("data-name")
		weaponRarity := element.Attr("data-rarity")
		weaponManufacturer := element.Attr("data-manufacturer")
		weaponClass := element.Attr("data-class")
		weaponContent := element.Attr("data-content")
		i := ClassModInfo{
			Name:         weaponName,
			Rarity:       weaponRarity,
			Manufacturer: weaponManufacturer,
			Class:        weaponClass,
			Content:      weaponContent,
		}
		allClassModsInfo = append(allClassModsInfo, i)
	})

	client.Visit(url)
	data, err := json.MarshalIndent(allClassModsInfo, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
