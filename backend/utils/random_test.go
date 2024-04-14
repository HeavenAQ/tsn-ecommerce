package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	type args struct {
		min int64
		max int64
	}
	tests := []struct {
		name string
		args args
		want func(val int64) bool
	}{
		{
			name: "Test RandomInt",
			args: args{
				min: 0,
				max: 100,
			},
			want: func(val int64) bool {
				return val >= 0 && val <= 100
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomInt(tt.args.min, tt.args.max); !tt.want(got) {
				t.Errorf("RandomInt() = %v, want 0 <= val <= 100", got)
			}
		})
	}
}

func TestRandomString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test RandomString",
			args: args{
				n: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomAlphabetString(tt.args.n); len(got) != tt.want {
				t.Errorf("RandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomUserName(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Test RandomUserName",
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomUserName(); len(got) != tt.want {
				t.Errorf("RandomUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomPrice(t *testing.T) {
	tests := []struct {
		name string
		want func(val int64) bool
	}{
		{
			name: "Test RandomPrice",
			want: func(val int64) bool {
				return val >= 0 && val <= 1000
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomPrice(); !tt.want(got) {
				t.Errorf("RandomPrice() = %v, want 0 <= val <= 1000", got)
			}
		})
	}
}

func TestRandomDiscount(t *testing.T) {
	tests := []struct {
		name string
		want func(val int64) bool
	}{
		{
			name: "Test RandomDiscount",
			want: func(val int64) bool {
				return val >= 0 && val <= 100
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomDiscount(); !tt.want(got) {
				t.Errorf("RandomDiscount() = %v, want 0 <= val <= 100", got)
			}
		})
	}
}

func TestRandomLanguage(t *testing.T) {
	tests := []struct {
		name string
		want func(val string) bool
	}{
		{
			name: "Test RandomLanguage",
			want: func(val string) bool {
				return val == "chn" || val == "jp"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomLanguage(); !tt.want(got) {
				t.Errorf("RandomLanguage() = %v, want chn or jp", got)
			}
		})
	}
}

func TestRandomNumberString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want func(val string) bool
	}{
		{
			name: "Test RandomNumberString",
			args: args{
				n: 10,
			},
			want: func(val string) bool {
				require.Len(t, val, 10)
				require.Regexp(t, "^[0-9]+$", val)
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomNumberString(tt.args.n); !tt.want(got) {
				t.Errorf("RandomNumberString() = %v, want length = 10 and contains only digit 0-9", got)
			}
		})
	}
}
