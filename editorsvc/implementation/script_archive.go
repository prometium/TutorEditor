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

type scriptArchiveController struct {
	ScriptJsonBytes      []byte
	imageFileBytesByName map[string][]byte
}

func (controller *scriptArchiveController) init(script *editorsvc.Script, imagesDir string) error {
	scriptJsonBytes, err := json.MarshalIndent(script, "", " ")
	if err != nil {
		return err
	}

	controller.ScriptJsonBytes = scriptJsonBytes

	controller.imageFileBytesByName = make(map[string][]byte)
	for _, frame := range script.Frames {
		imageFileBytes, err := ioutil.ReadFile(filepath.Join(imagesDir, frame.PictureLink))
		if err != nil {
			return err
		}

		controller.imageFileBytesByName[frame.PictureLink] = imageFileBytes

		// TODO: изображения switch pictures
	}

	return nil
}

func (controller *scriptArchiveController) getArchive() ([]byte, error) {
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

	for name, file := range controller.imageFileBytesByName {
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
