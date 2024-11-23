# day 2 part 2 asks us to determine based upon a set of random draws of red, green and blue cubes from a bag
# what the minimum number of red, green and blue balls would have been to make a set of draws possible
# multiply these numbers together to get the "power" for each game and sum the powers
import re

cube_quantities = {
    'red': 12,
    'green': 13,
    'blue': 14
}

file_obj = open("input.txt", "r") 
file_data = file_obj.read() 
file_obj.close()

sum = 0

lines = file_data.splitlines()

# iterate through the games. we just need to find maximum red, green and blue draws and multiply them together, summming the results as we go
for lineNum, line in enumerate(lines):
    draws = re.findall('(\\d+) (red|blue|green)', line)
    gamePossible = True
    maxRedCounts = max(map(lambda x: int(x[0]), (filter(lambda x: x[1] == 'red', draws))))
    maxGreenCounts = max(map(lambda x: int(x[0]), (filter(lambda x: x[1] == 'green', draws))))
    maxBlueCounts = max(map(lambda x: int(x[0]), (filter(lambda x: x[1] == 'blue', draws))))
    sum += (maxRedCounts * maxGreenCounts * maxBlueCounts)

print(sum)
