package implementation

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"editorsvc"
)

type service struct {
	repository editorsvc.Repository
}

// NewService makes a new Service
func NewService(rep editorsvc.Repository) editorsvc.Service {
	return &service{
		repository: rep,
	}
}

func (*service) Status(ctx context.Context) (string, error) {
	return "Ok", nil
}

func (s *service) AddRawScript(ctx context.Context, archiveReader io.Reader) (int, error) {
	zipBytes, err := ioutil.ReadAll(archiveReader)
	if err != nil {
		return 0, err
	}

	reader := bytes.NewReader(zipBytes)
	zipReader, err := zip.NewReader(reader, int64(len(zipBytes)))

	for _, file := range zipReader.File {
		//fmt.Printf("=%s\n", file.Name)

		if file.Name == "Script.json" {
			var raw script
			scriptJSON, err := readAll(file)
			if err != nil {
				return 0, err
			}

			json.Unmarshal([]byte(scriptJSON), &raw)
			s.repository.AddScript(ctx, convertToFrames(raw))
			break
		}
	}

	return 0, nil
}

func readAll(f *zip.File) ([]byte, error) {
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

func convertToFrames(raw script) []editorsvc.Frame {
	result := make([]editorsvc.Frame, len(raw.Frames))
	for i, frame := range raw.Frames {
		result[i] = editorsvc.Frame{
			UID:         fmt.Sprintf("_:%d", frame.FrameNumber),
			PictureLink: frame.PictureLink,
			Task:        frame.Task,
			Hint:        frame.Hint,
			DType:       []string{"Frame"},
		}
	}
	return result
}

// func convertToFramesActionsMutation(rawScript script) []editorsvc.Frame {

// }
