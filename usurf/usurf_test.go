package usurf

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		ua string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should be false",
			args: args{
				ua: THBadUserAgent,
			},
			want: false,
		},
		{
			name: "should be true ios",
			args: args{
				ua: IOSUserAgent,
			},
			want: true,
		},
		{
			name: "should be true android",
			args: args{
				ua: AndroidUserAgent,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.ua); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsValidBadUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := IsValid(THBadUserAgent)
		_ = x
	}
}

func BenchmarkIsValidIOS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := IsValid(THBadUserAgent)
		_ = x
	}
}

func BenchmarkIsValidAndroid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := IsValid(THBadUserAgent)
		_ = x
	}
}
