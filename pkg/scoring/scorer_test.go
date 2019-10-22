package scoring

import "testing"

func TestNoMatch(t *testing.T) {

	ref := make(map[string]bool)

	ref["tag1"] = true
	ref["tag2"] = true

	scorer := Scorer{ref}

	tags := []string{"tag3"}

	score := scorer.Score(tags)
	wanted := float64(0)
	if score != wanted {
		t.Errorf("got: %f, wanted: %f", score, wanted)
	}
}

func TestCompleteMatch(t *testing.T) {

	ref := make(map[string]bool)

	ref["tag1"] = true
	ref["tag2"] = true

	scorer := Scorer{ref}

	tags := []string{"tag1", "tag2"}

	score := scorer.Score(tags)
	wanted := 100.0
	if score != wanted {
		t.Errorf("got: %f, wanted: %f", score, wanted)
	}
}

func TestCompleteMatchWithDuplicates(t *testing.T) {

	ref := make(map[string]bool)

	ref["tag1"] = true
	ref["tag2"] = true

	scorer := Scorer{ref}

	tags := []string{"tag1", "tag1", "tag2", "tag2"}

	score := scorer.Score(tags)
	wanted := float64(100)
	if score != wanted {
		t.Errorf("got: %f, wanted: %f", score, wanted)
	}
}

func TestHalfMatches(t *testing.T) {

	ref := make(map[string]bool)

	ref["tag1"] = true
	ref["tag2"] = true

	scorer := Scorer{ref}

	tags := []string{"tag1", "tag3"}

	score := scorer.Score(tags)
	wanted := float64(50)
	if score != wanted {
		t.Errorf("got: %f, wanted: %f", score, wanted)
	}
}

func TestThirdMatches(t *testing.T) {

	ref := make(map[string]bool)

	ref["tag1"] = true
	ref["tag2"] = true
	ref["tag3"] = true

	scorer := Scorer{ref}

	tags := []string{"tag1", "tag4"}

	score := scorer.Score(tags)
	wanted := (float64(1) / float64(3)) * 100
	if score != wanted {
		t.Errorf("got: %f, wanted: %f", score, wanted)
	}
}

func TestCompleteInputTags(t *testing.T) {
	ref := make(map[string]bool)

	ref["Log | Packer"] = true
	ref["BusinessDev | Outbound"] = true
	ref["Event | Street Marketing"] = true

	scorer := Scorer{ref}

	tags := []string{"Log | Packer", "BusinessDev | Outbound", "Event | Street Marketing"}

	score := scorer.Score(tags)
	wanted := float64(100)
	if score != wanted {
		t.Errorf("got: %f, wanted: %f", score, wanted)
	}
}

func TestRefTagsLessThanInput(t *testing.T) {
	ref := make(map[string]bool)

	ref["Log | Packer"] = true
	scorer := Scorer{ref}

	tags := []string{"Log | Packer", "BusinessDev | Outbound", "Event | Street Marketing"}

	score := scorer.Score(tags)
	wanted := float64(100)
	if score != wanted {
		t.Errorf("got: %f, wanted: %f", score, wanted)
	}
}
