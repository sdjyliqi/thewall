package utils

type cycleType int

var FieldIdle cycleType = 0
var FieldPlanting cycleType = 1
var FieldHarvest cycleType = 2
var FieldWeight cycleType = 3
var FieldFinish cycleType = 4

func (c cycleType) Name() string {
	switch c {
	case FieldIdle:
		return "idle"
	case FieldPlanting:
		return "planting"
	case FieldHarvest:
		return "harvest"
	case FieldWeight:
		return "weight"
	case FieldFinish:
		return "finish"
	}
	return "unknown"
}
