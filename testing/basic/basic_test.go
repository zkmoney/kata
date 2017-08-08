package main

import "testing"

func TestMakeAwesome(t *testing.T) {
	tests := []struct {
		Name string
		In   string
		Want string
	}{
		{
			Name: "Conlin",
			In:   "Conlin",
			Want: "Conlin is awesome!",
		},
		{
			Name: "Nirag",
			In:   "Nirag",
			Want: "Nirag, eh",
		},
		{
			Name: "Shweta",
			In:   "Shweta",
			Want: "Shweta is awesome!",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			have := MakeAwesome(test.In)
			if have != test.Want {
				t.Errorf("Wanted: '%s', got '%s'", test.Want, have)
			}
		})
	}
}
