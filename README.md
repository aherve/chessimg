# chessimg
[![GoDoc](https://godoc.org/github.com/notnil/chessimg?status.svg)](https://godoc.org/github.com/notnil/chessimg)
[![Build Status](https://drone.io/github.com/notnil/chessimg/status.png)](https://drone.io/github.com/notnil/chessimg/latest)
[![Coverage Status](https://coveralls.io/repos/github/notnil/chessimg/badge.svg?branch=master)](https://coveralls.io/github/notnil/chessimg?branch=master)
[![Go Report Card](http://goreportcard.com/badge/notnil/chessimg)](http://goreportcard.com/report/notnil/chessimg)

### Code Example

```go
// create file
f, err := os.Create("example.svg")
if err != nil {
    log.Fatal(err)
}
// write image of position and marked squares to file
fenStr := "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq - 0 1"
mark := chessimg.MarkSquares(color.RGBA{255, 255, 0, 1}, chess.D2, chess.D4)
if err := chessimg.New(f, mark).EncodeSVG(fenStr); err != nil {
	log.Fatal(err)
}
```

### Resulting Image

![rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq - 0 1](/example.png)
 
