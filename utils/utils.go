package utils

import (
	"archive/zip"
	"bytes"
	"hash/fnv"
	"io"
	"io/ioutil"
	"os"
)

// Hash generate hash number of a string
func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// ReadAllFromZip reads from zip file until an error or EOF and returns the data it read
func ReadAllFromZip(f *zip.File) ([]byte, error) {
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	content, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// CreateZipReader returns a new zip Reader reading from r
func CreateZipReader(r io.Reader) (*zip.Reader, error) {
	zipBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(zipBytes)
	zipReader, err := zip.NewReader(reader, int64(len(zipBytes)))
	if err != nil {
		return nil, err
	}
	return zipReader, nil
}

// CopyZipFile copies the src zip file to dst
func CopyZipFile(src *zip.File, dst string) error {
	in, err := src.Open()
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
