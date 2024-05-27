package questions

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type DatabaseModel struct {
	Data         []question `json:"data"`
	CreatedAt    string     `json:"created-at"`
	LastEditedAt string     `json:"last-edited-at"`
}

type question struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Discription string `json:"discription"`
	SolvedAt    string `json:"solved-at"`
}

func (t *DatabaseModel) Load(filePath string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *DatabaseModel) Store(filePath string) error {
	t.LastEditedAt = getCurrentTime()
	if len(t.Data) == 1 {
		t.CreatedAt = getCurrentTime()
	}
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

func (t *DatabaseModel) Add(title string, link string, discription string) error {
	if title == "" || !validateLink(link) || discription == "" {
		return errors.New("Invalid args")
	}
	currentTime := getCurrentTime()
	newQuestion := question{Title: title, Link: link, Discription: discription, SolvedAt: currentTime}
	t.Data = append(t.Data, newQuestion)
	return nil
}
func (t *DatabaseModel) List() {
	for _, q := range t.Data {
		fmt.Println(q.Title, q.Discription)
	}
}
