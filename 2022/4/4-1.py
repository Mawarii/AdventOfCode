with open("./4/input") as f:
    lines = f.read().splitlines()

    counter = 0
    for line in lines:
        sectionOne = set(range(int(line.split(',')[0].split('-')[0]),int(line.split(',')[0].split('-')[1]) + 1))
        sectionTwo = set(range(int(line.split(',')[1].split('-')[0]), int(line.split(',')[1].split('-')[1]) + 1))

        if sectionOne.issubset(sectionTwo) or sectionTwo.issubset(sectionOne):
            counter += 1

print(counter)
