package main

import (
	"IoT-learning/client"
	"fmt"
	"time"
)



func main() {
	p := client.New("tcp://localhost:1883", "parham")
	r := client.New("tcp://localhost:1883", "raha")

	p.Connect()
	r.Connect("love")

	for i := 0;i < 5 ; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		p.Publish("love", text)
	}

	time.Sleep(3 * time.Second)

	//if token := r.C.Unsubscribe("love"); token.Wait() && token.Error() != nil {
	//	fmt.Println(token.Error())
	//	os.Exit(1)
	//}

	p.C.Disconnect(250)
	r.C.Disconnect(250)
}
