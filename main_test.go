package main

import "testing"

func Test_cashback(t *testing.T) {

	tests := []struct {
		name string
		amount int
		want int
	}{
		// TODO: Add test cases.
		{"Have cashback", amount:5000, want:250}
		{"No cashback", amount:1000, want:0}
		{"Cashback on bound", amount:3000, want:150}
	}
	for _, test := range tests {
		got := cashback(test.amount)
		if got != test.want{
			t.error("forcashbaxk test", test.name, "got:", got, "want:", test.want)
	}
	}

	}
}