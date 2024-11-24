<script setup lang="ts">

import type { NumberMatch } from '@/types/number-match';
import { ref } from 'vue';

const calibrationInput =
`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`;

let inputText = ref(calibrationInput);
let solution1 = ref();
let solution2 = ref();

// Solve for part 1 - sum of all numbers with adjacent symbols
function solvePart1() {
    const inputLines = inputText.value.split('\n');
    let sum = 0;

    const allNumberMatches = findAllNumbers(inputLines);

    // go through each number position and check for any bordering symbols (if none found then add to the sum)
    for(let potentialPart of allNumberMatches) {
        if (isPart(potentialPart, inputLines))
            sum += potentialPart.number;
    }

    solution1.value = sum;
}

// Solve for part 2 - now we only want to find * symbols specifically (cogs) and only ones adjacent to exactly 2 numbers. Then multiply those 2 numbers and sum all of the products.
function solvePart2() {
    const inputLines = inputText.value.split('\n');
    let sum = 0;

    let numberMatchesPerLine: NumberMatch[][] = [];
    // for efficiency we need the matched numbers to be indexed by the line they were on in this case. For convenience we'll reuse the eixsting function to find them
    for(let y = 0; y < inputLines.length; y++) {
        numberMatchesPerLine[y] = findAllNumbers([inputLines[y]], y);
    }
    
    const allCogPositions = findAllCogPositions(inputLines);
 
    for (let [x, y ] of allCogPositions) {
        // adjacent numbers can only be on the line above, the same line, or the line below - so look for adjacent numbers on those lines only
        let relevantLineMatches = numberMatchesPerLine.slice(Math.max(y - 1, 0), Math.min(y + 2, inputLines.length + 1)).flat();
        let adjacentNumbers: number[] = findAnyAdjacentNumbers(x, y, relevantLineMatches);
        if (adjacentNumbers.length === 2) {
            sum += adjacentNumbers[0] * adjacentNumbers[1];
        }
    }

    solution2.value = sum;
}

// find x,y positions of numbers. Return all matches. A NumberMatch also contains the actual number and the length in charatcers of that number.
function findAllNumbers(inputLines: string[], offset: number = 0): NumberMatch[] {
    let allNumberMatches: NumberMatch[] = [];

    for (let i = 0; i<inputLines.length; i++) {
        let numberMatches = [...inputLines[i].matchAll(/\d+/g)];
        for (let match of numberMatches) {
            const numberMatch: NumberMatch = {
                number: +match[0],
                x: match.index,
                y: offset + i,
                length: match[0].length
            }
            allNumberMatches.push(numberMatch);
        }
    }

    return allNumberMatches;
}

// find x,y positions of asterisks (cogs)
function findAllCogPositions(inputLines: string[]): [number, number][] {
    let allNumberMatches: [number, number][] = [];

    for (let i = 0; i<inputLines.length; i++) {
        let numberMatches = [...inputLines[i].matchAll(/\*/g)];
        for (let match of numberMatches) {
            allNumberMatches.push([match.index, i]);
        }
    }

    return allNumberMatches;
}

// for a given cog position return all adjacent numbers in the set provided
function findAnyAdjacentNumbers(cogX: number, cogY: number, numberMatches: NumberMatch[]): number[] {
    let adjacentNumbers: number[] = [];
    
    for(let match of numberMatches) {
        if ((match.y === cogY - 1  || match.y === cogY + 1) && cogX >= match.x - 1 && cogX <= match.x + match.length ) {
            adjacentNumbers.push(match.number);
        }

        // on the same line as the cog "adjacent" literally means it needs to be next to it
        if (match.y === cogY && (match.x + match.length === cogX /* left of cog */ || match.x === cogX + 1 /* right of cog */ ) ) {
            adjacentNumbers.push(match.number);
        }
    }
    return adjacentNumbers;
}

// following the definition from the puzzle - return true if any symbol is adjacent to the potentialPart (number from the input)
function isPart(potentialPart: NumberMatch, inputLines: string[]): boolean {
    // get line length... assuming the input has at least one line and that lines are all the same length
    const lineLength = inputLines[0].length;
    const startX = Math.max(potentialPart.x -1, 0);
    const endX = Math.min(potentialPart.x + potentialPart.length, lineLength);

    // check above for symbols
    if (potentialPart.y > 0 && /[^\d\\.\n]/g.test(inputLines[potentialPart.y - 1].substring(startX, endX + 1))) {
        return true;
    }
    //check left
    if (potentialPart.x > 0 && /[^\d\\.\n]/g.test(inputLines[potentialPart.y].charAt(potentialPart.x - 1))) {
        return true;
    }
    // check right
    if (potentialPart.x + potentialPart.length < lineLength && /[^\d\\.\n]/g.test(inputLines[potentialPart.y].charAt(potentialPart.x + potentialPart.length ))) {
        return true;
    }
    //check below
    if (potentialPart.y < inputLines.length -1 && /[^\d\\.\n]/g.test(inputLines[potentialPart.y + 1].substring(startX, endX + 1))) {
        return true;
    }

    return false;
}

function clearSolution() {
    solution1.value = undefined;
    solution2.value = undefined;
}

</script>

<template>
  <v-container fluid>
    <v-row>
      <v-textarea
      label="Input text"
      v-model="inputText"
      name="input-text"
      variant="filled"
      auto-grow
      class="fixed-width-font"
      max-rows="16"
      @change="clearSolution"
      ></v-textarea>
    </v-row>
    <v-row>
      <v-col>
        <v-btn prepend-icon="mdi-check-circle" class="solve1-btn" @click="solvePart1">
          Solve Part 1
        </v-btn>
        <v-chip v-if="solution1">Answer is: {{solution1}}</v-chip>
      </v-col>
      <v-col>
        <v-btn prepend-icon="mdi-check-circle" class="solve2-btn" @click="solvePart2">
          Solve Part 2
        </v-btn>
        <v-chip v-if="solution2">Answer is: {{solution2}}</v-chip>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
.solve-btn {
    background-color: aquamarine;
}

.fixed-width-font {
    font-family: 'Courier New', Courier, monospace;
}
</style>
