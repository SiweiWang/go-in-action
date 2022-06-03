package search

//default Matcher implements the default matcher
type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

//search implements the behavior for the default matchers.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
