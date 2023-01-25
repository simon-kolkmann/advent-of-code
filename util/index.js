import fs from 'fs'
import { EOL } from 'os'
import { dirname } from 'path'
import path from 'path'
import { fileURLToPath } from 'url'

export const __dirname = meta => dirname(fileURLToPath(meta.url))

export const getPuzzleInput = meta => {
	const input = fs.readFileSync(path.join(__dirname(meta), './input.txt'), 'utf-8')
	return input.split(EOL)
}

export const getPuzzleInputWithoutEmptyLastLine = meta => {
	const input = getPuzzleInput(meta)

	if (input[input.length - 1] === '') {
		input.pop()
	}

	return input
}
