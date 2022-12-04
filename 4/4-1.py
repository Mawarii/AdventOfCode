with open("./4/input") as f:
    lines = f.read().splitlines()

    counter = 0
    for line in lines:
        sectionOneStart = int(line.split(',')[0].split('-')[0])
        sectionOneEnd = int(line.split(',')[0].split('-')[1])

        sectionTwoStart = int(line.split(',')[1].split('-')[0])
        sectionTwoEnd = int(line.split(',')[1].split('-')[1])

        if sectionOneStart <= sectionTwoStart and sectionOneEnd >= sectionTwoEnd:
            counter += 1
        elif sectionTwoStart <= sectionOneStart and sectionTwoEnd >= sectionOneEnd:
            counter += 1

print(counter)
