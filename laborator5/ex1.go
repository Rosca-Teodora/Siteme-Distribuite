package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Pozitie struct {
	x, y float64
}

type Conexiune struct { // pastrez conexiunea mai clar astfel incat sa pot identifica usor robotii care comunica
	ch       chan Pozitie
	src, dst int // doar id-ul robotului
}

type Robot struct {
	id         int
	poz        Pozitie
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
func updateComunicare(roboti []Robot) {
	for i := 0; i < len(roboti)-1; i++ {
		var robot1 *Robot
		robot1 = &roboti[i]
		x1, y1 := robot1.poz.x, robot1.poz.y

		for j := i + 1; j < len(roboti); j++ {
			var robot2 *Robot
			robot2 = &roboti[j]
			x2, y2 := robot2.poz.x, robot2.poz.y

			distDeLa1La2 := math.Sqrt(math.Pow((x1-x2), 2) + math.Pow((y1-y2), 2))
			fmt.Printf("pt robotul cu id-ul %d si id-ul %d distanta este: %.2f\n", robot1.id, robot2.id, distDeLa1La2)
			if distDeLa1La2 <= float64(robot1.raza) {
				newCh := make(chan Pozitie)
				robot1.comunicari = append(robot1.comunicari, Conexiune{ch: newCh, src: robot1.id, dst: robot2.id})
				robot2.comunicari = append(robot2.comunicari, Conexiune{ch: newCh, src: robot2.id, dst: robot1.id})
			}
		}
	}
}

func printRoboti(n int, roboti []Robot) {
	fmt.Println(n)
	for _, robot := range roboti {
		fmt.Println(robot)
	}
}

func main() {
	roboti := make([]Robot, 0)
	r := 11

	input, err := os.Open("input.in")
	if err != nil {
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

		xC, _ := strconv.ParseFloat(line[0], 64)
		yC, _ := strconv.ParseFloat(line[1], 64)
		roboti = append(roboti, Robot{id: i + 1, poz: Pozitie{xC, yC}, raza: r})
		roboti[i].comunicari = make([]Conexiune, 0)
	}

	printRoboti(n, roboti)
	updateComunicare(roboti)
	printRoboti(n, roboti)
}
