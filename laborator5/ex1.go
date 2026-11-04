package laborator5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pozitie struct {
	x, y int
}

type Conexiune struct { // pastrez conexiunea mai clar astfel incat sa pot identifica usor robotii care comunica
	ch         chan Pozitie
	nod1, nod2 Robot
}

type Robot struct {
	x, y       int
	raza       int // raza pana la care vede un robot
	comunicari []Conexiune
}

func conversie(text string) int {
	nr, er := strconv.Atoi(text)
	if er != nil {
		fmt.Println("Eroare la conversie")
		return -1
	} else {
		return nr
	}
}

// simularea unui senzor care sa reactualizeze conexiunile dintre roboti
func updateComunicare(roboti []Robot, raza int) {
	for i := 0; i < len(roboti)-1; i++ {
		robot1 := roboti[i]

		for j := i + 1; j < len(roboti); j++ {

		}
	}
}

func main() {
	roboti := make([]Robot, 0)
	r := 5

	input, err := os.Open("input.in")
	if !err {
		fmt.Println("Eroare la citire")
	}
	defer input.Close()

	// citire valori
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	n := conversie(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())

		xC := conversie(line[0])
		yC := conversie(line[1])
		roboti = append(roboti, Robot{x: xC, y: yC, raza: r})
		roboti[i].comunicari = make([]Conexiune, 0)
	}

}
