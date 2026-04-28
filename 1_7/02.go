package main

import (
	"fmt"
	"strings"
)

type CityList []string

func (c *CityList) add(citiesToAdd ...string) error {
	for _, city := range citiesToAdd {
		if c.contains(city) {
			return fmt.Errorf("дублирование города %s", city)
		}
	}
	*c = append(*c, citiesToAdd...)

	return nil
}

func (c *CityList) contains(city string) bool {
	for _, entry := range *c {
		if entry == city {
			return true
		}
	}

	return false
}

func (c *CityList) remove(city string) error {
	var result CityList
	if !c.contains(city) {
		return fmt.Errorf("город %s отсутствует в списке", city)
	}
	for _, entry := range *c {
		if entry != city {
			result = append(result, entry)
		}
	}
	*c = result

	return nil
}

func (c *CityList) search(query string) CityList {
	var result CityList
	for _, entry := range *c {
		if strings.Contains(strings.ToLower(entry), strings.ToLower(query)) {
			result = append(result, entry)
		}
	}
	return result
}

func main() {
	var cities CityList

	err := cities.add("Москва", "Будапешт", "Братислава", "Варшава", "Белград")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cities)
	err = cities.remove("Будапешт")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cities)
	fmt.Println(cities.search("ава"))
}
