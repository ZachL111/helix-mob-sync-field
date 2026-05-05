package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 69, Capacity: 92, Latency: 23, Risk: 18, Weight: 8}
	if got := Score(signal); got != 119 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 73, Capacity: 84, Latency: 8, Risk: 13, Weight: 11}
	if got := Score(signal); got != 207 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 85, Capacity: 73, Latency: 9, Risk: 10, Weight: 9}
	if got := Score(signal); got != 220 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
}
