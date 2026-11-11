package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Pozitie struct {
	x, y float64
}

type Runda struct {
	canaleIntrare []chan Pozitie
	canaleIesire  []chan Pozitie
}

type Robot struct {
	id               int
	poz              Pozitie
	raza             int // raza pana la care vede un robot
	rounds           int
	chTopologiiRunde chan Runda
}

func (robot *Robot) send(runda Runda) {
	for _, canal := range runda.canaleIesire {
		canal <- robot.poz
	}
}

func (robot *Robot) recv(runda Runda) []Pozitie {
	pozitii := make([]Pozitie, len(runda.canaleIntrare))

	for _, canal := range runda.canaleIntrare {
		pozitii = append(pozitii, <-canal)
	}

	return pozitii
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
func updateComunicare(roboti []Robot) ([]Runda, int) {
	runde := make([]Runda, len(roboti))
	muchii := 0

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
				ch1 := make(chan Pozitie, 1) // de la i la j
				ch2 := make(chan Pozitie, 1) // de la j la i

				runde[i].canaleIesire = append(robot1.canaleIesire, ch1)
				runde[j].canaleIntrare = append(robot1.canaleIntrare, ch1)
				runde[j].canaleIesire = append(robot2.canaleIesire, ch2)
				runde[i].canaleIntrare = append(robot2.canaleIntrare, ch2)
				muchii++
			}
		}
	}
	return runde, muchii
}

func (robot *Robot) transformare(terminat chan<- struct{}, _wg *sync.WaitGroup) {
	defer _wg.Done()

	for topologieCurenta := range robot.chTopologiiRunde {
		robot.send(topologieCurenta)
		received := robot.recv(topologieCurenta)

		sumaX := robot.poz.x
		sumaY := robot.poz.y

		for _, pozitie := range received {
			sumaX += pozitie.x
			sumaY += pozitie.y
		}
		robot.poz.x = sumaX / float64(len(received)+1)
		robot.poz.y = sumaY / float64(len(received)+1)

		robot.rounds++
		terminat <- struct{}{}
	}
}

func returnDiametru(roboti []Robot) float64 {
	maxim := 0.0
	for i := 0; i < len(roboti); i++ {
		x1, y1 := roboti[i].poz.x, roboti[i].poz.y
		for j := i + 1; j < len(roboti); j++ {
			x2, y2 := roboti[j].poz.x, roboti[j].poz.y
			distanta := math.Sqrt(math.Pow((x1-x2), 2) + math.Pow((y1-y2), 2))
			maxim = math.Max(maxim, distanta)
		}
	}

	return maxim
}

func printRoboti(n int, roboti []Robot) {
	fmt.Println(n)
	for _, robot := range roboti {
		fmt.Printf("robotul %d(%.2f, %.2f)", robot.id, robot.poz.x, robot.poz.y)
	}
	fmt.Println()
}

func main() {
	roboti := make([]Robot, 0)
	r := 11

	input, err := os.Open("input.in")
	if err != nil {
		fmt.Println("Eroare la citire")
		return
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
		roboti[i].canaleIesire = make([]chan Pozitie, 0)
		roboti[i].canaleIntrare = make([]chan Pozitie, 0)
	}

	printRoboti(n, roboti)
	updateComunicare(roboti)
	printRoboti(n, roboti)
}
