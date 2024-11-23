# day 2 asks us to determine based upon a set of random draws of red, green and blue cubes from a bag, which sets of draws would have been possible
# if only 12 red cubes, 13 green cubes, and 14 blue cubes were in the bag
import re

cube_quantities = {
    'red': 12,
    'green': 13,
    'blue': 14
}

file_obj = open("input.txt", "r") 
file_data = file_obj.read() 
file_obj.close()

# function that filters a combination of colour/number that is over the known limit
def numberWasPossible(number, colour):
    if (cube_quantities[colour] < number):
        return False
    else:
        return True

sum = 0

lines = file_data.splitlines()

# iterate through the games checking for any impossible draws
# we won't parse the game number, but instead assume the lines have the games in sequential order
for lineNum, line in enumerate(lines):
    matches = re.findall('(\\d+) (red|blue|green)', line)
    gamePossible = True
    for cubeReveal in matches:
        if not numberWasPossible(int(cubeReveal[0]), cubeReveal[1]):
            gamePossible = False
    if (gamePossible):
        sum += lineNum + 1

print(sum)
