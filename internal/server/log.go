package server

import (
	"fmt"
	"log"
	"sync"
)

type Record struct {
	Offset uint64 `json:"offset"`
	Value  []byte `json:"value"`
}

type Log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

func (c *Log) StoreRecord(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	log.Printf("Storing record %v ", record)
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read(offset uint64) (Record, error) {
	if offset > uint64(len(c.records)) {
		return Record{}, ErrorOffsetNotFound
	}

	return c.records[offset], nil
}

var ErrorOffsetNotFound = fmt.Errorf("offset not found")
