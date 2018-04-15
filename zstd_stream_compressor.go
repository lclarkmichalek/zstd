package zstd

/*
#define ZSTD_STATIC_LINKING_ONLY
#include "zstd.h"
#include "zbuff.h"
#include "zstd_stream_helpers.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

var (
	ErrContextReleased = errors.New("context for this stream has been released")
)

func SuggestedStreamInputSize() int {
	return int(C.ZSTD_CStreamInSize())
}

func SuggestedStreamOutputSize() int {
	return int(C.ZSTD_CStreamOutSize())
}

type CompressionOptions struct {
	CompressionLevel int
	Dict             []byte
}

// DefaultCompressionOptions are the options that will be used when nil is
// passed to a function expecting CompressionOptions
var DefaultCompressionOptions = CompressionOptions{
	CompressionLevel: DefaultCompression,
	Dict:             nil,
}

// StreamCompressor is a wrapper around zstd's streaming compression primitives
type StreamCompressor struct {
	Options CompressionOptions
	ctx     *C.ZSTD_CStream
}

func NewStreamCompressor(options *CompressionOptions) *StreamCompressor {
	if options == nil {
		options = &DefaultCompressionOptions
	}

	ctx := C.ZSTD_createCStream()
	if options.Dict == nil {
		C.ZSTD_initCStream(ctx, C.int(options.CompressionLevel))
	} else {
		C.ZSTD_initCStream_usingDict(
			ctx,
			unsafe.Pointer(&options.Dict[0]),
			C.size_t(len(options.Dict)),
			C.int(options.CompressionLevel),
		)
	}

	return &StreamCompressor{
		Options: *options,
		ctx:     ctx,
	}
}

func (sc *StreamCompressor) Close() error {
	if sc.ctx != nil {
		C.ZSTD_freeCStream(sc.ctx)
		sc.ctx = nil
	}
	return nil
}

func (sc *StreamCompressor) Compress(dst, src []byte) (int, int, error) {
	if sc.ctx == nil {
		return 0, 0, ErrContextReleased
	}

	var dstPos, srcPos C.size_t
	retCode := C.GoZSTD_compressStream(
		sc.ctx,
		unsafe.Pointer(&dst[0]),
		C.size_t(len(dst)),
		&dstPos,
		unsafe.Pointer(&src[0]),
		C.size_t(len(src)),
		&srcPos,
	)
	if err := getError(int(retCode)); err != nil {
		return 0, 0, err
	}
	return int(dstPos), int(srcPos), nil
}

func (sc *StreamCompressor) Flush(dst []byte) (int, error) {
	if sc.ctx == nil {
		return 0, ErrContextReleased
	}

	var dstPos C.size_t
	retCode := C.GoZSTD_flushStream(
		sc.ctx,
		unsafe.Pointer(&dst[0]),
		C.size_t(len(dst)),
		&dstPos,
	)
	if err := getError(int(retCode)); err != nil {
		return 0, err
	}
	return int(dstPos), nil
}

func (sc *StreamCompressor) End(dst []byte) (int, error) {
	if sc.ctx == nil {
		return 0, ErrContextReleased
	}

	var dstPos C.size_t
	retCode := C.GoZSTD_endStream(
		sc.ctx,
		unsafe.Pointer(&dst[0]),
		C.size_t(len(dst)),
		&dstPos,
	)
	if err := getError(int(retCode)); err != nil {
		return 0, err
	}
	return int(dstPos), nil
}
