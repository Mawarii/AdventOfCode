#!/bin/bash

main() {
    redCubes=0
    greenCubes=0
    blueCubes=0
    sum=0

    input="/home/mawari/Desktop/git-projects/private/AdventOfCode/2023/2/input"

    while IFS= read -r line; do
        mapfile -t currentGames < <( echo "${line}" | tr ":" "\n" | tr ";" "\n")

        for game in "${currentGames[@]}"; do
            if [[ "${game}" == *"Game"* ]]; then
                temp=$(echo "${game}" | grep -oE '[0-9]+');
                continue
            fi
            currentRedCubes=$(echo "${game}" | grep -oE '[0-9]+ red' | grep -oE '[0-9]+')
            currentGreenCubes=$(echo "${game}" | grep -oE '[0-9]+ green' | grep -oE '[0-9]+')
            currentBlueCubes=$(echo "${game}" | grep -oE '[0-9]+ blue' | grep -oE '[0-9]+')

            if [[ "${currentRedCubes}" -gt "${redCubes}" ]]; then
                redCubes="${currentRedCubes}"
            fi
            
            if [[ "${currentGreenCubes}" -gt "${greenCubes}" ]]; then
                greenCubes="${currentGreenCubes}"
            fi
            
            if [[ "${currentBlueCubes}" -gt "${blueCubes}" ]]; then
                blueCubes="${currentBlueCubes}"
            fi
        done

        power=$(("${redCubes}"*"${greenCubes}"*"${blueCubes}"))
        sum=$(("${power}"+"${sum}"))

        redCubes=0
        greenCubes=0
        blueCubes=0

    done < "$input"

    echo "${sum}"
}

main
