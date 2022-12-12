def getPrio(char):
    if char >= 'a' and char <= 'z':
        return (ord(char)-96)
    else:
        return (ord(char)-38)

result = 0
with open("./2022/3/input") as f:
    lines = f.read().splitlines()

    for x in range(0, len(lines)-1, 3):
        elfGroups = [lines[x], lines[x + 1], lines[x + 2]]
        elfGroups.sort(key=len, reverse=True)

        for char in elfGroups[0]:
            if char in elfGroups[1] and char in elfGroups[2]:
                result += getPrio(char)
                break

print(result)
