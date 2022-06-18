package coinchange

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "11 = 5 + 5 + 1",
		// 	args: args{
		// 		coins:  []int{4, 5, 1},
		// 		amount: 11,
		// 	},
		// 	want: 3,
		// },
		{
			name: "6249 [186,419,83,408]",
			args: args{
				coins:  []int{186, 419, 83, 408},
				amount: 6249,
			},
			want: 20,
		},
		// {
		// 	name: "3 [2]",
		// 	args: args{
		// 		coins:  []int{2},
		// 		amount: 3,
		// 	},
		// 	want: -1,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
