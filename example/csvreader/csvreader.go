package csvreader

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"insure/insure"
)

type HeaderMap map[string]int

type File struct {
	path      string
	headerMap HeaderMap
	risks     insure.Risks
	csvReader *csv.Reader
	fields    []string
}

func NewFile(filepath string) *File {
	return &File{
		path:  filepath,
		risks: make(insure.Risks, 0),
	}
}

func (file File) get(col string) (value string, err error) {
	index, ok := file.headerMap[col]
	if !ok {
		err = fmt.Errorf("file does not contain column '%s'", col)
		goto end
	}
	if index >= len(file.fields) {
		log.Printf("Only %d fields; not enough to have a %d field '%s'",
			len(file.fields),
			index,
			col,
		)
		goto end
	}
	value = file.fields[index]
end:
	return value, err
}
func (file *File) read() (err error) {
	if file.headerMap == nil {
		err = file.readHeader()
	}
	if err != nil {
		goto end
	}
	file.fields, err = file.csvReader.Read()
end:
	return err
}

// ReadHeader add mapping: Column/property name --> record index
func (file *File) readHeader() (err error) {
	file.headerMap = make(HeaderMap)
	file.fields, err = file.csvReader.Read()
	if err != nil {
		goto end
	}
	for i, v := range file.fields {
		file.headerMap[v] = i
	}
end:
	return err
}

func (file *File) open() (*os.File, error) {
	// Load a csv file.
	f, err := os.Open(file.path)
	if err != nil {
		goto end
	}
	// Create a new reader.
	file.csvReader = csv.NewReader(f)
end:
	return f, err
}

func (file *File) ReadRisks() (_ insure.Risks, err error) {
	var f *os.File
	var recNo int

	f, err = file.open()
	if err != nil {
		goto end
	}
	defer f.Close()
	recNo = 0
	for {
		recNo++
		err = file.read()
		if err == io.EOF {
			err = nil
			goto end
		}
		if err != nil {
			log.Printf("Error reading line %d; %s", recNo, err.Error())
			continue
		}
		err = file.addRisk()
		if err != nil {
			log.Printf("Error adding risk from record #%d; %s", recNo, err.Error())
			continue
		}
	}
end:
	return file.risks, err
}

func (file *File) addRisk() (err error) {
	var field, include, deleted, limit string

	opts := insure.RiskOpts{}
	opts.GUID, err = file.get("Name")
	if err != nil {
		field = "Name"
		goto end
	}
	opts.ID, err = file.get("ID")
	if err != nil {
		field = "ID"
		goto end
	}
	include, err = file.get("Include")
	if err != nil {
		field = "Include"
		goto end
	}
	opts.Include = include == "1"

	deleted, err = file.get("Deleted")
	if err != nil {
		field = "Deleted"
		goto end
	}
	opts.Deleted = deleted == "1"

	limit, err = file.get("Limit")
	if err != nil {
		field = "Limit"
		goto end
	}
	opts.Limit, err = strconv.Atoi(limit)
	if err != nil {
		err = fmt.Errorf("failed to convert 'limit' field with value of '%s' to integer",
			limit,
		)
		goto end
	}
	file.risks = append(file.risks, insure.NewRisk(opts.ID, &opts))
end:
	if err != nil {
		err = fmt.Errorf("failed to get record's '%s' field",
			field,
		)
	}
	return err
}
