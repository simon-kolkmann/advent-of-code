import { getPuzzleInput } from '#root/util/index.js'

const stream = getPuzzleInput(import.meta)[0]

for (let i = 3; i < stream.length; i++) {
	const possibleStartOfPackageMarker = stream.slice(i - 3, i + 1)

	const set = new Set(possibleStartOfPackageMarker)

	if (set.size === 4) {
		console.log(i + 1)
		break
	}
}
