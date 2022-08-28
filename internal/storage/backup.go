package storage

import (
	"encoding/json"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/config"
	"os"
	"time"
)

type BackupConfig struct {
	file     string
	enable   bool
	Periodic bool
}

type producer struct {
	file    *os.File
	encoder *json.Encoder
}

type consumer struct {
	file    *os.File
	decoder *json.Decoder
}

func newProducer(filename string) *producer {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Could not open file for writing")
		return nil
	}
	return &producer{
		file:    file,
		encoder: json.NewEncoder(file),
	}
}

func newConsumer(filename string) *consumer {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Could not open file for reading")
		return nil
	}
	return &consumer{
		file:    file,
		decoder: json.NewDecoder(file),
	}
}

func BackupDo(storage *Storage) {
	p := newProducer(storage.BackupConfig.file)
	if p != nil {
		fmt.Println("Going to backup metrics")
		err := p.encoder.Encode(storage.Metrics)
		if err != nil {
			fmt.Println("Problem during encoding data during dumping")
			fmt.Println(err)
		}
		p.file.Close()
		fmt.Println("Backup dumpening is finished")
	} else {
		fmt.Println("Backup cancelled")
	}
}

func (s *Storage) ConfigureBackup(c *config.ServerConfig) {
	fmt.Println("COnfiguring backup mode")
	s.BackupConfig = BackupConfig{file: c.StoreFile}
	s.BackupConfig.enable = true

	if c.StoreInterval > 0 {
		s.BackupConfig.Periodic = true
		fmt.Printf("Periodic with interval %d\n", c.StoreInterval/1000000000)
		periodicBackup := func(c *config.ServerConfig) {
			ticker := time.NewTicker(c.StoreInterval)
			for {
				<-ticker.C
				BackupDo(s)
			}
		}
		go periodicBackup(c)
	} else {
		fmt.Println("OnDemand")
		s.BackupConfig.Periodic = false
	}
}

func (s *Storage) Restore(c *config.ServerConfig) {
	fmt.Printf("Restoring Data from file: %s\n", c.StoreFile)
	cons := newConsumer(c.StoreFile)
	mmap := metricMap{}
	if cons != nil {
		err := cons.decoder.Decode(&mmap)
		if err != nil {
			fmt.Println("Could not decode Json during Restore Stage")
		} else {
			fmt.Println("Backup Json succesfully decoded")
		}
	} else {
		fmt.Println("Could not find backup file")
	}
	s.Load(mmap)
	fmt.Println("Restoration finished")
}
