package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func timeTrack(tb testing.TB, start time.Time) {
	elapsed := time.Since(start)
	tb.Logf("This test took %s\n", elapsed)
}

func TestFile(t *testing.T) {
	testcases := []struct {
		name     string
		filename string
	}{
		{"10 lines", "file.csv"},
		{"duplicate rows", "file2.csv"},
//		{"500k lines", "jZZKJWqZVsFM.csv"},
//		{"duplicate 500k lines", "wKoVnoHBuqBe.csv"},
		{"file3.csv", "file3.csv"},
		{"file4.csv", "file4.csv"},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			start := time.Now()
			defer timeTrack(t, start)
			err := readFile(tc.filename)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestIsEqual(t *testing.T) {
	testcases := []struct {
		name        string
		d1          Date
		d2          Date
		expectation bool
	}{
		{"equal", Date{2017, 1, 1}, Date{2017, 1, 1}, true},
		{"not equal", Date{2016, 2, 3}, Date{2017, 3, 5}, false},
		{"not equal with zero", Date{2014, 2, 4}, Date{0, 0, 0}, false},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isEqual(tc.d1, tc.d2)
			assert.Equal(t, tc.expectation, actual, "Wrong!!")
		})
	}
}

func TestIsGreater(t *testing.T) {
	testcases := []struct {
		name        string
		d1          Date
		d2          Date
		expectation bool
	}{
		{"equal", Date{2017, 1, 1}, Date{2017, 1, 1}, false},
		{"smaller", Date{2016, 2, 3}, Date{2017, 3, 5}, false},
		{"greater", Date{2017, 7, 3}, Date{2017, 3, 5}, true},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isGreater(tc.d1, tc.d2)
			assert.Equal(t, tc.expectation, actual, "Wrong!!")
		})
	}
}

func TestUpdateExistedRecord(t *testing.T) {
	testcases := []struct {
		adate       Date
		dedate      Date
		record      *Record
		expectation *Record
	}{
		{Date{2017, 1, 1}, Date{0, 0, 0}, &Record{Date{2016, 2, 3}, Date{2016, 7, 8}, []Stuff{}}, &Record{Date{2017, 1, 1}, Date{0, 0, 0}, []Stuff{}}},
		{Date{2017, 1, 1}, Date{0, 0, 0}, &Record{Date{2016, 2, 3}, Date{2017, 1, 1}, []Stuff{}}, &Record{Date{2016, 2, 3}, Date{0, 0, 0}, []Stuff{}}},
		{Date{2017, 1, 1}, Date{2017, 3, 1}, &Record{Date{2017, 3, 1}, Date{2017, 5, 1}, []Stuff{}}, &Record{Date{2017, 1, 1}, Date{2017, 5, 1}, []Stuff{}}},
	}
	for _, tc := range testcases {
		updateExistedRecord(tc.record, tc.adate, tc.dedate)
		assert.Equal(t, tc.expectation.start, tc.record.start, "Wrong start day")
		assert.Equal(t, tc.expectation.finish, tc.record.finish, "Wrong finish day")
	}
}

func TestNormalizedString(t *testing.T) {
	testCases := []struct {
		name   string
		d      int
		n      int
		expect string
	}{
		{"year", 12, 4, "0012"},
		{"month", 12, 2, "12"},
		{"day", 1, 2, "01"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := normalizeString(tc.d, tc.n)
			assert.Equal(t, tc.expect, s, "Wrong")
		})
	}
}

func TestConvertString2Date(t *testing.T) {
	testCases := []struct {
		date     string
		expect   Date
		isFailed bool
	}{
		{"2013-02-03", Date{2013, 2, 3}, false},
		{"2012-03-04", Date{2012, 3, 4}, false},
		{"wewqewe", Date{}, true},
	}
	for _, tc := range testCases {
		d, err := convertString2Date(tc.date)
		if err != nil {
			if !tc.isFailed {
				t.Fatal(err)
			}
			continue
		}
		if !reflect.DeepEqual(d, tc.expect) {
			t.Fatal("Not match")
		}
	}
}

func TestConvertDate2String(t *testing.T) {
	testCases := []struct {
		date   Date
		expect string
	}{
		{Date{2013, 2, 3}, "2013-02-03"},
		{Date{2014, 5, 31}, "2014-05-31"},
		{Date{1, 1, 1}, "0001-01-01"},
	}
	for _, tc := range testCases {
		s := convertDate2String(tc.date)
		assert.Equal(t, tc.expect, s, "Wrong")
	}
}
