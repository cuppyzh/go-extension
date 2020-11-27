package strinformatter

import (
	"fmt"
	"testing"
)

func TestCompose(t *testing.T) {
	_, err := Compose("@<%v> Test Value %v", "VAL 1", "VAL 2")
	if err != nil {
		t.Errorf("Failed for test case 1")
	}

	val, err := Compose("@<%v> Test Value %v", "%v", "VAL 2")
	if err != nil || val != "@<%v> Test Value VAL 2" {
		t.Errorf("Failed for test case 2")
	}

	// Should error because parameter is mismatch
	_, err = Compose("@<%v> Test Value %v", "%v")
	if err == nil {
		t.Errorf("Failed for test case 3")
	}

	// Should error because parameter is mismatch
	_, err = Compose("@<%v> Test Value", "%v", "%v")
	if err == nil {
		t.Errorf("Failed for test case 4")
	}

	// Should error because parameter is mismatch
	_, err = Compose("@<%v> Test Value")
	if err == nil {
		t.Errorf("Failed for test case 5")
	}

	// Should error because parameter is mismatch
	resp, err := Compose("@<%v> Test Value", "test")
	if resp != "@<test> Test Value" {
		fmt.Println(resp)
		t.Errorf("Failed for test case 5")
	}
}
