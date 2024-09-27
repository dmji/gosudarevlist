package animelayer_model

type Item struct {
	Identifier  string
	Title       string
	IsCompleted bool
}

type DescriptionNote struct {
	Name string
	Text string
}

type Description struct {
	Identifier string

	TorrentFilesSize string

	RefImagePreview string
	RefImageCover   string

	UpdatedDate string
	CreatedDate string

	LastCheckedDate string

	Notes []DescriptionNote
}

type Difference struct {
	Name     string
	OldValue string
	NewValue string
}
