package reversebits

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{
				n: 43261596,
			},
			want: 964176192,
		},
		{
			name: "Test2",
			args: args{
				n: 2147483644,
			},
			want: 1073741822,
		},
		// エッジケース
		{
			name: "Test3",
			args: args{
				n: 0b10000000000000000000000000000000,
			},
			want: 0b00000000000000000000000000000001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.n); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
