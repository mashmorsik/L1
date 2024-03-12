package main

import (
	"fmt"
	"time"
)

/* Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

func main() {
	var c = make(chan string)

	go WriteToChan(c, time.Nanosecond)

	// читаем из канала пока приходят сообщения
	for {
		val, ok := <-c
		if !ok {
			break
		} else {
			fmt.Println(val)
		}
	}
}

// WriteToChan пишет в канал, пока не закончится таймер, после окончания таймера, канал закрывается
func WriteToChan(c chan string, dur time.Duration) {
	messages := []string{"hello", "from", "the", "other", "side", "I", "must", "have", "called"}

	timer := time.NewTimer(dur)

	for i := 0; i < len(messages); i++ {
		select {
		case <-timer.C:
			close(c)
			return
		default:
			c <- messages[i]
		}
	}

}
