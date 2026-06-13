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

/*
func Test_isValidValue(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		c    byte
		want bool
	}{
		{
			name: "Test1",
			c:    byte(48),
			want: true,
		},
		{
			name: "Test2",
			c:    byte(0),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidValue(tt.c)
			if got != tt.want {
				t.Errorf("isValidValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

*/
