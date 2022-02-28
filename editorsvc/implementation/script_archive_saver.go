package implementation

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/minio/minio-go/v7"
	"github.com/prometium/tutoreditor/editorsvc"
	"github.com/prometium/tutoreditor/editorsvc/utils"
)

type scriptArchiveSaver struct {
	Images []*zip.File
	Script *editorsvc.Script `json:"script,omitempty"`
}

func (controller *scriptArchiveSaver) init(r io.Reader) error {
	zipReader, err := utils.CreateZipReader(r)
	if err != nil {
		return err
	}

	controller.Images = make([]*zip.File, 0, len(zipReader.File))
	var scriptFile *zip.File = nil
	for _, file := range zipReader.File {
		if filepath.Ext(strings.TrimSpace(file.Name)) == ".png" {
			controller.Images = append(controller.Images, file)
		} else if file.Name == "Script.json" {
			scriptFile = file
		}
	}

	scriptJSON, err := utils.ReadAllFromZip(scriptFile)
	if err != nil {
		return err
	}

	controller.Script = new(editorsvc.Script)
	if err := json.Unmarshal(scriptJSON, controller.Script); err != nil {
		return err
	}

	return nil
}

func (controller *scriptArchiveSaver) saveImages(ctx context.Context, minioClient *minio.Client, bucketName string) error {
	lock := sync.RWMutex{}
	errs, _ := errgroup.WithContext(ctx)
	for _, file := range controller.Images {
		currentFile := file
		errs.Go(func() error {
			bytesArray, err := utils.ReadAllFromZip(currentFile)
			buf := bytes.NewBuffer(bytesArray)

			objectName := currentFile.Name + ".png"
			objectSize := int64(buf.Len())

			_, err = minioClient.PutObject(ctx, bucketName, objectName, buf, objectSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
			if err != nil {
				return err
			}

			lock.Lock()
			defer lock.Unlock()

			return nil
		})
	}

	if err := errs.Wait(); err != nil {
		return err
	}
	return nil
}

func (controller *scriptArchiveSaver) createScript(name string) *editorsvc.Script {
	controller.Script.Name = name
	return controller.Script
}
