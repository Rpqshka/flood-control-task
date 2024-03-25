package flood_control_task

type Flood struct {
	Id   int64 `bson:"id, omitempty"`
	Time int64 `bson:"time, omitempty"`
}
