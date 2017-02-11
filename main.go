package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type Command struct {
	Direction string
	Mode      string
	Order     int
}

type Commands []Command

func (a Commands) Len() int {
	return len(a)
}

func (a Commands) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Commands) Less(i, j int) bool {
	return a[i].Order < a[j].Order
}

func main() {
	resp, err := http.Get("https://raw.githubusercontent.com/WWGLondon/graduation/solution/map/release_party_map.json")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	commands := Commands{}
	err = json.Unmarshal(data, &commands)

	sort.Sort(commands)

	for _, command := range commands {
		fmt.Println(command)
	}

}
