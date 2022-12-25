import fs from 'fs'
import { EOL } from 'os'
import path from 'path'

import { __dirname } from '#root/util/index.js'

const input = fs.readFileSync(path.join(__dirname(import.meta), './input.txt'), 'utf-8')
const lines = input.split(EOL)

const divider = lines.indexOf('')

const initialStateMatrix = []
const moves = []

lines.forEach((line, index) => {
	// Ignore divider & final newline
	if (!line.length > 0) {
		return
	}

	if (index < divider) {
		initialStateMatrix.push(line)
	} else {
		moves.push(line)
	}
})

const numbers = initialStateMatrix[initialStateMatrix.length - 1]

const stacks = {}

for (let index = 0; index < numbers.length; index++) {
	const char = numbers[index]

	if (char !== ' ') {
		stacks[char] = index
	}
}

for (const [a, b] of Object.entries(stacks)) {
	stacks[a] = []
	initialStateMatrix.forEach(line => {
		if (line[b] !== ' ') {
			stacks[a].unshift(line[b])
		}
	})
}

for (const move of moves) {
	// eslint-disable-next-line prefer-const
	let [[count], [source], [dest]] = move.matchAll(/(\d+)/g)

	while (count > 0) {
		stacks[dest].push(stacks[source].pop())
		count--
	}
}

const solution = Object.entries(stacks).reduce((solution, [_, stack]) => {
	console.log(`Last element of stack ${_}: ${stack[stack.length - 1]}`)
	return solution + stack[stack.length - 1]
}, '')

console.log(`Solution: ${solution}`)
