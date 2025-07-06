package split

import (
	"fmt"
	"strings"
)

// 160403_124512-1600.jpg
// (name)	     (option)(ext)
func SplitFile(filep string) (*File, error) {
	if filep == "" {
		return nil, nil
	}
	if len(filep) >= len("vtt/x.de.vtt") && filep[:3] == "vtt" {
		return splitVTT(filep)
	}
	return splitMedia(filep)
}

// splitMedia: for images and video

func splitMedia(filep string) (*File, error) {
	i := strings.LastIndex(filep, "-")
	j := strings.LastIndex(filep, ".")

	return splitFileParameters(filep, i, j)
}

// Sample filepath: x.de.vtt

func splitVTT(filep string) (*File, error) {
	i := strings.Index(filep, ".")
	j := strings.LastIndex(filep, ".")

	return splitFileParameters(filep, i, j)
}

// new filename: 160403_124512-1600.jpg -> (160403_124512.jpg) (1600)

func splitFileParameters(filep string, i, j int) (*File, error) {
	if i < 0 || j < 0 || i == j {
		return nil, fmt.Errorf("splitFileParameters: errornous filename: %v", filep)
	}

	return &File{
		Name:   filep[:i] + filep[j:],
		Option: filep[i+1 : j],
		Ext:    filep[j+1:],
	}, nil
}
