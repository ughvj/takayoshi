package dml

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"errors"
)

type LoadedDML struct {
	sql *string
	debug bool
}

func Load(title string) (*LoadedDML, error) {
	loadedDML := &LoadedDML{}

	fp, err := os.Open("./dml/" + title + ".sql")
	if err != nil {
		fmt.Println("ERROR: file cannot be opened")
		return nil, err
	}
	defer fp.Close()

	data := make([]byte, 2048)
	count, err := fp.Read(data)
	if err != nil {
    	fmt.Println(err)
    	fmt.Println("fail to read file")
		return nil, err
	}

	sql := string(data[:count])
	loadedDML.sql = &sql
	return loadedDML, nil
}

func (l *LoadedDML) Debug() *LoadedDML {
	l.debug = true
	return l
}

func (l *LoadedDML) off() {
	l.debug = false
}

func (l *LoadedDML) GetSQL() string {
	defer l.off()
	if l.debug {
		fmt.Println(*l.sql)
	}
	return *l.sql
}

// It has vulnerability of sql-injection
func (l *LoadedDML) FillPlaceholders(placeholders ...string) error {
	defer l.off()

	if l.sql == nil {
		return errors.New("missing sql.")
	}
	if len(placeholders) != l.countPlaceholders() {
		return errors.New("len(args) != number of placeholder, pls chk target sql.")
	}

	if l.debug {
		fmt.Println("before filling: " + *l.sql)
	}
	for i, placeholder := range placeholders {
		strings.Replace(*l.sql, "$" + strconv.Itoa(i+1), placeholder, 1)
	}
	if l.debug {
		fmt.Println("after filling: " + *l.sql)
	}

	return nil
}

func (l *LoadedDML) countPlaceholders() (count int) {
	for {
		if strings.Contains(*l.sql, "$" + strconv.Itoa(count+1)) {
			count++
		} else {
			return count
		}
	}
}
