package utils

import (
	"errors"
	"fmt"
)

type Datastore struct {
	count     int
	data      map[int]string
	totalTime int64
	timeCount int
}

func InitializeDatastore() *Datastore {
	return &Datastore{
		count:     int(0),
		data:      make(map[int]string),
		totalTime: int64(0),
	}
}
func (db *Datastore) Increment() int {
	db.count++
	return db.count
}

func (db *Datastore) Insert(id int, str string) (int, error) {
	existing := db.FindOne(id)
	if existing != "" {
		return -1, errors.New(fmt.Sprintf("Duplicate id %v", id))
	}
	db.data[id] = str
	return id, nil
}

func (db *Datastore) CountTime(time int64) (int64, error) {
	if time == int64(0) {
		return -1, errors.New("non-zero time required")
	}
	db.totalTime += time
	db.timeCount++
	return db.totalTime, nil
}

func (db *Datastore) GetAverageTimeCount() int64 {
	if db.timeCount > 0 {
		return db.totalTime / int64(db.timeCount)
	}
	return 0
}

func (db *Datastore) GetTotal() int {
	return db.count
}

func (db *Datastore) FindOne(id int) string {
	return db.data[id]
}
