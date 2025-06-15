package operations

import "log"
func IsWork(err error, where string) bool{
	log.Print("Check: ", where)
	return err == nil
}
	

