package cryptoutil

import (
	"reflect"
	"testing"
)

func Test_generateRandomKey(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateRandomKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateRandomKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getKey(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
