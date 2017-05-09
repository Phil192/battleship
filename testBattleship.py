import battleship2
import unittest

class TestSea(unittest.TestCase):
    def test_random_direction(self):
        a = battleship2.Sea()
        for num in range(1, 10):
            is_correct_value = (a.random_direction() == 'col' or a.random_direction() == 'row')

            self.assertTrue(is_correct_value, 'text')



if __name__ == '__main__':
    unittest.main()

