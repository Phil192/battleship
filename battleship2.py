#Battleship game (Morskoy boy)
__author__= 'Filin Vadim'

import random

class Sea(object):
    board = []

    def painting_sea(self):
        for number in range(8):
            self.board.append(["~"] * 8)

    def random_direction(self):
        random_number = random.randint(1, 2)
        if random_number == 1:
            return 'row'
        return 'col'

    def searching_free_slots(self, direction, len_ship):
        ship_coordinates = []
        valid_slots = 0
        random_slot = random.randint(0, 7)
        ship_start_random = random.randint(0, 7 - len_ship)

        while valid_slots < len_ship:
            if direction == 'row' and self.board[ship_start_random][random_slot] == '~':

                ship_start_random += 1
                ship_coordinates.append({ship_start_random - 1: random_slot})
                valid_slots += 1
            elif direction == 'col' and self.board[random_slot][ship_start_random] == '~':
                ship_start_random += 1
                ship_coordinates.append({random_slot: ship_start_random - 1})
                valid_slots += 1
            else:
                ship_coordinates = []
                random_slot = random.randint(0, 7)
                ship_start_random = random.randint(0, 7 - len_ship)
                valid_slots = 0

        return ship_coordinates, len_ship


class Warship(Sea):
    used_ships = []

    def unrepeated_ship(self):
        len_ship = random.randint(1, 4)
        while len(self.used_ships) < 4:
            if len_ship in self.used_ships:
                len_ship = random.randint(1, 4)
            else:
                self.used_ships.append(len_ship)

                return len_ship

    def placing_ships_and_dots(self, coordinates_lenship_tuple):
        for coord in coordinates_lenship_tuple[0]:
            for row, col in coord.items():
                self.board[row][col] = str(coordinates_lenship_tuple[1])

                dot_coordinates = [[1, 1], [1, 0], [0, 1], [-1, -1], [-1, 0], [0, -1], [-1, 1], [1, -1]]
                for index in dot_coordinates:
                    if row + index[0] < 0 or row + index[0] > 7 or col + index[1] < 0 or col + index[1] > 7:
                        continue
                    elif self.board[row + index[0]][col + index[1]] != '~':
                        continue
                    else:
                        self.board[row + index[0]][col + index[1]] = '.'

a = Sea()
b = Warship()

a.painting_sea()
b.placing_ships_and_dots(a.searching_free_slots(a.random_direction(), b.unrepeated_ship()))
b.placing_ships_and_dots(a.searching_free_slots(a.random_direction(), b.unrepeated_ship()))
b.placing_ships_and_dots(a.searching_free_slots(a.random_direction(), b.unrepeated_ship()))
b.placing_ships_and_dots(a.searching_free_slots(a.random_direction(), b.unrepeated_ship()))

for x in a.board: print(" ".join(x))
