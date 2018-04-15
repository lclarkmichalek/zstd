package zstd

import "testing"

func TestStreamCompressor(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	orig := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sc := NewStreamCompressor(nil)
	defer sc.Close()
	buf := make([]byte, SuggestedStreamOutputSize())
	writableBuf := buf
	written := 0
	for len(data) != 0 {
		nWritten, nRead, err := sc.Compress(writableBuf, data)
		if err != nil {
			t.Errorf("compress failed: %v", err)
		}
		data = data[nRead:]
		writableBuf = writableBuf[nWritten:]
		written += nWritten
	}

	n, err := sc.End(writableBuf)
	if err != nil {
		t.Errorf("end failed: %v", err)
	}
	written += n

	t.Logf("compressed data: %v", buf[:written])

	dc := NewStreamDecompressor(nil)
	defer dc.Close()
	out := make([]byte, len(orig)*2)
	t.Logf("out len: %v", len(out))
	n, _, err = dc.Decompress(out, buf[:written])
	if err != nil {
		t.Errorf("decompress failed: %v", err)
	}

	t.Logf("decompressed data: %v", out[:n])
}
