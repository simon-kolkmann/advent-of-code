import { getPuzzleInputWithoutEmptyLastLine } from '#root/util/index.js'

const input = getPuzzleInputWithoutEmptyLastLine(import.meta)

const monkeys = []

class Monkey {
	constructor(id, items, operation, test) {
		this.id = id
		this.items = items
		this.operation = operation
		this.test = test
		this.inspections = 0
	}

	inspect(item) {
		this.inspections++

		// 1. Increase worry level by operation
		const evaluable = this.operation.replaceAll('old', item)

		// eslint-disable-next-line no-eval
		item = eval(evaluable)

		// 2. Decrease worry level (Part I)
		// item = Math.floor(item / 3)

		/**
		 * 2. Decrease worry level (Part II)
		 * Disclaimer: I only got this working thanks to https://www.reddit.com/r/adventofcode/comments/zizi43/comment/iztt8mx
		 */
		const productOfDivisors = monkeys.map(monkey => monkey.test.divisor).reduce((a, b) => a * b)
		item = item % productOfDivisors

		// 3. Test new worry level for condition, return target monkey
		if (item % this.test.divisor === 0) {
			return { item, target: this.test.onPass }
		} else {
			return { item, target: this.test.onFailure }
		}
	}
}

/**
 * Parse the input and create an array of monkeys.
 */
for (let i = 0; i < input.length; i += 7) {
	// "Monkey X"
	const id = Number(input[i][7])

	// "  Starting items: X, Y, Z"
	const items = input[i + 1]
		.substring(18, input[i + 1].length)
		.split(', ')
		.map(Number)

	// "  Operation: new = old <operator> <number|old>"
	const operation = input[i + 2].substring(19, input[i + 2].length)

	// "  Test: divisible by X"
	const divisor = Number(input[i + 3].substring(21, input[i + 3].length))

	// "    If true: throw to monkey X"
	const onPass = Number(input[i + 4].substring(29, input[i + 4].length))

	// "    If false: throw to monkey X"
	const onFailure = Number(input[i + 5].substring(30, input[i + 5].length))

	monkeys.push(new Monkey(id, items, operation, { divisor, onPass, onFailure }))
}

/**
 * Run the actual game.
 */

// Part I
// const rounds = 20

// Part II
const rounds = 10000

for (let round = 1; round <= rounds; round++) {
	for (const monkey of monkeys) {
		for (const current of monkey.items) {
			const { item, target } = monkey.inspect(current)
			monkeys.find(monkey => monkey.id === target).items.push(item)
		}

		monkey.items = []
	}
}

const [one, two] = monkeys.map(monkey => monkey.inspections).sort((a, b) => b - a)

const solutionPartOne = one * two

console.log(solutionPartOne)
