package main
import(
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.waitGroup
	wg.Add(2)
	str := "hello,world!"
	str1 := []byte(str)
	sc := make(chan byte, len(str))
	count :=make(chan int)
	for _, v := range str1{
		sc<-v
	}
	close(sc)

}

go func(){
	defer wg.Done()
	for{
		ball,ok:=<-count
		if ok {
			pr1,ok1:=<-sc
			if ok1{
				fmt.Printf("go 2:%c\n",pr1)
			}else{
				close(count)
				return
			}
			count <- ball
		}else{
			return
		}
	}

}()

go func(){
	defer wg.Done()
	for{
		time.Sleep(8 *time.Millisecond)
	}
}