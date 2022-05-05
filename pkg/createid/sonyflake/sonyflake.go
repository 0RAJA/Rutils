package sonyflake

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

const (
	Format = "2006-01-02"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

func Init(startTime string, machineId uint16) error {
	sonyMachineID = machineId
	st, err := time.Parse(Format, startTime)
	if err != nil {
		return err
	}
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return nil
}

func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		return 0, fmt.Errorf("sonyFlake not inited")
	}
	id, err = sonyFlake.NextID()
	return
}
