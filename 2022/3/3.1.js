import { getPuzzleInput } from '#root/util/index.js'

const rucksacks = getPuzzleInput(import.meta).filter(line => line !== '')

const getPriorityOfChar = char => {
	const priorities = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
	return priorities.indexOf(char) + 1
}

const sum = rucksacks.reduce((sum, rucksack) => {
	const width = rucksack.length / 2

	const one = rucksack.slice(0, width)
	const two = rucksack.slice(width, rucksack.length)

	for (const char of one) {
		if (two.includes(char)) {
			return sum + getPriorityOfChar(char)
		}
	}
}, 0)

console.log(sum)
