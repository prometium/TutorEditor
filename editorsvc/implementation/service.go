package implementation

import (
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/prometium/tutoreditor/editorsvc"
	"github.com/prometium/tutoreditor/editorsvc/utils"
)

type service struct {
	repository editorsvc.Repository
}

const VersionLen = 10
const ImagesDir = "assets/images/"

// NewService makes a new Service
func NewService(rep editorsvc.Repository) editorsvc.Service {
	return &service{
		repository: rep,
	}
}

func (s *service) AddScriptArchive(ctx context.Context, name string, fileReader io.ReadCloser) (string, error) {
	if fileReader == nil {
		return "", editorsvc.ErrFileNotAttached
	}
	defer fileReader.Close()

	var controller rawScriptArchiveSaver
	if err := controller.init(fileReader); err != nil {
		return "", err
	}

	linksMap, err := controller.saveImages(ctx, ImagesDir)
	if err != nil {
		return "", err
	}

	script, err := controller.createScript(name, linksMap)
	if err != nil {
		return "", err
	}

	script.Version = utils.RandSeq(VersionLen)
	script.ModificationDate = time.Now().Format("2006.01.02 15:04:05")
	id, err := s.repository.AddScript(ctx, script)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *service) AddScriptArchiveV2(ctx context.Context, name string, fileReader io.ReadCloser) (string, error) {
	if fileReader == nil {
		return "", editorsvc.ErrFileNotAttached
	}
	defer fileReader.Close()

	var controller scriptArchiveSaver
	if err := controller.init(fileReader); err != nil {
		return "", err
	}

	err := controller.saveImages(ctx, ImagesDir)
	if err != nil {
		return "", err
	}

	script := controller.createScript(name)

	script.Version = utils.RandSeq(VersionLen)
	script.ModificationDate = time.Now().Format("2006.01.02 15:04:05")
	id, err := s.repository.AddScript(ctx, script)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *service) GetScriptArchiveV2(ctx context.Context, id string) ([]byte, error) {
	script, err := s.repository.GetScript(ctx, id)
	if err != nil {
		return nil, err
	} else if script == nil {
		return nil, editorsvc.ErrScriptNotFound
	}

	var controller scriptArchiveDownloader
	if err := controller.init(script, ImagesDir); err != nil {
		return nil, err
	}

	return controller.getArchive()
}

func (s *service) GetScriptsList(ctx context.Context) ([]editorsvc.Script, error) {
	return s.repository.GetScriptsList(ctx)
}

func (s *service) GetScript(ctx context.Context, id string) (*editorsvc.Script, error) {
	script, err := s.repository.GetScript(ctx, id)
	if err != nil {
		return nil, err
	} else if script == nil {
		return nil, editorsvc.ErrScriptNotFound
	}
	return script, err
}

func (s *service) DeleteScript(ctx context.Context, id string) error {
	return s.repository.DeleteScript(ctx, id)
}

func (s *service) UpdateScript(
	ctx context.Context, script *editorsvc.Script, frameIdsToDel []string, actionIdsToDel []string,
) (map[string]string, error) {
	script.Version = utils.RandSeq(VersionLen)
	script.ModificationDate = time.Now().Format("2006.01.02 15:04:05")
	uids, err := s.repository.UpdateScript(ctx, script, frameIdsToDel, actionIdsToDel)
	if err != nil {
		return uids, err
	}
	return uids, nil
}

func (s *service) CopyScript(ctx context.Context, script *editorsvc.Script) (string, error) {
	script.Version = utils.RandSeq(VersionLen)
	return s.repository.AddScript(ctx, script)
}

func (s *service) AddImage(ctx context.Context, imageReader io.ReadCloser) (string, error) {
	defer imageReader.Close()

	var buf bytes.Buffer
	teeReader := io.TeeReader(imageReader, &buf)

	hash, err := utils.HashFileMD5(teeReader)
	if err != nil {
		return "", err
	}

	path := filepath.Join(ImagesDir, hash+".png")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = utils.CopyFile(&buf, path)
		if err != nil {
			return "", err
		}
	}

	return hash + ".png", nil
}
