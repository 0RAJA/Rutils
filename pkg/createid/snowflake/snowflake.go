package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

//雪花算法

const (
	Format = "2006-01-02"
)

var node *snowflake.Node

func Init(startTine string, machineID int64) error {
	st, err := time.Parse(Format, startTine)
	if err != nil {
		return err
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return nil
}

func GetID() int64 {
	return node.Generate().Int64()
}
