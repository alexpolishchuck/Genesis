package main

import (
	"encoding/csv"
	"log"
	"os"
	"sync"
)

type file_handler struct {
	filepath string
	rwlock   sync.RWMutex
}

func (fh file_handler) save_email(email string) {

	f, err := os.OpenFile(fh.filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
		return
	}

	writer := csv.NewWriter(f)

	fh.rwlock.Lock()
	err = writer.Write([]string{email})
	writer.Flush()
	fh.rwlock.Unlock()

	if err != nil {
		log.Fatal(err)
		return
	}
}

func (fh file_handler) read_all() [][]string {
	f, err := os.OpenFile(fh.filepath, os.O_RDONLY, 0600)
	if err != nil {
		log.Fatal(err)
		return [][]string{}
	}

	defer f.Close()

	writer := csv.NewReader(f)

	fh.rwlock.RLock()
	res, err := writer.ReadAll()
	fh.rwlock.RUnlock()

	return res
}

func (fh file_handler) delete_email(email string) {

	emails := fh.read_all()

	fh.rwlock.Lock()

	f, err := os.OpenFile(fh.filepath, os.O_WRONLY|os.O_TRUNC, 0600)

	if err != nil {
		log.Fatal(err)
		return
	}

	writer := csv.NewWriter(f)

	for _, i := range emails {
		for _, j := range i {
			if j != email {

				err = writer.Write([]string{j})

			}
		}
	}

	writer.Flush()

	fh.rwlock.Unlock()

	if err != nil {
		log.Fatal(err)
		return
	}
}
