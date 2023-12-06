#!/bin/bash

main() {
    sum=0

    input="/home/mawari/Desktop/git-projects/private/AdventOfCode/2023/1/input"

    while IFS= read -r line; do
        line=$(echo "$line" | sed -e 's/twone/twoone/g; s/oneight/oneeight/g; s/threeight/threeeight/g; s/fiveight/fiveeight/g; s/sevenine/sevennine/g; s/eightwo/eighttwo/g; s/eighthree/eightthree/g; s/nineight/nineeight/g')

        number=$(grep -oE 'one|two|three|four|five|six|seven|eight|nine|[0-9]' <<< "$line" | head -1 | tr -d '\n'; grep -oE 'one|two|three|four|five|six|seven|eight|nine|[0-9]' <<< "$line" | tail -1)

        number=$(echo "$number" | sed -e 's/one/1/g; s/two/2/g; s/three/3/g; s/four/4/g; s/five/5/g; s/six/6/g; s/seven/7/g; s/eight/8/g; s/nine/9/g')
        
        sum=$(("$number"+"$sum"))
    done < "$input"

    echo "$sum"
}

main
