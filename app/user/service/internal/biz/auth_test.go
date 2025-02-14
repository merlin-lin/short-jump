package biz

import (
	"fmt"
	"testing"
)

func Test_verifyPassword(t *testing.T) {
	type args struct {
		pwd     string
		pwdHash string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				pwd:     "123456",
				pwdHash: "$2a$10$WkHDQtA0bbCHfHtUnpXhmezfuCMt3lm8rgKPpa8Zr3m8t7kD0OAFK",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := verifyPassword(tt.args.pwd, tt.args.pwdHash); (err != nil) != tt.wantErr {
				t.Errorf("verifyPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_hashPassword(t *testing.T) {
	fmt.Println(hashPassword("123456"))

}
