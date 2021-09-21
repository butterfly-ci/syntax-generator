package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/butterfly-ci/syntax-generator/generated/test"
	"google.golang.org/protobuf/encoding/protojson"
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
			pipeline := &test.SearchRequest{}
			jsonBytes, err := yaml.YAMLToJSON(dat)
			if err != nil {
				log.Fatal("oh no")
			}
			protojson.Unmarshal(jsonBytes, pipeline)
			fmt.Printf("%v", pipeline)
		})
	}
}
