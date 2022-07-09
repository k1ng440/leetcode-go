package designundergroundsystem

type Checkin struct {
	StationName string
	CheckinTime int
}

type History struct {
	CheckIn  int
	CheckOut int
}

type UndergroundSystem struct {
	customer map[int]*Checkin
	history  map[[2]string][]*History
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		customer: make(map[int]*Checkin),
		history:  make(map[[2]string][]*History),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.customer[id] = &Checkin{
		StationName: stationName,
		CheckinTime: t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	mKey := [2]string{this.customer[id].StationName, stationName}

	this.history[mKey] = append(this.history[mKey], &History{
		CheckIn:  this.customer[id].CheckinTime,
		CheckOut: t,
	})
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	mKey := [2]string{startStation, endStation}
	totalTime := 0
	for _, h := range this.history[mKey] {
		totalTime += h.CheckOut - h.CheckIn
	}

	return float64(totalTime) / float64(len(this.history[mKey]))
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
