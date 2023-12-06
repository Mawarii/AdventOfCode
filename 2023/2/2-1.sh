#!/bin/bash

main() {
    redCubes=12
    greenCubes=13
    blueCubes=14
    sum=0
    add=true

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

            if [[ "${currentRedCubes}" -gt "${redCubes}" ]] || [[ "${currentGreenCubes}" -gt "${greenCubes}" ]] || [[ "${currentBlueCubes}" -gt "${blueCubes}" ]]; then
                add=false
                break
            fi
        done
        if "${add}"; then
            sum=$((sum + temp))
        else
            add=true
        fi

    done < "$input"

    echo "${sum}"
}

main
