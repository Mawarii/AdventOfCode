with open("./4/input") as f:
    lines = f.read().splitlines()

    counter = 0
    for line in lines:
        sectionOne = list(range(int(line.split(',')[0].split('-')[0]),int(line.split(',')[0].split('-')[1]) + 1))
        sectionTwo = list(range(int(line.split(',')[1].split('-')[0]), int(line.split(',')[1].split('-')[1]) + 1))

        if len(sectionOne) > len(sectionTwo):
            for numb in sectionOne:
                if numb in sectionTwo:
                    counter += 1
                    break
        else:
            for numb in sectionTwo:
                if numb in sectionOne:
                    counter += 1
                    break

print(counter)
