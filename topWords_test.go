package main

import (
	"reflect"
	"testing"
)

func TestTopWords(t *testing.T) {
	s := "a b c d a c a b a b a d b c ok l"
	expected := []string{"a", "b"}

	result := TopWords(s,2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect result.")
	}
}
