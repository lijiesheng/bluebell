package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

func Init(startTime string, machineId int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime) // string 转换为 time
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000 / 1000
	node, err = sf.NewNode(machineId)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
