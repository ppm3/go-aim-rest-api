package configs

import (
	"context"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	var projectDirName string = "go-aim-rest-api"
	var config *LoadServerConfig = NewLoadServerConfig()
	type args struct {
		ctx            context.Context
		env            string
		projectDirName string
	}
	tests := []struct {
		name    string
		args    args
		want    *ServerConfig
		wantErr bool
	}{
		{
			name: "Valid configuration file",
			want: &ServerConfig{
				Server: ServerInitConfig{
					Url:  "localhost",
					Port: "9999",
				},
				Mysql: MysqlDBConfig{
					Host:     "localhost",
					Port:     "3306",
					Username: "root",
					Password: "password",
					Database: "mydatabase",
				},
				Mongo: MongoDBConfig{
					Host:     "localhost",
					Port:     "27017",
					Username: "root",
					Password: "password",
					Database: "mydatabase",
				},
				ProjectName:    "test",
				ProjectVersion: "1.0.0",
				Environment:    "test",
			},
			args: args{
				env:            "testing",
				projectDirName: projectDirName,
				ctx:            context.Background(),
			},
			wantErr: false,
		},
		{
			name:    "Unset environment",
			want:    nil,
			wantErr: true,
			args: args{
				env:            "",
				projectDirName: projectDirName,
				ctx:            context.Background(),
			},
		},
		{
			name:    "Invalid configuration file",
			want:    nil,
			wantErr: true,
			args: args{
				env:            "noenv",
				projectDirName: projectDirName,
				ctx:            context.Background(),
			},
		},
		{
			name:    "wrong project dir name",
			want:    nil,
			wantErr: true,
			args: args{
				env:            "noenv",
				projectDirName: "notexist",
				ctx:            context.Background(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := config.Load(tt.args.ctx, tt.args.env, tt.args.projectDirName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
