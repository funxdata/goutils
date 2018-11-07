package crypto

import (
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
)

func HashFiles(h hash.Hash, files ...string) (string, error) {
	if len(files) == 0 {
		return "", fmt.Errorf("required at least one file.")
	}
	for _, v := range files {
		f, err := os.Open(v)
		if err != nil {
			return "", err
		}

		if _, err := io.Copy(h, f); err != nil {
			return "", fmt.Errorf("read %s failed, %s", f.Name(), err)
		}
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
