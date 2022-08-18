package examples

// todo 不知道是啥意思
type Feeder struct {
	collection string
	database string
	isDrop bool
	total int
}

type Model struct {
	ID string `json:"_id" bson:"_id"`
	Name string
	Description string
	Year int
}

type Task struct {
	For string `json:"for" bson:"for"`
	MinutesUsed int `json:"minutes_used" bson:"minutes_used"`
}

type Rebot struct {
	ID string `json:"_id" bson:"_id"`
	ModelID string `json:"modelId,omitempty" bson:"modelId,omitempty"`
	Notes string `json:"notes" bson:"notes"`
	BatteryPct float32 `json:"batteryPct,omitempty" bson:"batteryPct, omitempty"`
	Tasks      []Task  `json:"tasks" bson:"tasks"`
}

// NewFeeder establish seeding parameters
func NewFeeder() *Feeder {
	return &Feeder{isDrop: false, total: 1000}
}

// SetCollection set collection
func (f *Feeder) SetCollection(collection string) {
	f.collection = collection
}

// SetDatabase set database
func (f *Feeder) SetDatabase(database string) {
	f.database = database
}

// SetIsDrop set isDrop
func (f *Feeder) SetIsDrop(isDrop bool) {
	f.isDrop = isDrop
}

// SetTotal set total
func (f *Feeder) SetTotal(total int) {
	f.total = total
}

