package utils


type Datastore struct {
	count int
	data map[int]string
	totalTime int64
}

func InitializeDatastore() *Datastore {
	return &Datastore{
		count: int(0),
		data: make(map[int]string),
		totalTime: int64(0),
	}
}
func (db *Datastore) Increment() int {
	db.count++
	return db.count
}

func (db *Datastore) Insert(id int, str string) int {
	db.data[id] = str
	return id
}

func (db *Datastore) CountTime(time int64) int64 {
	db.totalTime += time
	return db.totalTime
}

func (db *Datastore) GetAverageTimeCount() int64 {
	if db.count > 0 {
		return db.totalTime / int64(db.count)
	}
	return 0
}

func (db *Datastore) GetTotal() int {
	return db.count
}

func (db *Datastore)  FindOne(id int) string {
	return db.data[id]
}