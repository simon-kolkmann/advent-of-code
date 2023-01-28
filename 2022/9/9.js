import { getPuzzleInputWithoutEmptyLastLine } from '#root/util/index.js'

const input = getPuzzleInputWithoutEmptyLastLine(import.meta)

const head = { x: 0, y: 0 }
const tail = { x: 0, y: 0 }

const move = (headOrTail, direction, steps) => {
	while (steps > 0) {
		switch (direction) {
			case 'U':
				head.y++
				break

			case 'D':
				head.y--
				break

			case 'R':
				head.x++
				break

			case 'L':
				head.x--
				break
		}

		/**
		 * Check if the distance between head and tail is greater than 1.
		 * If it is, move the tail accordingly.
		 */
		const distance = getDistanceBetween(head, tail)

		if (distance.x === 2) {
			tail.x = tail.x + distance.x - 1
		}

		if (distance.x > 2 || distance.y > 2) {
			// TODO: Move tail
		}

		console.log(head)
		console.log(tail)

		steps--
	}
}

const getDistanceBetween = (a, b) => {
	return {
		x: a.x - b.x,
		y: a.y - b.y
	}
}

for (const line of input) {
	const [direction, steps] = line.split(' ').map((value, i) => {
		return i === 0 ? value : Number(value)
	})

	move(head, direction, steps)
}
