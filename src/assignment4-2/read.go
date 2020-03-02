package main
import (
	"fmt"
	"io/ioutil"
)

type FullName struct {
	fname string
	lname string
}

func main() {
	var fileName string
	
	fmt.Scan(&fileName)
	sli := make([]FullName, 0, 5)
	dat, err := ioutil.ReadFile(fileName)
	if(err == nil) {
		for i := 0; i < len(dat); i+=44 {
			var fullname FullName
			j := i + 43
			fullname.fname = string(dat[i:j])[0:20]
			fullname.lname = string(dat[i:j])[21:41]
			sli = append(sli, fullname)
		}
		
		for _, v := range sli {
			fmt.Println(v.fname, v.lname)
		}
	}
}