package gotest

import (
	"errors"
	"testing"
)

type User struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

func validateUser(user *User) error {
	if user.EmailAddress == "" {
		return errors.New("user requires an email")
	}
	if len(user.Password) < 8 {
		return errors.New("user password requires at least 8 characters")
	}

	return nil
}

func Test_validateUser(t *testing.T) {
	type args struct {
		user *User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test validateUser raise",
			args:    args{&User{EmailAddress: "", Password: "12345678"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("validateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
