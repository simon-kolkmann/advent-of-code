import assert from 'assert'

import { getPuzzleInputWithoutEmptyLastLine } from '#root/util/index.js'

const input = getPuzzleInputWithoutEmptyLastLine(import.meta)

const head = { x: 0, y: 0 }
const tail = {
	x: 0,
	y: 0,
	history: [],
	move({ x, y }) {
		this.history.push({ x: this.x, y: this.y })
		this.x = x !== undefined ? x : this.x
		this.y = y !== undefined ? y : this.y
	}
}

const move = (direction, steps) => {
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

		if ([2, -2].includes(distance.x)) {
			if (distance.y === 0) {
				tail.move({ x: tail.x + distance.x / 2 })
			} else {
				tail.move({ x: tail.x + distance.x / 2, y: tail.y + distance.y })
			}
		}

		if ([2, -2].includes(distance.y)) {
			if (distance.x === 0) {
				tail.move({ y: tail.y + distance.y / 2 })
			} else {
				tail.move({ y: tail.y + distance.y / 2, x: tail.x + distance.x })
			}
		}

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

	move(direction, steps)
}

const getVisitedPoints = tail => {
	return [...tail.history, { x: tail.x, y: tail.y }]
}

function paint(points) {
	const { min, max } = points.reduce(
		({ min, max }, point) => {
			min.x = point.x < min.x ? point.x : min.x
			max.x = point.x > max.x ? point.x : max.x
			min.y = point.y < min.y ? point.y : min.y
			max.y = point.y > max.y ? point.y : max.y

			return { min, max }
		},
		{
			min: { x: Infinity, y: Infinity },
			max: { x: -Infinity, y: -Infinity }
		}
	)

	let grid = ''

	for (let y = max.x; y >= min.y; y--) {
		for (let x = min.x; x <= max.x; x++) {
			if (head.x === x && head.y === y) {
				grid += 'H'
			} else if (tail.x === x && tail.y === y) {
				grid += 'T'
			} else if (points.find(point => point.x === x && point.y === y)) {
				grid += `#`
			} else {
				grid += '.'
			}
		}

		grid += '\n'
	}

	console.log(grid)
}

const points = getVisitedPoints(tail)

// paint(points)

const distinct = new Set(points.map(point => `${point.x},${point.y}`))

console.log(distinct.size)

assert(distinct.size === 5513, `The solution must be 5513 but was ${distinct.size}`)

console.log('OK')
