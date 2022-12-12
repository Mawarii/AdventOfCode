results = {
    'A X' : 3,
    'A Y' : 4,
    'A Z' : 8,
    'B X' : 1,
    'B Y' : 5,
    'B Z' : 9,
    'C X' : 2,
    'C Y' : 6,
    'C Z' : 7,
}

finalValue = 0
with open("./2022/2/input") as f:
    lines = f.readlines()

    for line in lines:
        for key,val in results.items():
            if key in line:
                finalValue += val
    print(finalValue)
