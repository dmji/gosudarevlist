package animelayer

import "strconv"

const BaseUrl = "https://animelayer.ru"

func FormatUrlToItemsPage(category string, iPage int) string {
	return BaseUrl + category + "?page=" + strconv.FormatInt(int64(iPage), 10)
}

func FormatUrlToItem(guid string) string {
	return BaseUrl + "/torrent/" + guid
}
func FormatUrlToItemDownload(guid string) string {
	return BaseUrl + "/torrent/" + guid + "/download"
}
