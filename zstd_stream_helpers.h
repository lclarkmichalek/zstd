#ifndef GOZSTD_STREAM_HELPERS
#define GOZSTD_STREAM_HELPERS

#include "zstd.h"

size_t GoZSTD_compressStream(
  ZSTD_CStream* zds,
  void *dst_buf,
  size_t dst_size,
  size_t *dst_pos,
  void *src_buf,
  size_t src_size,
  size_t *src_pos
);

size_t GoZSTD_flushStream(
  ZSTD_CStream* zds,
  void *dst_buf,
  size_t dst_size,
  size_t *dst_pos
);

size_t GoZSTD_endStream(
  ZSTD_CStream* zds,
  void *dst_buf,
  size_t dst_size,
  size_t *dst_pos
);

size_t GoZSTD_decompressStream(
  ZSTD_DStream* zds,
  void *dst_buf,
  size_t dst_size,
  size_t *dst_pos,
  void *src_buf,
  size_t src_size,
  size_t *src_pos
);

#endif /* GOZSTD_STREAM_HELPERS */
