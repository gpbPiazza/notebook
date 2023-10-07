package main

import "testing"

// https://bismobaruno.medium.com/unit-test-http-request-in-golang-a96d146406e6

func Test_getRFCLetters(t *testing.T) {
	type args struct {
		rfcID int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "aa",
			args: args{
				rfcID: 100,
			},
			want:    "AAA",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRFCLetters(tt.args.rfcID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRFCLetters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getRFCLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
