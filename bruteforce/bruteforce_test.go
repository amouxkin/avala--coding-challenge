package bruteforce

import (
	"avala/common"
	"gorm.io/gorm"
	"math"
	"sync"
)

const batchSize = 0x10_000
const databaseBatch = 0x1_000
const upperLimit = 0xfffff

// createBatch processes batches of data that will be inserted to the database.
func createBatch(offset int, wg *sync.WaitGroup, db *gorm.DB) {

	start := offset * batchSize
	end := start + batchSize

	hexList := make([]common.HexCounter, batchSize)

	if end > upperLimit {
		newLimit := end - upperLimit
		hexList = hexList[:batchSize-newLimit+1]
	}

	for i := range hexList {
		currentHex := start + i

		if currentHex > upperLimit {
			break
		}

		hexList[i] = common.HexCounter{
			Hex:   currentHex,
			Count: 0,
		}
	}

	db.Begin().CreateInBatches(hexList, databaseBatch).Commit()
	wg.Done()
}

// Seed adds all the possible data values to the db.
func Seed() {
	iteration := int(math.Ceil(float64(upperLimit) / batchSize))

	wg := sync.WaitGroup{}

	wg.Add(iteration + 1)
	ch := make(chan int, iteration)

	for thread := 0; thread < 16; thread++ {
		db := common.GetDb()

		go func(ch <-chan int) {
			for i := range ch {
				createBatch(i, &wg, db)
			}
		}(ch)
	}

	go func(ch chan<- int) {
		defer func() {
			close(ch)
		}()

		for i := 0; i < iteration; i++ {
			ch <- i
		}
		wg.Done()
	}(ch)

	wg.Wait()
}
