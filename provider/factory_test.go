package provider

import (
	"reflect"
	"testing"
)

func TestNewProvider(t *testing.T) {
	type args struct {
		providerType Type
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name:    "speedtestnet provider",
			args:    args{providerType: SpeedTestNetProvider},
			wantErr: false,
		},
		{
			name:    "fastcom provider",
			args:    args{providerType: FastComProvider},
			wantErr: false,
		},
		{
			name:    "unknown provider",
			args:    args{providerType: -1},
			wantErr: true,
			err:     ErrUnknownProviderType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProvider(tt.args.providerType)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("NewProvider() = %v, want %v", got, tt.err)
			}
		})
	}
}
