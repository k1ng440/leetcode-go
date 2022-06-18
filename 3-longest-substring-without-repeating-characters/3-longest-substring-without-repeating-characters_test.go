package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "abba",
			args: args{
				s: "abba",
			},
			want: 2,
		},
		{
			name: "dvdk",
			args: args{
				s: "dvdk",
			},
			want: 3,
		},
		{
			name: "empty space",
			args: args{
				s: " ",
			},
			want: 1,
		},
		{
			name: "all duplicates",
			args: args{
				s: "aaaaaaaaaaaaaaaaa",
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
