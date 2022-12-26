import { getPuzzleInput } from '#root/util/index.js'

const stream = getPuzzleInput(import.meta)[0]

for (let i = 13; i < stream.length; i++) {
	const possibleStartOfMessageMarker = stream.slice(i - 13, i + 1)

	const set = new Set(possibleStartOfMessageMarker)

	if (set.size === 14) {
		console.log(i + 1)
		break
	}
}
