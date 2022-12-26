import { getPuzzleInput } from '#root/util/index.js'

const rucksacks = getPuzzleInput(import.meta)

const getPriorityOfChar = char => {
	const priorities = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
	return priorities.indexOf(char) + 1
}

let sum = 0

for (let i = 0; i < rucksacks.length; i += 3) {
	const one = rucksacks[i]
	const two = rucksacks[i + 1]
	const three = rucksacks[i + 2]

	for (const item of one) {
		if (two.includes(item) && three.includes(item)) {
			sum += getPriorityOfChar(item)
			break
		}
	}
}

console.log(sum)
