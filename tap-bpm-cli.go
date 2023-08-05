package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
)

func getAverage(arr []float32, size int) float32 {
   var i int
   var sum float32
   var avg float32  

   for i = 0; i < size;i++ {
      sum += arr[i]
   }

   avg = sum / float32(size)
   return avg;
}

/*func printSlice(s []float32) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}*/

func getBPM(diff float32) float32 {
	bpm := 60000.0 / ( float32(diff) / seconds )
	return float32(bpm)
}

const seconds float32 = 1000.0*1000.0

func main() {

	fmt.Print("Tap enter to calculate BPM (counter starts from the next tap)\n");

	scanner := bufio.NewScanner(os.Stdin)
	var taps []float32
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
		taps = append(taps, getBPM(float32(diff)))

		//printSlice(taps)
		//fmt.Printf("%#v", float32(diff))

		fmt.Printf("\033[2A\033[1K") // Remove previous two line
		fmt.Printf("Current tap BPM: %v \n", getBPM(float32(diff)) )
		fmt.Printf("Average BPM: %v ", getAverage(taps, len(taps)))
	}    

}

