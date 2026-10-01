package repository

const (
	// SourceLabelFindByName defines a specialized metadata scanning route identifier used inside DAL execution loops
	// to dynamically map and scan rows retrieved from custom, domain-specific query sequences targeting name strings.
	SourceLabelFindByName string = "find_by_name"
)
