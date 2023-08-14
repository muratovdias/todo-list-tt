package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const invalidTitle = "Ipd7VSJUJYQj43GbvnN2jslXV7VhwRPMZaxfaGCTlhJahDCSinldI1aE15nXs7jutakDxdQU5F0YPUZ5P9JQNlmk2A1zwNtdpooDgez4V0x9A2DGp3yLl6AWd0MOZCgCXPXEr2ZtAWlD58GPggamk0Db759MOBD5HqSKMhg5f3JI7CZMgnYAbOxRmSVTHGcX2f089UXCU"

func Test_isWeekend(t *testing.T) {

	tests := []struct {
		name    string
		arg     string
		want    bool
		wantErr bool
	}{
		{
			name: "OK",
			arg:  "2023-08-13",
			want: true,
		},
		{
			name: "Not Weekend",
			arg:  "2023-08-09",
		},
		{
			name: "Invalid Date",
			arg:  "2023-08-09",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isWeekend(tt.arg)
			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			} else {
				require.Error(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_validateDate(t *testing.T) {

	tests := []struct {
		name    string
		arg     string
		wantErr bool
	}{
		{
			name: "OK",
			arg:  "2023-08-14",
		},
		{
			name: "Invalid Date",
			arg:  "2023--14",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.wantErr {
				require.NoError(t, validateDate(tt.arg))
			} else {
				require.Equal(t, validateDate(tt.arg), ErrInvalidDate)
			}
		})
	}
}

func Test_validateTitle(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		wantErr bool
	}{
		{
			name: "OK",
			arg:  "test",
		},
		{
			name: "Invalid Titile",
			arg:  invalidTitle,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTitle(tt.arg)
			if !tt.wantErr {
				require.NoError(t, err)
			} else {
				require.Equal(t, err, ErrInvalidTitle)
			}
		})
	}
}
