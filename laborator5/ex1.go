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

type Robot struct {
	id            int
	poz           Pozitie
	sendPoz       Pozitie
	raza          int // raza pana la care vede un robot
	canaleIntrare []chan Pozitie
	canaleIesire  []chan Pozitie
	rounds        int
	diametru      int
	epsilon 	  float64
}

func (robot *Robot) send() {
	for _, canal := range robot.canaleIesire {
		canal <- robot.poz
	}
}

func (robot *Robot) recv() []Pozitie{
	pozitii := make([]Pozitie, 0)

	for _, canal := range robot.canaleIntrare {
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
				// fiecare robot are canal prin care primeste de la un alt robot si canal separat prin care trimite
				ch1Send := make(chan Pozitie, 1)
				ch1Recv := make(chan Pozitie, 1)
				ch2Send := make(chan Pozitie, 1)
				ch2Recv := make(chan Pozitie, 1)

				robot1.canaleIesire = append(robot1.canaleIesire, ch1Send)   // deci practic sunt 4 canale ca sa le pastrez "unidirectionale" si sa evit deadlock-ul peste tot
				robot1.canaleIntrare = append(robot1.canaleIntrare, ch1Recv) // sincer mai tb sa ma gandesc la chestia asta, sanse sa schimb, vedem daca merge
				robot2.canaleIesire = append(robot2.canaleIesire, ch2Send)
				robot2.canaleIntrare = append(robot2.canaleIntrare, ch2Recv)
			}
		}
	}
}

func (robot *Robot) updateIndividual(roboti []Robot) {
	x1, y1 := robot.poz.x, robot.poz.y
	canaleNoiIesire := make([]chan Pozitie, 0)
	canaleNoiIntrare := make([]chan Pozitie, 0)

	for i := 0; i < len(roboti); i++ {
		altRobot := &roboti[i]
		x2, y2 := altRobot.poz.x, altRobot.poz.y

		distDeLa1La2 := math.Sqrt(math.Pow((x1-x2), 2) + math.Pow((y1-y2), 2))
		if distDeLa1La2 <= float64(robot.raza) {
			ch1Send := make(chan Pozitie)
			ch1Recv := make(chan Pozitie)

			canaleNoiIesire = append(canaleNoiIesire, ch1Send)
			canaleNoiIntrare = append(canaleNoiIntrare, ch1Recv)
		}
	}
}

func (robot *Robot) pozitieEgalaVecini() bool { // "egala" inseamna o diferenta mai mica de epsilon ft mic

}

func (robot *Robot) transformare(roboti []Robot, _wg *sync.WaitGroup) {
	defer _wg.Done()

	robot.sendPoz = robot.poz

	robot.updateIndividual(roboti)
	robot.send()

	for true {
		robot.rounds++
		received := robot.recv()

		sumaX := robot.poz.x 
		sumaY := robot.poz.y

		for _, pozitie := received {
			sumaX += pozitie.x 
			sumaY += pozitie.y 
		}
		robot.sendPoz.x = sumaX / float64(len(received) + 1)
		robot.sendPoz.y = sumaY / float64(len(received) + 1)

		if robot.rounds % robot.diametru == 0{
			if robot.pozitieEgalaVecini() {
				return
			}
			robot.updateIndividual(roboti)
		}
		robot.send()
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
		roboti[i].canaleIesire = make([]chan Pozitie, 0)
		roboti[i].canaleIntrare = make([]chan Pozitie, 0)
	}

	printRoboti(n, roboti)
	updateComunicare(roboti)
	printRoboti(n, roboti)
}
