package cities

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ReadCities load cities from file filename
func ReadCities(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		msg := fmt.Sprintf("Error while reading cities<%v>\n", err.Error())
		panic(msg)
	}
	datastr := string(data)
	fmt.Printf("#%v#", datastr)
}

// ReadCities2 load cities from file filename using Reader
func ReadCities2(filename string) {
	f, _ := os.Open(filename)
	defer f.Close() // appel en fin de fonction
	reader := bufio.NewReader(f)
	for sizeRead := -1; sizeRead != 0; {
		var r rune
		r, sizeRead, _ = reader.ReadRune()
		fmt.Printf("Read: %v (%v, %v)\n", string(r), r, sizeRead)
	}
}

// ReadCities3 load cities from file filename using buffered Reader
func ReadCities3(filename string) {
	f, _ := os.Open(filename)
	defer f.Close() // appel en fin de fonction
	reader := bufio.NewReader(f)
	for line := "U"; line != ""; {
		line, _ = reader.ReadString(10)
		line = strings.TrimSpace(line)
		words := strings.Split(line, ",")
		if len(line) >= 3 { // line with at least 3 characters
			fmt.Printf("Read: <%v> <%v> => %v\n", line, line[:3], words)
		}
	}
}

// ReadCities4 load cities from file filename using csv Reader
func ReadCities4(filename string) []City {
	var res []City
	f, _ := os.Open(filename)
	defer f.Close() // appel en fin de fonction
	r := csv.NewReader(f)
	for row := []string{"dummy"}; row != nil; {
		row, _ = r.Read()
		if len(row) > 0 {
			c := City{Name: row[0], Zipcode: row[1]}
			res = append(res, c)
			// fmt.Printf("Read: %v\n", c)
		}
	}
	return res
}
