package playground

import (
	"testing"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func removeAccents(s string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, err := transform.String(t, s)
	if err != nil {
		return "", err
	}
	return output, nil
}

func Test_removeAccents(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "given áááá should return aaaa",
			args: args{
				s: "áááá",
			},
			want:    "aaaa",
			wantErr: false,
		},
		{
			name: "given ãã should return aa",
			args: args{
				s: "ãã",
			},
			want:    "aa",
			wantErr: false,
		},
		{
			name: "given âââ should return aaa",
			args: args{
				s: "âââ",
			},
			want:    "aaa",
			wantErr: false,
		},
		{
			name: "given àà should return aa",
			args: args{
				s: "àà",
			},
			want:    "aa",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := removeAccents(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("removeAccents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("removeAccents() = %v, want %v", got, tt.want)
			}
		})
	}
}
