package fswebcam

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestFrameGrab(t *testing.T) {

	args := strings.Split("-d v4l2:/dev/video1 -S 2 -p MJPEG -q -r 960x720 -F 2 -fps 30", " ")
	bytes, err := GetCaptureJPEGBytes(args, 100)
	if err != nil {
		t.Error(err)
	}
	if len(bytes) < 1000 {
		t.Errorf("Len of captured bytes should be greater %d\n", len(bytes))
	}
	ioutil.WriteFile("output.jpg", bytes, 0666)
}

func BenchmarkFrameGrab(b *testing.B) {

	args := strings.Split("-d v4l2:/dev/video1 -S 1 -p MJPEG -q -r 960x720 -F 1", " ")
	for i := 0; i < b.N; i++ {
		GetCaptureJPEGBytes(args, 100)
	}
}
