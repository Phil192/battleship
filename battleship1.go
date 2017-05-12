package main

import (
	"fmt"
  "math/rand"
  "strconv"
	)


var board [8][8]string

func painting_sea() {
  for i := 0; i < 8; i++ {
    for j := 0; j < 8; j++ {
      board[i][j] = "~"
    }
  }
}

func random_direction() string {
  random_number := rand.Intn(1)
  if random_number == 0 {
    return "row"
  } else {
    return "col"
  }
}

func searching_free_slots(direction string, len_ship int) ([]map[int]int) {
  ship_coordinates := make([]map[int]int, len_ship)
  valid_slots := 0
  random_slot := rand.Intn(7)
  ship_start_random := rand.Intn(7 - len_ship)

  for valid_slots < len_ship {
    if direction == "row" && board[ship_start_random][random_slot] == "~" {
      coord := map[int]int{ship_start_random: random_slot}
      ship_coordinates = append(ship_coordinates, coord)
      ship_start_random++
      valid_slots++
    } else if direction == "col" && board[random_slot][ship_start_random] == "~" {
      coord := map[int]int{random_slot: ship_start_random}
      ship_coordinates = append(ship_coordinates, coord)
      ship_start_random++
      valid_slots++
    } else {
      ship_coordinates = make([]map[int]int, len_ship)
      random_slot = rand.Intn(7)
      ship_start_random = rand.Intn(7 - len_ship)
      valid_slots = 0
    }
  }
  return ship_coordinates
}

var used_ships []int

func unrepeated_ship() int {
  len_ship := rand.Intn(3) + 1
  for len(used_ships) < 4 {
    for _, ship := range used_ships{
      if len_ship == ship {
        len_ship = rand.Intn(3) + 1
      } else {
        used_ships = append(used_ships, len_ship)
      }
    }
  }
  return len_ship
}


func placing_ships_and_dots(coordinates []map[int]int) {
  len_ship := len(coordinates)
  dot_coord := [][]int{{1, 1}, {1, 0}, {0, 1}, {-1, -1}, {-1, 0}, {0, -1}, {-1, 1}, {1, -1}}
  for _, coord := range coordinates {
    for row, col := range coord {
      board[row][col] = strconv.Itoa(len_ship)
      for _, value := range dot_coord {
        if (row + value[0]) < 0 || (row + value[0]) > 0 || col + value[1] < 0 || col + value[1] > 7 {
          continue
        } else if board[row + value[0]][col + value[1]] != "~" {
          continue
        } else {
          board[row + value[0]][col + value[1]] = "."
        }
      }
    }
  }
}


func main(){
  painting_sea()
  placing_ships_and_dots(searching_free_slots(random_direction(), unrepeated_ship()))
  fmt.Println(board)
}
