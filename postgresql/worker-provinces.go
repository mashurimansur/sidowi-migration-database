package postgresql

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"math"
	"sync"
	"time"
)

func (posgres *PostgresConnection) WorkerProvince() {
	start := time.Now()

	csvReader, csvFile, err := openCsvFile("indonesia_provinces.csv")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csvFile.Close()

	jobs := make(chan []interface{}, 0)
	wg := new(sync.WaitGroup)

	go dispatchWorkersProvince(posgres.DB, jobs, wg)
	readCsvFilePerLineThenSendToWorker(csvReader, jobs, wg)

	wg.Wait()

	duration := time.Since(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
}

func dispatchWorkersProvince(db *gorm.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
	for workerIndex := 0; workerIndex <= totalWorker; workerIndex++ {
		go func(workerIndex int, db *gorm.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
			counter := 0

			for job := range jobs {
				doTheJobProvince(workerIndex, counter, db, job)
				wg.Done()
				counter++
			}
		}(workerIndex, db, jobs, wg)
	}
}

func doTheJobProvince(workerIndex, counter int, db *gorm.DB, values []interface{}) {
	for {
		var outerError error
		func(outerError *error) {
			defer func() {
				if err := recover(); err != nil {
					*outerError = fmt.Errorf("%v", err)
				}
			}()

			// query here
			conn, err := db.DB().Conn(context.Background())
			query := fmt.Sprintf("INSERT INTO id_provinces (id, name) VALUES ($1, $2)")

			_, err = conn.ExecContext(context.Background(), query, values...)
			if err != nil {
				log.Fatal(err.Error())
			}

			err = conn.Close()
			if err != nil {
				log.Fatal(err.Error())
			}
		}(&outerError)
		if outerError == nil {
			break
		}
	}

	if counter%100 == 0 {
		log.Println("provinces => worker", workerIndex, "inserted", counter, "data")
	}
}
