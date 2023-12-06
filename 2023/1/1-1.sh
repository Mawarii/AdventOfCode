#!/bin/bash

main() {
    sum=0

    input="/home/mawari/Desktop/git-projects/private/AdventOfCode/2023/1/input"

    while IFS= read -r line; do
        number=$(printf "${line}" | grep -oE '[0-9]' | head -1 | tr -d '\n'; printf "${line}" | grep -oE '[0-9]' | tail -1)
        sum=$(("${number}"+"${sum}"))
    done < "$input"

    echo "${sum}"
}

main
