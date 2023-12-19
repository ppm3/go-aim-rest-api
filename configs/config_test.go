package configs

import (
	"context"
	"os"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	var projectDirName string = "go-aim-rest-api"
	var cxt context.Context = context.Background()
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
					Host:           "localhost",
					Port:           "27017",
					Username:       "root",
					Password:       "password",
					Database:       "mydatabase",
					Protocol:       "mongodb",
					AuthSource:     "admin",
					MaxPoolSize:    1,
					ConnectTimeout: 1,
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
			name: "Returned error for max env",
			want: nil,
			args: args{
				env:            "testing_error_max",
				projectDirName: projectDirName,
				ctx:            cxt,
			},
			wantErr: true,
		},
		{
			name: "Returned error for connect env",
			want: nil,
			args: args{
				env:            "testing_error_connect",
				projectDirName: projectDirName,
				ctx:            cxt,
			},
			wantErr: true,
		},
		{
			name:    "Unset environment",
			want:    nil,
			wantErr: true,
			args: args{
				env:            "",
				projectDirName: projectDirName,
				ctx:            cxt,
			},
		},
		{
			name:    "Invalid configuration file",
			want:    nil,
			wantErr: true,
			args: args{
				env:            "noenv",
				projectDirName: projectDirName,
				ctx:            cxt,
			},
		},
		{
			name:    "wrong project dir name",
			want:    nil,
			wantErr: true,
			args: args{
				env:            "noenv",
				projectDirName: "notexist",
				ctx:            cxt,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv() // Clear environment variables between runs
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
