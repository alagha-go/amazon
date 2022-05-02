package crawler

import (
	mydepartments "amazon/lib/crawler/departments"
	"time"
)


func Main() {
	for {
		mydepartments.CollectAllDepartments()
		for index:=0; index<5; index++ {
			mydepartments.UpdateDepartments()
		} 
		time.Sleep(720*time.Hour)
	}
}