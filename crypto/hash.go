package crypto

import (
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
)

func HashBytes(h hash.Hash, bs []byte) (string, int64, error) {
	n, err := h.Write(bs)
	if err != nil {
		return "", 0, err
	}
	return hex.EncodeToString(h.Sum(nil)), int64(n), nil
}

func HashFiles(h hash.Hash, files ...string) (string, int64, error) {
	if len(files) == 0 {
		return "", 0, fmt.Errorf("required at least one file.")
	}
	var count int64
	for _, v := range files {
		f, err := os.Open(v)
		if err != nil {
			return "", 0, err
		}

		if n, err := io.Copy(h, f); err != nil {
			return "", 0, fmt.Errorf("read %s failed, %s", f.Name(), err)
		} else {
			count += n
		}
	}

	return hex.EncodeToString(h.Sum(nil)), count, nil
}
