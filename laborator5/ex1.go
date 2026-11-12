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
	raza             float64 // raza pana la care vede un robot
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
			fmt.Printf("distanta dintre robot%d robot%d este: %.2f\n", robot1.id, robot2.id, distDeLa1La2)
			if distDeLa1La2 <= float64(robot1.raza) {
				ch1 := make(chan Pozitie, 1) // de la i la j
				ch2 := make(chan Pozitie, 1) // de la j la i

				runde[i].canaleIesire = append(runde[i].canaleIesire, ch1)
				runde[j].canaleIntrare = append(runde[j].canaleIntrare, ch1)
				runde[j].canaleIesire = append(runde[j].canaleIesire, ch2)
				runde[i].canaleIntrare = append(runde[i].canaleIntrare, ch2)
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
		fmt.Printf("robotul %d(%.2f, %.2f)\n", robot.id, robot.poz.x, robot.poz.y)
	}
	fmt.Println()
}

func main() {
	roboti := make([]Robot, 0)
	r := 11.0
	epsilon := 0.00000001
	maxRunde := 1000

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
	}

	printRoboti(n, roboti)
	updateComunicare(roboti)
	printRoboti(n, roboti)

	status := make(chan struct{})
	_wg := sync.WaitGroup{}

	for i := range roboti {
		_wg.Add(1)
		go roboti[i].transformare(status, &_wg)
	}

	muchiiAnterior := 0
	for runda := 1; runda <= maxRunde; runda++ {
		topologieCurenta, muchii := updateComunicare(roboti) // update comunicare returneaza toate rundele
		if muchii != muchiiAnterior {
			muchiiAnterior = muchii
		}

		for i := 0; i < len(roboti); i++ {
			roboti[i].chTopologiiRunde <- topologieCurenta[i]
		}
		for i := 0; i < n; i++ {
			_ = <-status //
		}

		fmt.Printf("Runda %d:\n", runda)
		printRoboti(n, roboti)

		if returnDiametru(roboti) <= epsilon {
			fmt.Printf("Robotii converg intr-un singur punct dupa %d runde", runda)
			break
		}
	}

	for i := range roboti {
		close(roboti[i].chTopologiiRunde)
	}

	_wg.Wait()
}
