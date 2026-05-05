package validpalindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[正常系] 回文をはtrueを返却する",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "[正常系] 回文ではない文はfalseを返却する",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "[正常系] 空白文字の場合はtrueを返却する",
			args: args{
				s: " ",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
