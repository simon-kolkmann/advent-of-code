import { getPuzzleInputWithoutEmptyLastLine } from '#root/util/index.js'

const input = getPuzzleInputWithoutEmptyLastLine(import.meta)

const CRT = {
	pixels: '',
	add(pixel) {
		this.pixels += pixel
	},
	print() {
		for (let i = 0; i < this.pixels.length; i += 40) {
			console.log(this.pixels.substring(i, i + 40))
		}
	}
}

const CPU = {
	cycle: 0,
	register: 1,
	get sprite() {
		return [this.register - 1, this.register, this.register + 1]
	},
	interestingSignalStrengths: [],
	tick() {
		if (this.sprite.includes(this.cycle % 40)) {
			CRT.add('#')
		} else {
			CRT.add('.')
		}

		this.cycle++

		if (this.cycle === 20 || (this.cycle + 20) % 40 === 0) {
			const signalStrength = this.cycle * this.register
			this.interestingSignalStrengths.push(signalStrength)
		}
	},
	addToRegister(n) {
		this.register += n
	},
	run(cmd, arg) {
		if (cmd === 'addx') {
			AddX.execute(this, arg)
		}

		if (cmd === 'noop') {
			Noop.execute(this)
		}
	}
}

const AddX = {
	execute(cpu, arg) {
		cpu.tick()
		cpu.tick()
		cpu.addToRegister(arg)
	}
}

const Noop = {
	execute(cpu) {
		cpu.tick()
	}
}

for (const instruction of input) {
	const [cmd, arg] = instruction.split(' ')
	CPU.run(cmd, Number(arg))
}

const solutionPartOne = CPU.interestingSignalStrengths.reduce((sum, current) => sum + current, 0)

console.log(solutionPartOne)

// Part II
console.log('')
CRT.print()
