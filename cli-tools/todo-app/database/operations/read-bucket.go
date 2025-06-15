package operations

import (
	"errors"
	"log"
	"github.com/boltdb/bolt"
)
func ReadDB(db *bolt.DB) error{
	// failed at db.View
	var err error
	if err = db.View(func(tx *bolt.Tx) error { 
			b := tx.Bucket([]byte("tasks"))
			if b == nil {
//				log.Fatal("No bucket found")
				return errors.New("No bucket found")
			}
			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				log.Printf("key=%s, value=%s\n", k, v)
			}
		return nil
	}); err != nil{
//		log.Print("failed to read")
//		log.Fatal(err, "Here")
		return err
	} else{
//		log.Print("Success to read")
		return nil
	}
}
