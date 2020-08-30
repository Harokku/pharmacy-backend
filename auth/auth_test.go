package auth

import "testing"

var (
	plainPwd  string = "TestPwd"
	hashedPwd string = "$2a$10$1rMPOk2n.p/J2/.MZHqD2enzDdICvWHlAzMyRuiUGM9Zwx/gfZDnS"
)

func TestComparePassword(t *testing.T) {
	type args struct {
		hashPwd  string
		plainPwd string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty password",
			args: args{
				hashPwd:  hashedPwd,
				plainPwd: "",
			},
			want: false,
		},
		{
			name: "Wrong password",
			args: args{
				hashPwd:  hashedPwd,
				plainPwd: "1234",
			},
			want: false,
		},
		{
			name: "Correct password",
			args: args{
				hashPwd:  hashedPwd,
				plainPwd: plainPwd,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComparePassword(tt.args.hashPwd, tt.args.plainPwd); got != tt.want {
				t.Errorf("ComparePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashAndSalt(t *testing.T) {
	type args struct {
		pwd string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Empty password",
			args:    args{pwd: ""},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashAndSalt(tt.args.pwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashAndSalt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HashAndSalt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
