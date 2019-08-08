package main

import (
	"reflect"
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{
			arg:  "Yogesh",
			want: "Hello Yogesh!",
		}, {
			arg:  "Rishabh",
			want: "Hello Rishabh!",
		},
	}
	for _, test := range tests {
		if got := Greet(test.arg); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Test failed: got %+v, want %+v", got, test.want)
		}
	}
}
