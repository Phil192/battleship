package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var board [8][8]string

func getRandom(num int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	randomNumber := random.Intn(num)
	return randomNumber
}

func paintingSea() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = "~"
		}
	}
}

func randomDirection() string {
	randomNumber := getRandom(2)
	if randomNumber == 0 {
		return "row"
	}
	return "col"
}

func searchingFreeSlots(direction string, lenShip int) []map[int]int {
	shipCoordinates := make([]map[int]int, 0, lenShip)
	validSlots := 0
	randomSlot := getRandom(8)
	shipStartRandom := getRandom(8 - lenShip)

	for validSlots < lenShip {
		if direction == "row" && board[shipStartRandom][randomSlot] == "~" {
			coord := map[int]int{shipStartRandom: randomSlot}
			shipCoordinates = append(shipCoordinates, coord)
			shipStartRandom++
			validSlots++
		} else if direction == "col" && board[randomSlot][shipStartRandom] == "~" {
			coord := map[int]int{randomSlot: shipStartRandom}
			shipCoordinates = append(shipCoordinates, coord)
			shipStartRandom++
			validSlots++
		} else {
			shipCoordinates = make([]map[int]int, 0, lenShip)
			randomSlot = getRandom(8)
			shipStartRandom = getRandom(8 - lenShip)
			validSlots = 0
		}
	}
	return shipCoordinates
}

var usedShips []int

func unrepeatedShip() int {
	lenShip := getRandom(4) + 1
	for len(usedShips) < 4 {
		usedShips = append(usedShips, lenShip)
		for _, ship := range usedShips {
			if lenShip != ship {
				usedShips = append(usedShips, lenShip)
			} else {
				lenShip = getRandom(4) + 1
			}
		}
	}
	return lenShip
}

func placingShipsAndDots(coordinates []map[int]int) {
	lenShip := len(coordinates)
	dotCoord := [][]int{{1, 1}, {1, 0}, {0, 1}, {-1, -1}, {-1, 0}, {0, -1}, {-1, 1}, {1, -1}}
	for _, coord := range coordinates {
		for row, col := range coord {
			board[row][col] = strconv.Itoa(lenShip)
			for _, value := range dotCoord {
				if (row+value[0]) < 0 || (row+value[0]) > 7 || col+value[1] < 0 || col+value[1] > 7 {
					continue
				} else if board[row+value[0]][col+value[1]] != "~" {
					continue
				} else {
					board[row+value[0]][col+value[1]] = "."
				}
			}
		}
	}
}

func main() {
	paintingSea()
	placingShipsAndDots(searchingFreeSlots(randomDirection(), unrepeatedShip()))
	fmt.Println(board)
}
