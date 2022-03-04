package config

import (
	"os"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	type env struct {
		OauthBase         string
		ClientCallbackUrl string

		JwtPublicKey   string
		JwtPrivateKey  string
		JwtKeyId       string
		AuthGrpcServer string

		VkCleintId     string
		VkClientSercet string
	}

	type args struct {
		path string
		env  env
	}

	setEnv := func(env env) {
		os.Setenv("OAUTH_BASE", env.OauthBase)
		os.Setenv("CLIENT_CALLBACK_URL", env.ClientCallbackUrl)
		os.Setenv("AUTH_GRPC_SERVER", env.AuthGrpcServer)
		os.Setenv("JWT_PUBLIC_KEY", env.JwtPublicKey)
		os.Setenv("JWT_PRIVATE_KEY", env.JwtPrivateKey)
		os.Setenv("JWT_KEY_ID", env.JwtKeyId)
		os.Setenv("VKONTAKTE_CLIENT_ID", env.VkCleintId)
		os.Setenv("VKONTAKTE_CLIENT_SECRET", env.VkClientSercet)
	}

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "test config",
			args: args{
				env: env{
					OauthBase: "http://localhost:3000/oauth/",
					ClientCallbackUrl: "http://localhost/auth/callback",
					JwtPublicKey: "./data/certs/is_rsa.pub",
					JwtPrivateKey: "./data/certs/is_rsa",
					JwtKeyId: "qwerty",
					AuthGrpcServer: ":9002",
					VkCleintId: "12312",
					VkClientSercet: "secret",
				},
			},
			want: &Config{
				OauthBase: "http://localhost:3000/oauth/",
				ClientCallbackUrl: "http://localhost/auth/callback",
				JwtPublicKey: "./data/certs/is_rsa.pub",
				JwtPrivateKey: "./data/certs/is_rsa",
				JwtKeyId: "qwerty",
				AuthGrpcServer: ":9002",
				VkCleintId: "12312",
				VkClientSercet: "secret",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(tt.args.env)

			got, err := GetConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfig() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
