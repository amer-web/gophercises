package manage

import (
	"encoding/binary"
	"encoding/json"
	"github.com/boltdb/bolt"
	"task/app"
)

type Task struct {
	Id      int    `json:"id"`
	Details string `json:"details"`
	Status  bool   `json:"status"`
}

func (t *Task) CreateTask(u *Task) error {
	db := app.GetDB()
	return db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("tasks"))
		id, _ := b.NextSequence()
		u.Id = int(id)

		buf, err := json.Marshal(u)
		if err != nil {
			return err
		}
		return b.Put(itob(u.Id), buf)
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func ListTasks() []Task {
	db := app.GetDB()
	var tasks []Task
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))

		b.ForEach(func(k, v []byte) error {
			var task Task
			json.Unmarshal(v, &task)
			tasks = append(tasks, task)
			return nil
		})
		return nil
	})
	return tasks
}
