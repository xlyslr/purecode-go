package q0567

import (
	"fmt"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	inclusion := checkInclusion("eidbaooo", "ab")
	fmt.Printf("%#v\n", inclusion)
}
