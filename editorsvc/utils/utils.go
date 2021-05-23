package utils

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandSeq generates a random sequence
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// HashZipFileMD5 generates MD5 hash number of a zip file
func HashZipFileMD5(f *zip.File) (string, error) {
	fReader, err := f.Open()
	if err != nil {
		return "", err
	}
	defer fReader.Close()

	return HashFileMD5(fReader)
}

// HashFileMD5 generates MD5 hash number of a file
func HashFileMD5(fReader io.ReadCloser) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, fReader); err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String := hex.EncodeToString(hashInBytes)
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
	inReader, err := src.Open()
	if err != nil {
		return err
	}
	defer inReader.Close()

	return CopyFile(inReader, dst)
}

// CopyFile copies the inReader file to dst
func CopyFile(inReader io.Reader, dst string) error {
	outReader, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer outReader.Close()

	_, err = io.Copy(outReader, inReader)
	if err != nil {
		return err
	}
	return outReader.Close()
}
