package main

import (
	"time"
	"encoding/json"
	"log"
	"sync"
)

type Bug struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Location Locations `json:"location"`
	Integer  Integers  `json:"integer"`
}
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type Locations [] Location
type Integers [] int

func main() {
	//cl := Bug{Name: "wenz", Location: Locations{{Latitude: 34.56, Longitude: 17.23}}, Integer: []int{1, 2, 3}}
	//bc, _ := json.Marshal(cl)
	//log.Println(string(bc))
	var bug Bug

	c := true
	wg := new(sync.WaitGroup)
	ch := make(chan Bug, 1000)
	wg.Add(2)
	go func() {
		//var bug Bug
		defer wg.Done()

		for i := 0; i < 100; i++ {
			//var bug Bug
			if c {
				json.Unmarshal([] byte( `{"id":1,"name":"Opeth","location":[{"latitude":66.666,"longitude":17.666}],"integer":[1,2,3]}`), &bug)
				//log.Println(err)
			} else {
				json.Unmarshal([] byte( `{"id":2,"name":"Metal","location":[{"latitude":14.666,"longitude":26.666}],"integer":[4,5,6]}`), &bug)
				//log.Println(err)
			}
			ch <- bug
			//time.Sleep(100 * time.Microsecond)
			c = !c

		}
		log.Println("fin")
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for i := range ch {

			log.Printf("%v\n", i)
			time.Sleep(100 * time.Microsecond)
		}

	}()

	wg.Wait()
}

