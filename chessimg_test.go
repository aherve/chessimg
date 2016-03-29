package chessimg_test

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/loganjspears/chessimg"
)

const (
	expectedMD5 = "ba032bdf69ead7837d9ed7fbbc81511a"
)

func TestSVG(t *testing.T) {
	// create buffer of actual svg
	buf := bytes.NewBuffer([]byte{})
	fenStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	if err := chessimg.New(buf).EncodeSVG(fenStr); err != nil {
		t.Error(err)
	}

	// compare to expected svg
	actualSVG := strings.TrimSpace(buf.String())
	actualMD5 := fmt.Sprintf("%x", md5.Sum([]byte(actualSVG)))
	if actualMD5 != expectedMD5 {
		t.Errorf("expected actual md5 hash to be %x but got %x", expectedMD5, actualMD5)
	}

	// create actual svg file for visualization
	f, err := os.Create("actual.svg")
	defer f.Close()
	if err != nil {
		t.Error(err)
	}
	if _, err := io.Copy(f, bytes.NewBufferString(actualSVG)); err != nil {
		t.Error(err)
	}
}
