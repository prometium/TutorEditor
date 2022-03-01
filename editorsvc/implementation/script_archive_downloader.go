package implementation

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/prometium/tutoreditor/editorsvc"
)

type scriptArchiveDownloader struct {
	ScriptJsonBytes      []byte
	ImageFileBytesByName map[string][]byte
}

func (controller *scriptArchiveDownloader) init(ctx context.Context, script *editorsvc.Script, minioClient *minio.Client, bucketName string) error {
	scriptJsonBytes, err := json.MarshalIndent(script, "", " ")
	if err != nil {
		return err
	}

	controller.ScriptJsonBytes = scriptJsonBytes

	controller.ImageFileBytesByName = make(map[string][]byte)
	for _, frame := range script.Frames {
		imageFileReader, err := minioClient.GetObject(ctx, bucketName, frame.PictureLink, minio.GetObjectOptions{})
		if err != nil {
			return err
		}
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(imageFileReader)
		if err != nil {
			return err
		}

		imageFileBytes := buf.Bytes()

		controller.ImageFileBytesByName[frame.PictureLink] = imageFileBytes

		for _, action := range frame.Actions {
			if action.SwitchPictures == nil {
				continue
			}

			for _, switchPicture := range action.SwitchPictures {
				imageFileReader, err := minioClient.GetObject(ctx, bucketName, switchPicture.PictureLink, minio.GetObjectOptions{})
				if err != nil {
					return err
				}
				buf := new(bytes.Buffer)
				_, err = buf.ReadFrom(imageFileReader)
				if err != nil {
					return err
				}

				imageFileBytes := buf.Bytes()

				controller.ImageFileBytesByName[switchPicture.PictureLink] = imageFileBytes
			}
		}
	}

	return nil
}

func (controller *scriptArchiveDownloader) getArchive() ([]byte, error) {
	var zipBuffer *bytes.Buffer = new(bytes.Buffer)
	var zipWriter *zip.Writer = zip.NewWriter(zipBuffer)
	var scriptJsonWriter io.Writer

	scriptJsonWriter, err := zipWriter.Create("Script.json")
	if err != nil {
		return nil, err
	}

	_, err = scriptJsonWriter.Write(controller.ScriptJsonBytes)
	if err != nil {
		return nil, err
	}

	for name, file := range controller.ImageFileBytesByName {
		imageFileWriter, err := zipWriter.Create(name)
		if err != nil {
			return nil, err
		}

		_, err = imageFileWriter.Write([]byte(file))
		if err != nil {
			return nil, err
		}
	}

	err = zipWriter.Close()
	if err != nil {
		return nil, err
	}

	return zipBuffer.Bytes(), nil
}
