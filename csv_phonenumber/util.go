package main

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const blockSize = 5000000

type Block struct {
	start  int
	finish int
}

func setUp(n int) (string, error) {
	filename := generateFileName(12) + ".csv"
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}

	jobs := make(chan string)
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()

	firstLine := []string{"PHONE_NUMBER", "ACTIVATION_DATE", "DEACTIVATION_DATE"}
	w.Write(firstLine)

	times := n / blockSize
	if n%blockSize != 0 {
		times++
	}
	blocks := []Block{}
	for i := 0; i < times; i++ {
		finish := blockSize * (i + 1)
		if i == times-1 && n%blockSize != 0 {
			finish = i*blockSize + n%blockSize
		}
		blocks = append(blocks, Block{blockSize * i, finish})
	}
	for _, b := range blocks {
		go writeBlock(b.start, b.finish, jobs, filename)
	}
	for i := 0; i < times; i++ {
		tempFile := <-jobs
		appendTempFile(w, tempFile)
	}
	return filename, nil
}

func writeBlock(start int, finish int, jobs chan string, filename string) {
	name := []string{"temp", filename, strconv.Itoa(start), strconv.Itoa(finish), "csv"}
	temp := strings.Join(name, ".")

	f, err := os.Create(temp)
	if err != nil {
		jobs <- ""
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for i := start; i < finish; i++ {
		//phoneNumber := generatePhoneNumber(10)
		line := []string{"09432408329432", "2016-01-02", "2017-01-01"}
		//line := []string{phoneNumber, generateDate(), generateDate()}
		err := w.Write(line)
		if err != nil {
			panic(err)
		}
	}
	jobs <- temp
}

func appendTempFile(w *csv.Writer, tempFile string) {
	defer os.Remove(tempFile)
	temp, err := os.Open(tempFile)
	if err != nil {
		panic(err)
	}
	defer temp.Close()

	r := csv.NewReader(temp)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	err = w.WriteAll(records)
	if err != nil {
		panic(err)
	}
}

func generatePhoneNumber(n int) string {
	return randStringRunes(n, []rune("1234567890"))
}

func generateFileName(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return randStringRunes(n, letters)
}

func generateNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func generateDate() string {
	year := generateNumber(2000, 2018)
	month := generateNumber(1, 13)
	day := 31
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 4, 6, 9, 11:
		day = 30
	case 2:
		if year%4 == 0 {
			day = 29
		} else {
			day = 28
		}
	default:
		day = 0
	}
	date := Date{year, month, day}
	return convertDate2String(date)
}

func randStringRunes(n int, letters []rune) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
