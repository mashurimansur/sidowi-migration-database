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

func (posgres *PostgresConnection) WorkerVillage() {
	start := time.Now()

	csvReader, csvFile, err := openCsvFile("indonesia_villages.csv")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csvFile.Close()

	jobs := make(chan []interface{}, 0)
	wg := new(sync.WaitGroup)

	go dispatchWorkersVillage(posgres.DB, jobs, wg)
	readCsvFilePerLineThenSendToWorker(csvReader, jobs, wg)

	wg.Wait()

	duration := time.Since(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
}

func dispatchWorkersVillage(db *gorm.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
	for workerIndex := 0; workerIndex <= totalWorker; workerIndex++ {
		go func(workerIndex int, db *gorm.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
			counter := 0

			for job := range jobs {
				doTheJobVillage(workerIndex, counter, db, job)
				wg.Done()
				counter++
			}
		}(workerIndex, db, jobs, wg)
	}
}

func doTheJobVillage(workerIndex, counter int, db *gorm.DB, values []interface{}) {
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
			query := "INSERT INTO id_villages (id, district_id, name) VALUES ($1, $2, $3)"

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
		log.Println("village => worker", workerIndex, "inserted", counter, "data")
	}
}
