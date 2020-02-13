// +build !better

package robotname


import (
	//"crypto/rand"
	"fmt"
	"math/rand"
	"time"
)


type Robot struct {
	name string
}


//name `^[A-Z]{2}[0-9]{3}$``
//TODO: see how slow it will be with crypto/rand
//fyi : playground always begin with the same time to this wont work .
func init() {
	/**
	key:= [256]byte{}
	_, err := rand.Read(key[:])
	if err != nil{
		panic(err)
	}
	fmt.Println(key)
	
	*/
	rand.Seed(time.Now().UnixNano())
}

func (r *Robot ) Name() (string,error) {
	var s string
	if r.name == "" {
		i,j := 65+ rand.Intn(26) ,65+ rand.Intn(26)
		s = fmt.Sprintf("%s%s%03d",string( i) , string(j) ,rand.Intn(999))
		//r.name = string(s)
		r.name = s
	} else	 {
		s = r.name
		
	}
	return string(s),nil
}

func (r *Robot ) Reset() {
	r.name = ""
}