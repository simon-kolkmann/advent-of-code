import { getPuzzleInput } from '#root/util/index.js'

const assignments = getPuzzleInput(import.meta).filter(line => line !== '')

class Range {
	constructor(a, b) {
		this.a = a
		this.b = b
	}

	overlaps(range) {
		return this.a <= range.a && this.b >= range.a
	}
}

const count = assignments.reduce((count, assignment) => {
	const [one, two] = assignment.split(',')

	const [a, b] = one.split('-').map(Number)
	const [c, d] = two.split('-').map(Number)

	if (new Range(a, b).overlaps(new Range(c, d))) {
		return count + 1
	} else if (new Range(c, d).overlaps(new Range(a, b))) {
		return count + 1
	} else {
		return count
	}
}, 0)

console.log(count)
