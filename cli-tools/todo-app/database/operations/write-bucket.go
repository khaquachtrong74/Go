package operations

import (
//	"log"
	"strconv"
	"github.com/boltdb/bolt"
)
/* 
	task: string
	number: for index task
*/
func Write(db *bolt.DB, task string, number int) (string, error){
	if	err := db.Update(func(tx *bolt.Tx) error{
		b, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil{
//			log.Fatal("Create bucket failed")
			return err
		}
//		log.Print("Create bucket success")
		if err := b.Put([]byte(strconv.Itoa(number)), []byte(task)); err != nil{
//			log.Fatal("Failed to put task")
			return err
		}
//		log.Print("Put task success")
		return nil
	}); err != nil{
//		log.Fatal(err)
		return "Write task failed", err
	}
	return "Write task success",nil 
}
