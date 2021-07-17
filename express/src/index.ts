import express, { Request, Response } from 'express'
import morgan from 'morgan'

const app = express()

app.use(morgan('combined'))

app.get('/', (req: Request, res: Response) => {
  res.send('ok')
})

app.listen(3000)
console.log('Listening at localhost:3000')
