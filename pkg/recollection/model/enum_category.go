package model

type Category string

type categories struct {
	Anime       Category
	AnimeHentai Category
	Manga       Category
	MangaHentai Category
	Music       Category
	Dorama      Category
	All         Category
}

// Categories - object to emulate enum class
var Categories = categories{
	Anime:       "anime",
	AnimeHentai: "anime_hentai",
	Manga:       "manga",
	MangaHentai: "manga_henai",
	Music:       "music",
	Dorama:      "dorama",
	All:         "",
}
