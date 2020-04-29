package visits

import (
	"fmt"
)

// Visitor struct
type Visitor struct {
	IP 			string
	Counter	uint64
}

// Visitors struct.
type Visitors struct {
	All		[]Visitor
	Total	uint64
}


// AddNewVisitor visitor to memory.
func (v *Visitors) AddNewVisitor(ip string) {
	v.All = append(v.All, Visitor{
		IP: ip,
		Counter: 1,
	})
	v.Total++
}

// IterateVisitByIP makes Visitor.Count++.
func (v *Visitors) IterateVisitByIP(ip string) error {
	v.Total++

	for i := 0; i < len(v.All); i++ {
		if v.All[i].IP == ip {
			v.All[i].Counter++
			return nil
		}
	}
	return fmt.Errorf("Visitor with such ip haven't been found")
}

// GetByIP returns visitor's count of visits.
func (v Visitors) GetByIP(ip string) (Visitor, bool) {
	for _, value := range v.All {
		if value.IP == ip {
			return value, true
		}
	}
	return Visitor{}, false
}

// GetTotalVisits returns total amount of visits of this page.
func (v Visitors) GetTotalVisits() uint64 {
	return v.Total
}

// GetUniqueVisitors returns count of all unique visitors.
func (v Visitors) GetUniqueVisitors() uint64 {
	return uint64(len(v.All))
}