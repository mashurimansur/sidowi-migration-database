package postgresql

import (
	"context"
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

func (posgres *PostgresConnection) WorkerDistrict() {
	start := time.Now()

	csvReader, csvFile, err := openCsvFile("indonesia_districts.csv")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csvFile.Close()

	jobs := make(chan []interface{}, 0)
	wg := new(sync.WaitGroup)

	go dispatchWorkersDistrict(posgres.DB, jobs, wg)
	readCsvFilePerLineThenSendToWorker(csvReader, jobs, wg)

	wg.Wait()

	duration := time.Since(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
}

func dispatchWorkersDistrict(db *gorm.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
	for workerIndex := 0; workerIndex <= totalWorker; workerIndex++ {
		go func(workerIndex int, db *gorm.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
			counter := 0

			for job := range jobs {
				doTheJobDistrict(workerIndex, counter, db, job)
				wg.Done()
				counter++
			}
		}(workerIndex, db, jobs, wg)
	}
}

func doTheJobDistrict(workerIndex, counter int, db *gorm.DB, values []interface{}) {
	for {
		var outerError error
		func(outerError *error) {
			defer func() {
				if err := recover(); err != nil {
					*outerError = fmt.Errorf("%v", err)
				}
			}()

			// query here
			conn, _ := db.DB().Conn(context.Background())
			query := "INSERT INTO id_districts (id, city_id, name) VALUES ($1, $2, $3)"

			_, err := conn.ExecContext(context.Background(), query, values...)
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
		log.Println("district => worker", workerIndex, "inserted", counter, "data")
	}
}
