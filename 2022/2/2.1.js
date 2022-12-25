import { getPuzzleInput } from '#root/util/index.js'

const input = getPuzzleInput(import.meta)

class Game {
	static getShapeByChar(char) {
		if (['A', 'X'].includes(char)) {
			return new Rock()
		} else if (['B', 'Y'].includes(char)) {
			return new Paper()
		} else {
			return new Scissor()
		}
	}
}

class Rock {
	constructor() {
		this.worth = 1
	}

	compare(shape) {
		if (shape instanceof Rock) {
			return 3 + this.worth
		} else if (shape instanceof Scissor) {
			return 6 + this.worth
		} else {
			return 0 + this.worth
		}
	}
}

class Paper {
	constructor() {
		this.worth = 2
	}

	compare(shape) {
		if (shape instanceof Paper) {
			return 3 + this.worth
		} else if (shape instanceof Rock) {
			return 6 + this.worth
		} else {
			return 0 + this.worth
		}
	}
}

class Scissor {
	constructor() {
		this.worth = 3
	}

	compare(shape) {
		if (shape instanceof Scissor) {
			return 3 + this.worth
		} else if (shape instanceof Paper) {
			return 6 + this.worth
		} else {
			return 0 + this.worth
		}
	}
}

const rounds = input.filter(line => line !== '').map(line => line.split(' '))

let total = 0

for (const [theirs, ours] of rounds) {
	total += Game.getShapeByChar(ours).compare(Game.getShapeByChar(theirs))
}

console.log(`Total score: ${total}`)
