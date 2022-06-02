package utils

import (
	"testing"
)

var db Datastore

func setup() {
	db = Datastore{
		count:     0,
		data:      make(map[int]string),
		totalTime: 0,
	}
}

func TestDatastore_Insert(t *testing.T) {
	setup()

	want := 1
	result, err := db.Insert(want, "test")
	if err != nil {
		t.Fatalf(`Insert error %q`, err)
	}
	if want != result {
		t.Fatalf(`Insert = %q, Expected: %q`, result, want)
	}
}

func TestDatastore_InsertDuplicate(t *testing.T) {
	setup()

	want := 1
	_, initialError := db.Insert(want, "test")
	if initialError != nil {
		t.Fatalf(`Insert error %q`, initialError)
	}
	_, err := db.Insert(want, "test")
	if err == nil {
		t.Fatalf(`Insert expected error, got nil`)
	}

}

func TestDatastore_Increment(t *testing.T) {
	setup()
	n := 0
	for n < 5 {
		result := db.Increment()
		if result != n+1 {
			t.Fatalf(`Increment error, expected %v, got %v`, n+1, result)
		}
		n += 1
	}
}

func TestDatastore_CountTime(t *testing.T) {
	setup()
	inputList := [8]int64{1, 2, 3, 4, 5, 123123, 41234, 41241241241}
	var result = int64(0)
	var err error
	expectedSum := int64(0)
	for _, v := range inputList {
		result, err = db.CountTime(v)
		if err != nil {
			t.Fatalf(`CountTime error %v`, err)
		}
		expectedSum += v
	}

	if result != expectedSum {
		t.Fatalf(`CountTime error, expected %v, got %v`, expectedSum, result)
	}

}

func TestDatastore_CountTime_Zero(t *testing.T) {
	setup()
	_, err := db.CountTime(0)
	if err == nil {
		t.Fatalf(`CountTime expected error, got nil`)
	}
}

func TestDatastore_GetAverageTimeCount(t *testing.T) {
	setup()
	inputList := [8]int64{1, 2, 3, 4, 5, 123123, 41234, 41241241241}
	var err error
	for _, v := range inputList {
		_, err = db.CountTime(v)
		if err != nil {
			t.Fatalf(`CountTime error %v`, err)
		}
	}
	result := db.GetAverageTimeCount()
	expected := db.totalTime / int64(len(inputList))
	if result != expected {
		t.Fatalf(`GetAverageTimeCount error, expected %v got %v`, expected, result)
	}
}

func TestDatastore_FindOne(t *testing.T) {
	TestDatastore_Insert(t)
	result := db.FindOne(1)
	expected := "test"
	if result != expected {
		t.Fatalf(`FindOne error, expected %v got %v`, expected, result)
	}

}