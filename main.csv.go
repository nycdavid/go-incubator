package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("./seeds.csv")
	rdr := bufio.NewReader(file)
	cr := csv.NewReader(rdr)
	slc, _ := cr.ReadAll()
	fmt.Println(slc)
}
