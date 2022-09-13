package memory

import (
	"encoding/json"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/config"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
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
		logging.Debug("Could not open file for writing")
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
		logging.Warn("Could not open file for reading")
		return nil
	}
	return &consumer{
		file:    file,
		decoder: json.NewDecoder(file),
	}
}

func BackupDo(storage *MemStorage) {
	p := newProducer(storage.BackupConfig.file)
	if p != nil {
		logging.Debug("Going to backup metrics")
		err := p.encoder.Encode(storage.Metrics)
		if err != nil {
			logging.Warn("Problem during encoding data during dumping", err.Error())
		}
		p.file.Close()
		logging.Debug("Backup dumpening is finished")
	} else {
		logging.Warn("Backup failed")
	}
}

func (s *MemStorage) ConfigureBackup(c *config.ServerConfig) {
	logging.Debug("Configuring backup mode:")
	s.BackupConfig = BackupConfig{file: c.StoreFile}
	s.BackupConfig.enable = true

	if c.StoreInterval > 0 {
		s.BackupConfig.Periodic = true
		logging.Debug("Periodic with interval %d\n", c.StoreInterval/1000000000)
		periodicBackup := func(c *config.ServerConfig) {
			ticker := time.NewTicker(c.StoreInterval)
			for {
				<-ticker.C
				BackupDo(s)
			}
		}
		go periodicBackup(c)
	} else {
		logging.Debug("OnDemand")
		s.BackupConfig.Periodic = false
	}
}

func (s *MemStorage) Restore(c *config.ServerConfig) {
	fmt.Printf("Restoring Data from file: %s\n", c.StoreFile)
	cons := newConsumer(c.StoreFile)
	mmap := metric.MetricMap{}
	if cons != nil {
		err := cons.decoder.Decode(&mmap)
		if err != nil {
			logging.Warn("Could not decode Json during Restore Stage")
		} else {
			logging.Debug("Backup Json succesfully decoded")
		}
	} else {
		logging.Warn("Could not find backup file")
	}
	s.Load(mmap)
	logging.Debug("Restoration finished")
}
