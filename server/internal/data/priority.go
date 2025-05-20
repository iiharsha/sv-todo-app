package data

import (
	"encoding/json"
	"fmt"
)

type PriorityLevel int8

const (
	Low PriorityLevel = iota
	Med
	High
)

var priorityNames = [...]string{"low", "medium", "high"}
var priorityMap = map[string]PriorityLevel{
	"low":    Low,
	"medium": Med,
	"high":   High,
}

func (p PriorityLevel) String() string {
	if int(p) < 0 {
		return "low"
	}

	if int(p) >= len(priorityNames) {
		return "high"
	}

	return priorityNames[p]
}

func (p PriorityLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *PriorityLevel) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	level, ok := priorityMap[s]
	if !ok {
		return fmt.Errorf("invalid priority level: %s", s)
	}
	*p = level
	return nil
}
