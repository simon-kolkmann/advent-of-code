import { getPuzzleInput } from '#root/util/index.js'

const assignments = ['2-4,6-8', '2-3,4-5', '5-7,7-9', '2-8,3-7', '6-6,4-6', '2-6,4-8']

for (const assignment of assignments) {
	const [one, two] = assignment.split(',')

	const [a, b] = one.split('-').map(Number)
	const [c, d] = two.split('-').map(Number)

	console.log(a, b, c, d)
}
