/*
 The package even implemets functions to test if a number is even or odd
*/

package even

//returns true when number is even
func Even(i int) bool { // exported function
	return i%2 == 0
}

// returns true when the number is odd
func odd(i int) bool { // private function
	return i%2 == 1
}
