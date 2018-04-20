package toy

import "testing"

func Test_nextPowerOf2(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "5",
			args: args{n: 5},
			want: 8,
		},
		{
			name: "17",
			args: args{n: 17},
			want: 32,
		},
		{
			name: "32",
			args: args{n: 32},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPowerOf2(tt.args.n); got != tt.want {
				t.Errorf("nextPowerOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_RandRange(t *testing.T) {
	tests := []struct {
		name string
		size uint
	}{
		{
			name: "5",
			size: 5,
		},
		{
			name: "17",
			size: 17,
		},
		{
			name: "32",
			size: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandRange(tt.size); got >= tt.size {
				t.Errorf("RandRange() = %v, size %v", got, tt.size)
			} else {
				t.Logf("RandRange() = %v, size %v", got, tt.size)
			}
		})
	}
}
