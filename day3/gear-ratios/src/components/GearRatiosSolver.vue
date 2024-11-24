<script setup lang="ts">

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
let solution = ref();

// represents info about a potential "part" in the input grid - i.e. a number and its position
interface NumberMatch {
    number: number,
    x: number,
    y: number,
    length: number
}

function solve() {

    const inputLines = inputText.value.split('\n');

    let allNumberMatches: NumberMatch[] = [];
    let sum = 0;

    // find x,y positions of numbers. remember them
    for (let i = 0; i<inputLines.length; i++) {
        let numberMatches = [...inputLines[i].matchAll(/\d+/g)];
        for (let match of numberMatches) {
            const numberMatch: NumberMatch = {
                number: +match[0],
                x: match.index,
                y: i,
                length: match[0].length
            }
            allNumberMatches.push(numberMatch);
        }
    }

    // go through each number position and check for any bordering symbols (if none found then add to the sum)
    for(let potentialPart of allNumberMatches) {
        if (isPart(potentialPart, inputLines))
            sum += potentialPart.number;
    }

    solution.value = sum;
}

// following the definition form the puzzle - return true if any symbol is adjacent to the potentialPart (number from the input)
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
    solution.value = undefined;
}

</script>

<template>
    <v-container fluid>
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
    <v-btn prepend-icon="mdi-check-circle" class="solve-btn" @click="solve">
        Solve
    </v-btn>
    <v-chip v-if="solution">Answer is: {{solution}}</v-chip>
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
