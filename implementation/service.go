package implementation

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"strconv"

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

func (s *service) AddRawScript(ctx context.Context, name string, archiveReader io.Reader) (string, error) {
	zipBytes, err := ioutil.ReadAll(archiveReader)
	if err != nil {
		return "", err
	}

	reader := bytes.NewReader(zipBytes)
	zipReader, err := zip.NewReader(reader, int64(len(zipBytes)))
	if err != nil {
		return "", err
	}

	for _, file := range zipReader.File {
		//fmt.Printf("=%s\n", file.Name)

		if file.Name == "Script.json" {
			var rs rawScript
			scriptJSON, err := readAll(file)
			if err != nil {
				return "", err
			}

			json.Unmarshal([]byte(scriptJSON), &rs)
			id, err := s.repository.AddScript(ctx, name, convertToFrames(rs))
			if err != nil {
				return "", err
			}
			return id, nil
		}
	}

	return "", nil
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

func convertToFrames(rs rawScript) []editorsvc.Frame {
	frames := make([]editorsvc.Frame, len(rs.Frames))
	for i, frame := range rs.Frames {
		frames[i] = editorsvc.Frame{
			UID:         strconv.Itoa(frame.FrameNumber),
			PictureLink: frame.PictureLink,
			Task:        frame.Task,
			Hint:        frame.Hint,
		}
	}
	for i, frame := range rs.Frames[1:] {
		action := frame.ActionSwitch
		frames[i].Actions = []editorsvc.Action{
			editorsvc.Action{
				NextFrame: editorsvc.NextFrame{
					UID: strconv.Itoa(frame.FrameNumber),
				},
				ActionType:   action.ActionType,
				XLeft:        action.XLeft,
				XRight:       action.XRight,
				YLeft:        action.YLeft,
				YRight:       action.YRight,
				StartXLeft:   action.StartXLeft,
				StartYLeft:   action.StartYLeft,
				StartXRight:  action.StartXRight,
				StartYRight:  action.StartYRight,
				FinishXLeft:  action.FinishXLeft,
				FinishYLeft:  action.FinishYLeft,
				FinishXRight: action.FinishXRight,
				FinishYRight: action.FinishYRight,
				TicksCount:   action.TicksCount,
				Key:          action.Key,
				ModKey:       action.ModKey,
			},
		}
	}
	return frames
}
