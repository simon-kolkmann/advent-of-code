import fs from 'fs'

const input = fs.readFileSync('./input', 'utf-8')
const lines = input.split('\r\n')

const divider = lines.indexOf('')

let matrix = []
const moves = []

for (let line = 0; line < lines.length - 1; line++) {
    if (line === divider) 
        continue

    if (line < divider) {
        matrix.push(lines[line])
    } else {
        moves.push(lines[line])
    }
}

const numbers = matrix[matrix.length - 1]
const stacks = {}

for (let index = 0; index < numbers.length; index++) {
    const char = numbers[index];

    if (char !== ' ') {
        stacks[char] = index
    }
}

for (const [a, b] of Object.entries(stacks)) {
    stacks[a] = []
    matrix.forEach(line => {
        if (line[b] !== ' ') {
            stacks[a].unshift(line[b])
        }
    })
}

for (const line of lines) {
    let [count, source, dest] = line.replace('move ', '').replace('from ', '').replace('to ', '').split(' ')

    while(count > 0) {
        const element = stacks[source].pop();
        stacks[dest].push(element)
        count--
    }
}

for (const [stack, elements] of Object.entries(stacks)) {
    console.log('Last element of stack ' + stack  + ': ' + elements[elements.length - 1])
}
