#include "zstd_stream_helpers.h"

size_t GoZSTD_compressStream(
  ZSTD_CStream* zds,
  void *dst_buf,
  size_t dst_size,
  size_t *dst_pos,
  void *src_buf,
  size_t src_size,
  size_t *src_pos
) {
  struct ZSTD_outBuffer_s dst;
  dst.pos = 0;
  dst.size = dst_size;
  dst.dst = dst_buf;
  struct ZSTD_inBuffer_s src;
  src.pos = 0;
  src.size = src_size;
  src.src = src_buf;

  size_t ret = ZSTD_compressStream(zds, &dst, &src);
  *dst_pos = dst.pos;
  *src_pos = src.pos;
  return ret;
};

size_t GoZSTD_flushStream(
  ZSTD_CStream* zds,
  void *dst_buf,
  size_t dst_size,
  size_t *dst_pos
) {
  struct ZSTD_outBuffer_s dst;
  dst.pos = 0;
  dst.size = dst_size;
  dst.dst = dst_buf;

  size_t ret = ZSTD_flushStream(zds, &dst);
  *dst_pos = dst.pos;
  return ret;
};

size_t GoZSTD_endStream(
  ZSTD_CStream* zds,
  void *dst_buf,
  size_t dst_size,
  size_t *dst_pos
) {
  struct ZSTD_outBuffer_s dst;
  dst.pos = 0;
  dst.size = dst_size;
  dst.dst = dst_buf;

  size_t ret = ZSTD_endStream(zds, &dst);
  *dst_pos = dst.pos;
  return ret;
};

size_t GoZSTD_decompressStream(
  ZSTD_DStream* zds,
  void *dst_buf,
  size_t dst_size,
  size_t *dst_pos,
  void *src_buf,
  size_t src_size,
  size_t *src_pos
) {
  struct ZSTD_outBuffer_s dst;
  dst.pos = 0;
  dst.size = dst_size;
  dst.dst = dst_buf;
  struct ZSTD_inBuffer_s src;
  src.pos = 0;
  src.size = src_size;
  src.src = src_buf;

  size_t ret = ZSTD_decompressStream(zds, &dst, &src);
  *dst_pos = dst.pos;
  *src_pos = src.pos;
  return ret;
};
