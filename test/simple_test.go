package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/butterfly-ci/syntax-generator/generated/proto"
	"github.com/gogo/protobuf/jsonpb"
	"sigs.k8s.io/yaml"
)

func TestSimple(t *testing.T) {
	tests := []struct {
		name string
		file string
	}{
		{
			name: "simple.yaml",
			file: "yaml/simple.yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dat, err := ioutil.ReadFile(tt.file)
			if err != nil {
				log.Fatal("dodgy...")
			}
			pipeline := &proto.Main{}
			a := &jsonpb.Unmarshaler{
				AllowUnknownFields: true,
			}

			jsonBytes, err := yaml.YAMLToJSON(dat)
			if err != nil {
				log.Fatal("oh no")
			}
			r := bytes.NewReader(jsonBytes)
			err = a.Unmarshal(r, pipeline)
			if err != nil {
				log.Fatalf("oh no %v", err)
			}
			fmt.Printf("%v\n", pipeline.Init.BaseImage)
			fmt.Printf("%v\n", pipeline.Init.Env)
			fmt.Printf("%v\n", pipeline.Init.Secrets)
			fmt.Printf("%v\n", pipeline.Init.Commands)
		})
	}
}
