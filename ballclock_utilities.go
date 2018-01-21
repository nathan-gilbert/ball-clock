package main

//RemoveIndex -- removes a ball from a queue and right shifts all
//following elements. Returns value at index as well as new queue
//with index removed
func RemoveIndex(queue []int, index int) []int {
	return append(queue[:index], queue[index+1:]...)
}

//ReverseQueue -- simple reversing of the balls in a given array
func ReverseQueue(queue []int) []int {
	for i := 0; i < len(queue)/2; i++ {
		j := len(queue) - i - 1
		queue[i], queue[j] = queue[j], queue[i]
	}
	return queue
}
