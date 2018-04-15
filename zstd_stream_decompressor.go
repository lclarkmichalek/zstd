package zstd

/*
#define ZSTD_STATIC_LINKING_ONLY
#include "zstd.h"
#include "zbuff.h"
#include "zstd_stream_helpers.h"

*/
import "C"
import (
	"unsafe"
)

// StreamDecompressor is a wrapper around zstd's streaming decompression primitives
type StreamDecompressor struct {
	Options CompressionOptions
	ctx     *C.ZSTD_DStream
}

func NewStreamDecompressor(options *CompressionOptions) *StreamDecompressor {
	if options == nil {
		options = &DefaultCompressionOptions
	}

	ctx := C.ZSTD_createDStream()
	if options.Dict == nil {
		C.ZSTD_initDStream(ctx)
	} else {
		C.ZSTD_initDStream_usingDict(
			ctx,
			unsafe.Pointer(&options.Dict[0]),
			C.size_t(len(options.Dict)),
		)
	}

	return &StreamDecompressor{
		Options: *options,
		ctx:     ctx,
	}
}

func (sc *StreamDecompressor) Close() error {
	if sc.ctx != nil {
		C.ZSTD_freeDStream(sc.ctx)
		sc.ctx = nil
	}
	return nil
}

func (sc *StreamDecompressor) Decompress(dst, src []byte) (int, int, error) {
	if sc.ctx == nil {
		return 0, 0, ErrContextReleased
	}

	var dstPos, srcPos C.size_t
	retCode := C.GoZSTD_decompressStream(
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
