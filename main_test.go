package main

import "testing"

func Test_main(t *testing.T) {
	status := goMain([]string{"oilio", "-h"})
	if status != 0 {
		t.Errorf("status code 0, but got %d", status)
	}
}
