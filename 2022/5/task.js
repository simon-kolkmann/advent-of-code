import fs from 'fs'
import { EOL } from 'os'
import path from 'path'

import { __dirname } from '#root/util/index.js'

const input = fs.readFileSync(path.join(__dirname(import.meta), './input.txt'), 'utf-8')
const lines = input.split(EOL)

const divider = lines.indexOf('')

const matrix = []
const moves = []

for (let line = 0; line < lines.length - 1; line++) {
	if (line === divider) {
		continue
	}

	if (line < divider) {
		matrix.push(lines[line])
	} else {
		moves.push(lines[line])
	}
}

const numbers = matrix[matrix.length - 1]

const stacks = {}

for (let index = 0; index < numbers.length; index++) {
	const char = numbers[index]

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
	// eslint-disable-next-line prefer-const
	let [count, source, dest] = line.replace('move ', '').replace('from ', '').replace('to ', '').split(' ')

	while (count > 0) {
		const element = stacks[source].pop()
		stacks[dest].push(element)
		count--
	}
}

const solution = Object.entries(stacks).reduce((solution, [_, stack]) => {
	console.log(`Last element of stack ${_}: ${stack[stack.length - 1]}`)
	return solution + stack[stack.length - 1]
}, '')

console.log(`Solution: ${solution}`)
