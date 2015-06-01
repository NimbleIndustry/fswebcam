package fswebcam

//fswebcam -d v4l2:/dev/video1 -S 2 -p MJPEG -v -r 1280x720 -fps 30 foo.jpg

/*

#include <stdio.h>
#include <getopt.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <time.h>
#include <locale.h>
#include <gd.h>
#include <errno.h>
#include <signal.h>
#include <sys/types.h>
#include <sys/stat.h>
#include "fswebcam.h"
#include "log.h"
#include "src.h"
#include "dec.h"
#include "effects.h"
#include "parse.h"

#cgo LDFLAGS: -lgd
#cgo CFLAGS: -g -O2 -DHAVE_CONFIG_H

extern gdImage* getCapture(int argc, char *argv[]);

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func GetCaptureJPEGBytes(args []string, quality int) ([]byte, error) {
	argc := C.int(len(args))
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
	}

	img := C.getCapture(argc, &argv[0])
	if img == nil {
		return nil, fmt.Errorf("Error capturing frame with %v", args)
	}
	bytes := ImageToJPEGBuffer(img, quality)

	C.gdImageDestroy(img)
	for _, arg := range argv {
		C.free(unsafe.Pointer(arg))
	}
	return bytes, nil
}

func ImageToJPEGBuffer(p *C.gdImage, quality int) []byte {
	var imgSize int
	pimgSize := (*C.int)(unsafe.Pointer(&imgSize))

	buf := C.gdImageJpegPtr(p, pimgSize, C.int(quality))
	defer C.gdFree(buf)

	return C.GoBytes(buf, *pimgSize)
}
