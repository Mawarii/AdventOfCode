with open("./1/input") as f:
    lines = f.readlines()

    sum = 0
    allValues = list()
    for line in lines:
        if line.strip():
            sum += int(line)
        else:
            allValues.append(sum)
            sum = 0
            
print(max(allValues))
