# chessimg
[![GoDoc](https://godoc.org/github.com/loganjspears/chessimg?status.svg)](https://godoc.org/github.com/loganjspears/chessimg)
[![Build Status](https://drone.io/github.com/loganjspears/chessimg/status.png)](https://drone.io/github.com/loganjspears/chessimg/latest)
[![Coverage Status](https://coveralls.io/repos/loganjspears/chessimg/badge.svg?branch=master&service=github)](https://coveralls.io/github/loganjspears/chessimg?branch=master)

chessimg is a go library that creates a board image from FEN notation.

## Usage

```go
// populate buffer w/ SVG of the starting position
buf := bytes.NewBuffer([]byte{}) 
fenStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
if err := chessimg.New(buf).EncodeSVG(fenStr); err != nil {
	log.Fatal(err)
}
```