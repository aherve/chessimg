package chessimg_test

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"image/color"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/notnil/chess"
	"github.com/notnil/chessimg"
)

const expectedMD5 = "a76e587e930921260e71b6bb3e75e441"
const expectedReversedMD5 = "8b61a395f4e123d76582171064827de9"

func TestSVG(t *testing.T) {
	// create buffer of actual svg
	buf := bytes.NewBuffer([]byte{})
	reversedBuf := bytes.NewBuffer([]byte{})
	fenStr := "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq - 0 1"
	pos := &chess.Position{}
	if err := pos.UnmarshalText([]byte(fenStr)); err != nil {
		t.Error(err)
	}
	mark := chessimg.MarkSquares(color.RGBA{255, 255, 0, 1}, chess.D2, chess.D4)
	if err := chessimg.SVG(buf, pos.Board(), mark); err != nil {
		t.Error(err)
	}

	if reversedErr := chessimg.ReversedSVG(reversedBuf, pos.Board(), mark); reversedErr != nil {
		t.Error(reversedErr)
	}

	// compare to expected svg
	actualSVG := strings.TrimSpace(buf.String())
	actualMD5 := fmt.Sprintf("%x", md5.Sum([]byte(actualSVG)))
	if actualMD5 != expectedMD5 {
		t.Errorf("expected actual md5 hash to be %s but got %s", expectedMD5, actualMD5)
	}

	// compare to expected reversed svg
	reversedSVG := strings.TrimSpace(reversedBuf.String())
	reversedMD5 := fmt.Sprintf("%x", md5.Sum([]byte(reversedSVG)))
	if reversedMD5 != expectedReversedMD5 {
		t.Errorf("expected reversed md5 hash to be %s but got %s", expectedReversedMD5, reversedMD5)
	}

	// create actual svg file for visualization
	f, err := os.Create("example.svg")
	defer f.Close()
	if err != nil {
		t.Error(err)
	}
	if _, err := io.Copy(f, bytes.NewBufferString(actualSVG)); err != nil {
		t.Error(err)
	}

	f, err = os.Create("example_reversed.svg")
	defer f.Close()
	if err != nil {
		t.Error(err)
	}
	if _, err := io.Copy(f, bytes.NewBufferString(reversedSVG)); err != nil {
		t.Error(err)
	}

}
