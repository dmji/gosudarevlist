package model

type Entity struct {
	Name string
}

type AnimeLayerItem struct {
	GUID      string
	Name      string
	Completed bool
}

type DescriptionPoint struct {
	Key   string
	Value string
}

type AnimeLayerItemDescription struct {
	GUID string

	TorrentFilesSize string
	RefImagePreview  string

	UpdatedDate string
	CreatedDate string

	LastCheckedDate string

	Descriptions []DescriptionPoint
}
