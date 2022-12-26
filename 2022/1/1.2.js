import { getPuzzleInput } from '#root/util/index.js'

const lines = getPuzzleInput(import.meta)
const calories = []

let currentCalories = 0

lines.forEach((line, index, array) => {
	if (line === '') {
		return
	}

	currentCalories += Number(line)

	if (array[index + 1] === '') {
		calories.push(currentCalories)
		currentCalories = 0
	}
})

calories.sort((a, b) => b - a)

const [one, two, three] = calories

console.log(one + two + three)
