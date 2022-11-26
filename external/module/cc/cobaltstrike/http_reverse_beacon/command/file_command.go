package command

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type FileCommand struct {
	*Commnad

	FilePath    string
	FileContent []byte
	Operate     string
}

func (c *FileCommand) Unmarshall(data []byte) error {
	err := json.Unmarshal(data, &c)
	if err != nil {
		return errors.New("unmarshall failed")
	}
	return nil
}

// Implement large file uploads
func (c *FileCommand) Upload() {

}

// Implement large file downloads
func (c *FileCommand) Download() {

}

func (c *FileCommand) ListDir() {

}

func (c *FileCommand) WriteFile() {
	ioutil.WriteFile(c.FilePath, c.FileContent, 0777)
}

func (c *FileCommand) ReadFile() {
	ioutil.ReadFile(c.FilePath)
}
