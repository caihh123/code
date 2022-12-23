package code

import (
	"errors"
	"testing"

	"github.com/boltdb/bolt"
)

func TestBolt(t *testing.T) {
	db, err := bolt.Open("testdata/bolt/mu.db", 0666, nil)
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("test"))
		if err != nil {
			return err
		}
		return bucket.Put([]byte("1"), []byte("haha"))
	})
	if err != nil {
		t.Error(err)
		return
	}

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("test"))
		if bucket == nil {
			return errors.New("bucket empty")
		}
		t.Log(string(bucket.Get([]byte("1"))))
		t.Log(string(bucket.Get([]byte("2"))))
		return nil
	})
}
