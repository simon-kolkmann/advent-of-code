import fs from 'fs'
import { EOL } from 'os'
import path from 'path'

import { __dirname } from '#root/util/index.js'

const input = fs.readFileSync(path.join(__dirname(import.meta), './input.txt'), 'utf-8')
const lines = input.split(EOL)
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
