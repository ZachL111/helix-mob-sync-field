package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 69, Capacity: 92, Latency: 23, Risk: 18, Weight: 8}, wantScore: 119, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 73, Capacity: 84, Latency: 8, Risk: 13, Weight: 11}, wantScore: 207, wantDecision: "accept"},
		{name: "case_3", signal: Signal{Demand: 85, Capacity: 73, Latency: 9, Risk: 10, Weight: 9}, wantScore: 220, wantDecision: "accept"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
