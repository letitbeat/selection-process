package scoring

// Scorer which holds a Ref map, which is used by the
// Score method.
type Scorer struct {
	// Ref map to compare with.
	Ref map[string]bool
}

// Score which calculates the score from a list of provided tags
// against a Ref set, if all Ref tags are included in the
// provided tags list, then a score of 100 will be returned, otherwise,
// a percentage accordingly, if none of the tags in the provided list match,
// then a 0 value will be returned.
func (s *Scorer) Score(tags []string) float64 {

	// Remove duplicates from input list
	tagsMap := make(map[string]bool)
	for _, tag := range tags {
		tagsMap[tag] = true
	}

	total := len(s.Ref)
	count := 0

	for tag := range tagsMap {
		if _, ok := s.Ref[tag]; ok {
			count++
		}
	}

	if count > 0 {
		return (float64(count) / float64(total)) * 100
	}

	return 0
}
