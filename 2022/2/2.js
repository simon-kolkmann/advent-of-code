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

	static getWinningShapeAgainst(shape) {
		if (shape instanceof Rock) {
			return new Paper()
		} else if (shape instanceof Paper) {
			return new Scissor()
		} else {
			return new Rock()
		}
	}

	static getLosingShapeAgainst(shape) {
		if (shape instanceof Rock) {
			return new Scissor()
		} else if (shape instanceof Paper) {
			return new Rock()
		} else {
			return new Paper()
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

for (const [theirsAsChar, outcome] of rounds) {
	const theirs = Game.getShapeByChar(theirsAsChar)

	if (outcome === 'X') {
		total += Game.getLosingShapeAgainst(theirs).compare(theirs)
	} else if (outcome === 'Y') {
		total += theirs.worth + 3
	} else {
		total += Game.getWinningShapeAgainst(theirs).compare(theirs)
	}
}

console.log(`Total score: ${total}`)
