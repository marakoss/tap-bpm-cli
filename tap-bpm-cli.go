package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
	_ "math"
)

/*type Queue struct {
    elements chan interface{}
}

func NewQueue(size int) *Queue {
    return &Queue{
        elements: make(chan interface{}, size),
    }
}

func (queue *Queue) ReadAll() interface{} {
	select {
	case _ := <-queue.elements:
		return queue.elements
	default:
		panic("Queue empty")
	}
	return nil
}

func (queue *Queue) Push(element interface{}) {
    select {
    case queue.elements <- element:
    default:
        panic("Queue full")
    }
}

func (queue *Queue) Pop() interface{} {
    select {
    case e := <-queue.elements:
        return e
    default:
        panic("Queue empty")
    }
    return nil
}
*/

func getAverage(arr []int, size int) float32 {
   var i, sum int
   var avg float32  

   for i = 0; i < size;i++ {
      sum += arr[i]
   }

   avg = float32(sum / size)
   return avg;
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func getBPM(diff int) int {
	bpm := 60000 / ( int(diff) / seconds )
	return int(bpm)
}

const seconds int = 1000*1000 

func main() {

	fmt.Print("Tap enter to calculate BPM (counter starts from the next tap)\n");
	//fmt.Print("Current tap BPM: \n")
	//fmt.Print("Average BPM: ")

	scanner := bufio.NewScanner(os.Stdin)
	var taps []int
	var i int

	for {

		if i < 1 {
			scanner.Scan()
			i +=1
        		fmt.Print("Current tap BPM: \n")
        		fmt.Print("Average BPM: ")
			continue
		}

		start := time.Now()
        	scanner.Scan()
		end := time.Now()
		diff := end.Sub(start)

		i += 1

		if i >= 10  {
			copy(taps, taps[1:])
			taps = taps[:len(taps) - 1]
		}
		taps = append(taps, getBPM(int(diff)))

		//printSlice(taps)
		//fmt.Printf("%#v", int(diff))

		fmt.Printf("\033[2A\033[1K") // Remove previous two line
		fmt.Printf("Current tap BPM: %v \n", getBPM(int(diff)) )
		fmt.Printf("Average BPM: %v ", getAverage(taps, len(taps)))
	}    

}

