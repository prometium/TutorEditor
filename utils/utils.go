package utils

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
)

// HashZipFileMD5 generate MD5 hash number of a zip file
func HashZipFileMD5(f *zip.File) (string, error) {
	var returnMD5String string
	file, err := f.Open()
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
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
