package operations

import "github.com/boltdb/bolt"
// reset database mean clear all task in bucketName - tasks
func DeleteBucket(db *bolt.DB, bucketName []byte) error{
	return db.Update(func(tx *bolt.Tx) error{
		return tx.DeleteBucket(bucketName)
	})
}
