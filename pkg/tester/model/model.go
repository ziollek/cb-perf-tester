package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Difficulty int

const (
	Easy       Difficulty = 0
	Medium                = 1
	Hard                  = 2
	Impossible            = 3
	Unknown               = 4
)

func FromString(difficultyName string) Difficulty {
	switch strings.ToLower(difficultyName) {
	case "easy":
		return Easy
	case "medium":
		return Medium
	case "hard":
		return Hard
	case "impossible":
		return Impossible
	}
	return Unknown
}

func (d Difficulty) ToString() string {
	switch d {
	case Easy:
		return "Easy"
	case Medium:
		return "Medium"
	case Hard:
		return "Hard"
	case Impossible:
		return "Impossible"
	}
	return "Unknown"
}

type Document struct {
	Key  string            `json:"key"`
	Data map[string]string `json:"data"`
	Size int               `json:"size"`
}

func Generate(key string, size int) *Document {
	data := make(map[string]string, size)
	for i := 0; i < size; i++ {
		data[subKey(i)] = fmt.Sprintf("value-%06d", i)
	}
	return &Document{
		Key:  key,
		Data: data,
		Size: size,
	}
}

func subKey(i int) string {
	return fmt.Sprintf("subkey-%06d", i)
}

func subKeyPath(i int) string {
	return fmt.Sprintf("data.%s", subKey(i))
}

func (doc *Document) JsonSize() int {
	data, err := json.Marshal(doc)
	if err != nil {
		return -1
	}
	return len(data)
}

func (doc *Document) GetSubKeysByDifficulty(difficulty Difficulty, num int) (result []string) {
	switch difficulty {
	case Easy:
		for i := 0; i < num; i++ {
			result = append(result, subKeyPath(i))
		}
	case Medium:
		current := doc.Size / 2
		sign := 1
		for i := 0; i < num; i++ {
			sign = -sign
			current = current + (sign * i)
			result = append(result, subKeyPath(current+(sign*i)))
		}
	case Hard:
		for i := 1; i <= num; i++ {
			result = append(result, subKeyPath(doc.Size-i))
		}
	case Impossible:
		for i := 0; i < num; i++ {
			result = append(result, subKeyPath(doc.Size+i))
		}
	}
	return
}

type Report struct {
	Errors    int
	Successes int
	Duration  time.Duration
}

func (r Report) AverageRPS() float64 {
	if r.Duration == 0 {
		return 0
	}
	return float64(r.Successes+r.Errors) / (float64(r.Duration) / float64(time.Second))
}

func (r Report) AverageSuccessRPS() float64 {
	if r.Duration == 0 {
		return 0
	}
	return float64(r.Successes) / (float64(r.Duration) / float64(time.Second))
}

func (r Report) ToString() string {
	return fmt.Sprintf("successes: %d, errors: %d, duration: %s, rps: %f, success rps: %f", r.Successes, r.Errors, r.Duration, r.AverageRPS(), r.AverageSuccessRPS())
}
