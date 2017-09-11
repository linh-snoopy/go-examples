package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const PhoneNumber = 0
const ActivationDate = 1
const DeactivationDate = 2

type Date struct {
	Year  int
	Month int
	Day   int
}

type Stuff struct {
	start  Date
	finish Date
}

type Record struct {
	start  Date
	finish Date
	stuffs []Stuff
}

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "filename csv")
	flag.Parse()
	if filename == "" {
		flag.Usage()
		return
	}
	err := readFile(filename)
	if err != nil {
		fmt.Println(err)
	}
}

func readFile(filename string) error {
	// Load csv file
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Create a new reader
	m := readFileCsv(f)

	err = writeFile(m, filename)
	if err != nil {
		return err
	}

	return nil
}

func readFileCsv(f *os.File) map[string]Record {
	r := csv.NewReader(f)
	m := make(map[string]Record)
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			continue
		}

		if invalidRow(row) {
			continue
		}

		record, err := updateRecord(m, row)
		if err != nil {
			continue
		}
		m[row[PhoneNumber]] = record
	}
	return m
}

func invalidRow(row []string) bool {
	if row[PhoneNumber] == "PHONE_NUMBER" || row[ActivationDate] == "" || row[PhoneNumber] == "" {
		return true
	}
	return false
}

func convertString2Date(date string) (Date, error) {
	s := strings.Split(date, "-")
	d := [3]int{}
	for i := 0; i < 3; i++ {
		n, err := strconv.Atoi(s[i])
		if err != nil {
			return Date{}, err
		}
		d[i] = n
	}
	return Date{d[0], d[1], d[2]}, nil
}

func convertDate2String(date Date) string {
	year := normalizeString(date.Year, 4)
	month := normalizeString(date.Month, 2)
	day := normalizeString(date.Day, 2)
	s := []string{year, month, day}
	return strings.Join(s, "-")
}

func normalizeString(d int, n int) string {
	s := strconv.Itoa(d)
	if len(s) < n {
		s = strings.Repeat("0", n-len(s)) + s
	}
	return s
}

func updateRecord(m map[string]Record, row []string) (Record, error) {
	record, exists := m[row[PhoneNumber]]
	adate, err := convertString2Date(row[ActivationDate])
	if err != nil {
		return Record{}, err
	}

	dedate := Date{}
	if row[DeactivationDate] != "" {
		dedate, err = convertString2Date(row[DeactivationDate])
		if err != nil {
			return Record{}, err
		}
		if isGreater(adate, dedate) {
			return Record{}, errors.New("activation date is greater than deactivation date")
		}
	}

	if !exists {
		record = Record{adate, dedate, []Stuff{}}
	} else {
		updateExistedRecord(&record, adate, dedate)
	}
	return record, nil
}

func updateExistedRecord(record *Record, adate Date, dedate Date) {
	changed := true
	if isEqual(dedate, Date{0, 0, 0}) {
		if !isEqual(record.finish, adate) && !isEqual(record.finish, Date{0, 0, 0}) {
			record.start = adate
		}
		record.finish = dedate
	} else if isEqual(record.start, dedate) {
		record.start = adate
	} else if isEqual(record.finish, adate) {
		record.finish = dedate
	} else if isGreater(adate, record.finish) && !isEqual(record.finish, Date{0, 0, 0}) {
		record.stuffs = append(record.stuffs, Stuff{record.start, record.finish})
		record.finish = dedate
		record.start = adate
	} else {
		changed = false
		record.stuffs = append(record.stuffs, Stuff{adate, dedate})
	}

	if changed {
		// check current start/finish with stuffs
		for i, stuff := range record.stuffs {
			compareWithStuff(record, stuff, i)
		}
	}
}

func compareWithStuff(record *Record, stuff Stuff, pos int) {
	if isEqual(record.start, stuff.finish) {
		record.start = stuff.start
		if pos < len(record.stuffs) {
			record.stuffs = append(record.stuffs[:pos], record.stuffs[pos+1:]...)
		}
	}
}

func isEqual(d1 Date, d2 Date) bool {
	if d1.Year == d2.Year && d1.Month == d2.Month && d1.Day == d2.Day {
		return true
	}
	return false
}

func isGreater(d1 Date, d2 Date) bool {
	if d1.Year > d2.Year {
		return true
	}
	if d1.Year < d2.Year {
		return false
	}

	if d1.Month > d2.Month {
		return true
	}
	if d1.Month < d2.Month {
		return false
	}

	if d1.Day > d2.Day {
		return true
	}
	if d1.Day < d2.Day {
		return false
	}

	// equal
	return false
}

func writeFile(m map[string]Record, filename string) error {
	output := "result." + filename
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{"PHONE_NUMBER", "REAL_ACTIVATION_DATE"})
	for key, val := range m {
		line := []string{key, convertDate2String(val.start)}
		err := w.Write(line)
		if err != nil {
			return err
		}
	}
	fmt.Println("Open file", output)
	return nil
}
