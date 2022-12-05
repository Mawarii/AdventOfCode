import re

with open("./5/input-crane") as startFile:
    allBoxes = startFile.read().splitlines()

    allRows = list()
    for x in range(1, len(allBoxes[0]), 4):
        temp = list()
        for boxNumber in allBoxes:
            if ' ' not in boxNumber[x]:
                temp.append(boxNumber[x])
        allRows.append(temp)

with open("./5/input") as f:
    lines = f.read().splitlines()

    numbersOnly = list()
    for line in lines:
        numbersOnly.append(re.findall(r'\d+', line))

for move in numbersOnly:
    for x in range(0, int(move[0])):
        allRows[int(move[2]) - 1].insert(0, allRows[int(move[1]) - 1][0])
        allRows[int(move[1]) - 1].pop(0)

for row in allRows:
    print(row[0], end='')
