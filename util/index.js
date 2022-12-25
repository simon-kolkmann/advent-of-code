import { dirname } from 'path'
import { fileURLToPath } from 'url'

export const __dirname = meta => dirname(fileURLToPath(meta.url))
