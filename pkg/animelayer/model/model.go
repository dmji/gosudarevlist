package animelayer_model

type Item struct {
	GUID      string
	Name      string
	Completed bool
}

type DescriptionPoint struct {
	Key   string
	Value string
}

type ItemDescription struct {
	GUID string

	TorrentFilesSize string

	RefImagePreview string
	RefImageCover   string

	UpdatedDate string
	CreatedDate string

	LastCheckedDate string

	Descriptions []DescriptionPoint
}

type Difference struct {
	Name     string
	OldValue string
	NewValue string
}
