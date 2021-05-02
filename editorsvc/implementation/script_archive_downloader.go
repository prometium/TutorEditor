package implementation

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/prometium/tutoreditor/editorsvc"
)

type scriptArchiveDownloader struct {
	ScriptJsonBytes      []byte
	ImageFileBytesByName map[string][]byte
}

func (controller *scriptArchiveDownloader) init(script *editorsvc.Script, imagesDir string) error {
	scriptJsonBytes, err := json.MarshalIndent(script, "", " ")
	if err != nil {
		return err
	}

	controller.ScriptJsonBytes = scriptJsonBytes

	controller.ImageFileBytesByName = make(map[string][]byte)
	for _, frame := range script.Frames {
		var imageFileBytes []byte

		imageFileBytes, err := ioutil.ReadFile(filepath.Join(imagesDir, frame.PictureLink))
		if err != nil {
			return err
		}

		controller.ImageFileBytesByName[frame.PictureLink] = imageFileBytes

		for _, action := range frame.Actions {
			if action.SwitchPictures == nil {
				continue
			}

			for _, switchPicture := range action.SwitchPictures {
				imageFileBytes, err := ioutil.ReadFile(filepath.Join(imagesDir, switchPicture.PictureLink))
				if err != nil {
					return err
				}

				controller.ImageFileBytesByName[frame.PictureLink] = imageFileBytes
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
