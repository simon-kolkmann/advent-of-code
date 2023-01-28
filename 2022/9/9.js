import { getPuzzleInputWithoutEmptyLastLine } from '#root/util/index.js'

const input = getPuzzleInputWithoutEmptyLastLine(import.meta)

const head = { x: 0, y: 0 }
const tail = {
	x: 0,
	y: 0,
	history: [],
	move({ x, y }) {
		this.history.push({ x: this.x, y: this.y })
		this.x = x || this.x
		this.y = y || this.y
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

		console.log(distance)

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

const points = [...tail.history, { x: tail.x, y: tail.y }]

const maxX = points.reduce((x, point) => (point.x > x ? point.x : x), 0)
const minX = points.reduce((x, point) => (point.x < x ? point.x : x), 0)
const maxY = points.reduce((y, point) => (point.y > y ? point.y : y), 0)
const minY = points.reduce((y, point) => (point.y < y ? point.y : y), 0)

for (let y = maxY; y >= minY; y--) {
	let row = ''
	for (let x = minX; x <= maxX + 1; x++) {
		if (points.find(point => point.x === x && point.y === y)) {
			row += `#`
		} else {
			row += '.'
		}
	}

	console.log(row)
}

points.map(point => `${point.x},${point.y}`).forEach(point => console.log(point))

const distinct = new Set(points.map(point => `${point.x},${point.y}`))

console.log(distinct.size)
