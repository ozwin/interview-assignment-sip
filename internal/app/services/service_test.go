package services

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/ozwin/interview-assignment-sip/internal/app/models"
)

func Test_readTransactionsFromFile(t *testing.T) {

	fileName := "sample_registrations"
	createSampledataFile(t, fileName)
	defer deleteSampleDataFile(t, fileName)

	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Transactions
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Valid Data",
			args: args{fileName: fileName},
			want: &models.Transactions{
				{
					Address:      "1",
					TenantId:     "1",
					Uri:          "1",
					Contact:      "1",
					Path:         []string{"a", "b"},
					Source:       "1",
					Target:       "1",
					UserAgent:    "1",
					RawUserAgent: "1",
					Created:      "2016-12-12T22:40:40.764Z",
					LineId:       "1",
				}, {
					Address:      "2",
					TenantId:     "2",
					Uri:          "2",
					Contact:      "2",
					Path:         []string{"a"},
					Source:       "2",
					Target:       "2",
					UserAgent:    "2",
					RawUserAgent: "2",
					Created:      "2018-12-12T22:40:40.764Z",
					LineId:       "2",
				},
			},
			wantErr: false,
		}, {
			name:    "Invalid file name",
			args:    args{fileName: fmt.Sprintf("%v.txt.csv.pdf", fileName)},
			want:    nil,
			wantErr: true,
		},
		// {
		// 	name:    "Main file name",
		// 	args:    args{fileName: "../../data/sample_registrations"},
		// 	want:    nil,
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readTransactionsFromFile(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("readTransactionsFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readTransactionsFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createSampledataFile(t *testing.T, fileName string) {
	//ToDo: check later for this file existence to avoid redundency
	testFile, err := os.Create(fileName)
	if err != nil {
		t.Fatalf("error creating temp file for testing %v", err)
	}
	defer testFile.Close()
	data := `{
  	"addressOfRecord": "1",
  	"tenantId": "1",
  	"uri": "1",
  	"contact": "1",
  	"path": [
  	  "a","b"
  	],
  	"source": "1",
  	"target": "1",
  	"userAgent": "1",
  	"rawUserAgent": "1",
  	"created": "2016-12-12T22:40:40.764Z",
  	"lineId": "1"
	}{
	  "addressOfRecord": "2",
	  "tenantId": "2",
	  "uri": "2",
	  "contact": "2",
	  "path": [
	    "a"
	  ],
	  "source": "2",
	  "target": "2",
	  "userAgent": "2",
	  "rawUserAgent": "2",
	  "created": "2018-12-12T22:40:40.764Z",
  	  "lineId": "2"
  }`
	if _, err := testFile.Write([]byte(data)); err != nil {
		t.Fatalf("error writing to temp file: %v", err)
	}
}

func deleteSampleDataFile(t *testing.T, fileName string) {
	time.Sleep(1 * time.Second)
	if err := os.Remove(fileName); err != nil {
		t.Logf("failed to remove the sample data file: %v", err.Error())
	}
}
