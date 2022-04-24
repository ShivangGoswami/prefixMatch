package filereader

import (
	"log"
	"testing"
)

func TestReadFile(t *testing.T) {
	logFatal = log.Println
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"invalid",
			args{path: ""},
			0,
		},
		{
			"valid",
			args{path: "../source/sample_prefixes.txt"},
			264986,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadFile(tt.args.path); len(got) != tt.want {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
