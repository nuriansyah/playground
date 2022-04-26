package main

func squareWorker(input <-chan int, output chan<- int) {
	for i := 1; i <= 10; i++ {
		//mengirim ke channel output hasil	pangkat 2 nilai dari channel input
		output <- i * i
	}
}

//mengirim 1-n angka ke squareWorker
func numberGenerator(n int, input chan<- int) {
	for i := 1; i <= n; i++ {
		//mengirim ke channel input angka i
		input <- i
	}
}
