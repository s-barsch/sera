package paths

import (
	"reflect"
	"testing"
)

func TestSplitFile(t *testing.T) {
	tests := []struct {
		name     string
		filep    string
		expected *File
		wantErr  bool
	}{
		{
			name:  "VTT file",
			filep: "vtt/subtitle.en.vtt",
			expected: &File{
				Name:   "vtt/subtitle.vtt",
				Option: "en",
				Ext:    "vtt",
			},
			wantErr: false,
		},
		{
			name:  "Media file",
			filep: "160403_124512-1600.jpg",
			expected: &File{
				Name:   "160403_124512.jpg",
				Option: "1600",
				Ext:    "jpg",
			},
			wantErr: false,
		},
		{
			name:     "Invalid file",
			filep:    "invalid-file",
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SplitFile(tt.filep)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("SplitFile() = %v, want %v", got, tt.expected)
			}
		})
	}
}
