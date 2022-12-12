def getPrio(char):
    if char >= 'a' and char <= 'z':
        return (ord(char)-96)
    else:
        return (ord(char)-38)

result = 0
with open("./2022/3/input") as f:
    lines = f.read().splitlines()

    for line in lines:
        firstHalf = line[:len(line)//2]
        secondHalf = line[len(line)//2:]

        for char in firstHalf:
            if char in secondHalf:
                result += getPrio(char)
                break

print(result)
