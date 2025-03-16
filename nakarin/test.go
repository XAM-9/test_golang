package nakarin

import (
	"fmt"
)

func generateTest() { // private function
	fmt.Println("Test")
}

func SayTest() { // public function
	generateTest()
}
