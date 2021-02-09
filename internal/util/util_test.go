package util

import (
	"reflect"
	"testing"
)


func Test_splitSep(t *testing.T) {
	type args struct {
		str string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "t1",
			args: args{
				str: "aaa|aab",
				sep: "|",
			},
			want: []string{"aaa", "aab"},
		},
		
		{
			name: "t2",
			args: args{
				str: "aaa  aab",
				sep: " ",
			},
			want: []string{"aaa", "aab"},
		},
		
		{
			name: "t3",
			args: args{
				str: "aaa",
				sep: " ",
			},
			want: []string{"aaa"},
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitSep(tt.args.str, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitSep() = %v, want %v", got, tt.want)
			}
		})
	}
}
