package configs

import (
	"context"
	"testing"
)

func TestReadEnvfile(t *testing.T) {
	type args struct {
		ctx            context.Context
		env            string
		projectDirName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Read the .env file and works successfully",
			args: args{
				ctx:            context.Background(),
				env:            "test",
				projectDirName: "go-aim-rest-api",
			},
			wantErr: false,
		},
		{
			name: "Error when reading the .env file",
			args: args{
				ctx:            context.Background(),
				env:            "",
				projectDirName: "go-aim-rest-api",
			},
			wantErr: true,
		},
		{
			name: "Error when recived wrong directory name",
			args: args{
				ctx:            context.Background(),
				env:            "test",
				projectDirName: "x",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadEnvfile(tt.args.ctx, tt.args.env, tt.args.projectDirName); (err != nil) != tt.wantErr {
				t.Errorf("ReadEnvfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
