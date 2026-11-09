package v1_test

import (
	"bytes"
	"encoding/json"
	"os"
	"reflect"
	"testing"

	apiv1 "github.com/mroefs1/endpoint_manager/pkg/api/v1"
)

func TestInventoryRoundTrip(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
	}{
		{
			name:     "Full Inventory",
			filepath: "testdata/inventory.json",
		},
		{
			name:     "Failed Collector",
			filepath: "testdata/failed_collector.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test opening input file
			inputFile, err := os.Open(tt.filepath)
			if err != nil {
				t.Fatalf("Failed to open test file %s: %v", tt.filepath, err)
			}
			defer inputFile.Close()

			// Test JSON -> Struct (decoding stage)
			var inv apiv1.Inventory
			decoder := json.NewDecoder(inputFile)
			decoder.DisallowUnknownFields()
			if err := decoder.Decode(&inv); err != nil {
				t.Fatalf("Failed to decode initial file stream: %v", err)
			}

			// Test Struct -> JSON buffer (encoding stage)
			var buf bytes.Buffer
			encoder := json.NewEncoder(&buf)
			if err := encoder.Encode(inv); err != nil {
				t.Fatalf("Failed to encode struct back to stream: %v", err)
			}

			// Reset file pointer to beginning instead of reopening file
			if _, err := inputFile.Seek(0, 0); err != nil {
				t.Fatalf("Failed to rewind input file: %s: %v", tt.filepath, err)
			}

			// Compare structure of original file vs newly generated buffer
			var expectedMap, actualMap map[string]any

			if err := json.NewDecoder(inputFile).Decode(&expectedMap); err != nil {
				t.Fatalf("Failed to decode original file for comparison: %v", err)
			}

			if err := json.NewDecoder(&buf).Decode(&actualMap); err != nil {
				t.Fatalf("Failed to decode generated buffer for comparison: %v", err)
			}

			// Assert deep equality
			if !reflect.DeepEqual(expectedMap, actualMap) {
				want, _ := json.MarshalIndent(expectedMap, "", "  ")
				got, _ := json.MarshalIndent(actualMap, "", "  ")
				t.Errorf("round-trip mismatch for %s\nwant:\n%s\ngot:\n%s", tt.filepath, want, got)
			}

			for _, c := range inv.Collectors {
				section, ok := actualMap[string(c.Name)]
				if !ok {
					t.Errorf("collector %q has no matching section key", c.Name)
					continue
				}
				// Contract rule: a failed collector's section is null, and a
				// successful collector's section is not.
				if c.Error != "" && section != nil {
					t.Errorf("collector %q failed (%q) but its section is not null", c.Name, c.Error)
				}
				if c.Error == "" && section == nil {
					t.Errorf("collector %q succeeded but its section is null", c.Name)
				}
			}
		})
	}
}
