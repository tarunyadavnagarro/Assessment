package main

import "fmt"

// A Month specifies a month of the year (January = 1, ...).
type Planet int

const (
              Mercury Planet = 1 + iota
              Venus
              Earth
              Mars
              Jupitar
              Saturn
              Naptune
              Uranus
              Pluto
)

var planet = [12]string{
              "January",
              "Venus",
              "Earth",
              "Mars",
              "Jupitar",
              "Saturn",
              "Naptune",
              "Uranus",
              "Pluto",
}


func (p Planet) String() string { return planet[p-1] }

func main() {

              var planetName Planet
              planetName = 3
              fmt.Printf("%s \n", planetName)

}











