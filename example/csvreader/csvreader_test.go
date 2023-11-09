package csvreader_test

import (
	"os"
	"reflect"
	"testing"

	"insure/csvreader"
	"insure/insure"
)

var data = `
Name,ID,Include,Deleted,Limit
Risk1,1,1,0,25
Risk2,2,1,0,22
Risk3,3,1,0,21
Risk4,4,1,0,15
`

func getRisks() insure.Risks {
	risks := make(insure.Risks, 4)
	risks[0] = insure.NewRisk("1", &insure.RiskOpts{
		GUID:    "Risk1",
		Include: true,
		Deleted: false,
		Limit:   25,
	})
	risks[1] = insure.NewRisk("2", &insure.RiskOpts{
		GUID:    "Risk2",
		Include: true,
		Deleted: false,
		Limit:   22,
	})
	risks[2] = insure.NewRisk("3", &insure.RiskOpts{
		GUID:    "Risk3",
		Include: true,
		Deleted: false,
		Limit:   21,
	})
	risks[3] = insure.NewRisk("4", &insure.RiskOpts{
		GUID:    "Risk4",
		Include: true,
		Deleted: false,
		Limit:   15,
	})
	return risks
}

func makeDatafile() (string, error) {
	file, err := os.CreateTemp("/tmp", "csvdata-")
	if err != nil {
		goto end
	}
	_, err = file.WriteString(data)
	if err != nil {
		goto end
	}
	err = file.Close()
	if err != nil {
		goto end
	}
end:
	return file.Name(), err
}

func TestFile_ReadRisks(t *testing.T) {
	filepath, err := makeDatafile()
	if err != nil {
		t.Fatal("unable to create temp file like '/tmp/csvdata-*'")
	}
	defer os.Remove(filepath)
	reader := csvreader.NewFile(filepath)
	tests := []struct {
		name    string
		want    insure.Risks
		wantErr bool
	}{
		{
			name: "Read Risks",
			want: getRisks(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			line := insure.NewLine("Line1", &insure.LineOpts{})
			err = reader.ReadRisks(line)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadRisks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := line.Risks()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadRisks() got = %v, want %v", got, tt.want)
			}
		})
	}
}
