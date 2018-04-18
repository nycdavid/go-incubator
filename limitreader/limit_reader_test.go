package limitreader

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestLimitReader(t *testing.T) {
	fileRdr, err := os.Open("./scratch.txt")
	if err != nil {
		log.Print(err)
	}
	rdr := LimitReader(fileRdr, 4)
	expected := 4
	data := make([]byte, 6)
	got, err := rdr.Read(data)

	if string(data) != "Lore" {
		errMsg := fmt.Sprintf("Expected string to be %s, got %s", "Lore", string(data))
		t.Error(errMsg)
	}

	if got != expected {
		errMsg := fmt.Sprintf("Expected read to limit to %d, but read %d", expected, got)
		t.Error(errMsg)
	}
}
