package postgresql

import (
	"encoding/csv"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"os"
	"sync"
)

var (
	totalWorker = 100
)

func (postgres *PostgresConnection) RunningWorkerIndonesia() {
	postgres.WorkerProvince()
	postgres.WorkerCity()
	postgres.WorkerDistrict()
	postgres.WorkerVillage()
}

func openCsvFile(nameFile string) (*csv.Reader, *os.File, error) {
	log.Println("=> open csv file")

	f, err := os.Open("./database/csv_indonesia/" + nameFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("%s tidak ditemukan.\n", nameFile)
		}

		return nil, nil, err
	}

	reader := csv.NewReader(f)
	return reader, f, nil
}

func readCsvFilePerLineThenSendToWorker(csvReader *csv.Reader, jobs chan<- []interface{}, wg *sync.WaitGroup) {
	isHeader := true
	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if isHeader {
			isHeader = false
			continue
		}

		rowOrdered := make([]interface{}, 0)
		for _, each := range row {
			rowOrdered = append(rowOrdered, each)
		}

		wg.Add(1)
		jobs <- rowOrdered
	}
	close(jobs)
}
