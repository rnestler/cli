package models

import (
	"encoding/json"
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

type JobV1Alpha struct {
	ApiVersion string `json:"apiVersion,omitempty" yaml:"apiVersion"`
	Kind       string `json:"kind,omitempty" yaml:"kind"`
	Metadata   struct {
		Name       string      `json:"name,omitempty"`
		Id         string      `json:"id,omitempty"`
		CreateTime json.Number `json:"create_time,omitempty,string" yaml:"create_time,omitempty"`
		UpdateTime json.Number `json:"update_time,omitempty,string" yaml:"update_time,omitempty"`
		StartTime  json.Number `json:"start_time,omitempty,string" yaml:"start_time,omitempty"`
		FinishTime json.Number `json:"finish_time,omitempty,string" yaml:"finish_time,omitempty"`
	} `json:"metadata,omitempty"`

	Status struct {
	}
}

func NewJobV1Alpha(name string) JobV1Alpha {
	j := JobV1Alpha{}

	j.Metadata.Name = name
	j.setApiVersionAndKind()

	return j
}

func NewJobV1AlphaFromJson(data []byte) (*JobV1Alpha, error) {
	j := JobV1Alpha{}

	err := json.Unmarshal(data, &j)

	if err != nil {
		return nil, err
	}

	j.setApiVersionAndKind()

	return &j, nil
}

func NewJobV1AlphaFromYaml(data []byte) (*JobV1Alpha, error) {
	j := JobV1Alpha{}

	err := yaml.UnmarshalStrict(data, &j)

	if err != nil {
		return nil, err
	}

	j.setApiVersionAndKind()

	return &j, nil
}

func (j *JobV1Alpha) setApiVersionAndKind() {
	j.ApiVersion = "v1alpha"
	j.Kind = "Job"
}

func (j *JobV1Alpha) ObjectName() string {
	return fmt.Sprintf("Jobs/%s", j.Metadata.Id)
}

func (j *JobV1Alpha) ToJson() ([]byte, error) {
	return json.Marshal(j)
}

func (j *JobV1Alpha) ToYaml() ([]byte, error) {
	return yaml.Marshal(j)
}
