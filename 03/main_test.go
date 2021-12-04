package main

import (
	"testing"
)

func Test_findPowerConsumption(t *testing.T) {
	type args struct {
		fname string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{"test.txt"},
			198,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPowerConsumption(readInput(tt.args.fname)); got != tt.want {
				t.Errorf("findPowerConsumption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLifeSupportRating(t *testing.T) {
	type args struct {
		fname string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{"test.txt"},
			230,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLifeSupportRating(readInput(tt.args.fname)); got != tt.want {
				t.Errorf("findLifeSupportRating() = %v, want %v", got, tt.want)
			}
		})
	}
}
