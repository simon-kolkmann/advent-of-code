import { getPuzzleInputWithoutEmptyLastLine } from '#root/util/index.js'

const input = getPuzzleInputWithoutEmptyLastLine(import.meta).map(rowOfTrees => rowOfTrees.split('').map(Number))

const getColumn = x => input.map(row => row[x])

class Tree {
	constructor({ x, y, height, up, down, left, right }) {
		this.x = x
		this.y = y
		this.height = height
		this.up = up
		this.down = down
		this.left = left
		this.right = right
	}

	get isVisible() {
		const containsLargerTree = direction => Math.max(...direction) >= this.height

		return !(
			containsLargerTree(this.up) &&
			containsLargerTree(this.down) &&
			containsLargerTree(this.left) &&
			containsLargerTree(this.right)
		)
	}

	get scenicScore() {
		const calculateVisibleTrees = trees => {
			for (let i = 0; i < trees.length; i++) {
				const tree = trees[i]

				if (tree >= this.height) {
					return i + 1
				}
			}

			return trees.length
		}

		const left = calculateVisibleTrees(this.left.reverse())
		const right = calculateVisibleTrees(this.right)
		const up = calculateVisibleTrees(this.up.reverse())
		const down = calculateVisibleTrees(this.down)

		return left * right * up * down
	}
}

const forest = input.map((rowOfTrees, y) => {
	return rowOfTrees.map((height, x) => {
		const left = rowOfTrees.slice(0, x)
		const right = rowOfTrees.slice(x + 1, rowOfTrees.length)
		const up = getColumn(x).slice(0, y)
		const down = getColumn(x).slice(y + 1, input.length)

		return new Tree({ x, y, height, left, right, up, down })
	})
})

const solutionPartOne = forest
	.flatMap(rowOfTrees => rowOfTrees.map(tree => tree.isVisible))
	.reduce((counter, visible) => (visible ? counter + 1 : counter), 0)

console.log(solutionPartOne)

// Part II

const scores = forest.flatMap(rowOfTrees => rowOfTrees).map(tree => tree.scenicScore)
const solutionPartTwo = Math.max(...scores)

console.log(solutionPartTwo)
