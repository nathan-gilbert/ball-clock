package main

//RemoveIndex -- removes a ball from a queue and right shifts all following elements.
func RemoveIndex(queue []int, index int) []int {
	return append(queue[:index], queue[index+1:]...)
}

//ReverseQueue -- simple reversal of an array
func ReverseQueue(queue []int) []int {
	for i := 0; i < len(queue)/2; i++ {
		j := len(queue) - i - 1
		queue[i], queue[j] = queue[j], queue[i]
	}
	return queue
}
