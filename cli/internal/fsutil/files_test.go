package fsutil

import "testing"

func TestWriteToFile(t *testing.T) {
	type args struct {
		content string
		path    string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteToFile(tt.args.content, tt.args.path)
		})
	}
}

func TestReadFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadFile(tt.args.file); got != tt.want {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
