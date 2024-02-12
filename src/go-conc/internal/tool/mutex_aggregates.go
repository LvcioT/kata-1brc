package tool

import (
	"lvciot/go-conc/internal/model"
	"sort"
	"sync"
)

type MutexAggregates struct {
	mutex      sync.Mutex
	aggregates map[string]*model.StationAggregate
}

func NewMutexAggregates() (ma MutexAggregates) {
	ma.aggregates = make(map[string]*model.StationAggregate)

	return
}

func (ma *MutexAggregates) AddDetection(d model.Detection) {
	ma.mutex.Lock()
	defer ma.mutex.Unlock()

	a, exist := ma.aggregates[d.Station]

	if exist {
		a.AddDetection(d)
	} else {
		a := model.NewStationAggregateFromDetection(d)
		ma.aggregates[d.Station] = &a
	}
}

func (ma *MutexAggregates) SortedRows() []string {
	totalStations := len(ma.aggregates)

	stations := make([]string, totalStations)
	aggregateRows := make([]string, totalStations)

	j := 0
	for station, _ := range ma.aggregates {
		stations[j] = station
		j++
	}

	sort.Strings(stations)

	for j, station := range stations {
		aggregate := ma.aggregates[station]
		aggregateRows[j] = aggregate.String()
	}

	return aggregateRows
}
