package expose_header_utils

import (
	"encoding/json"
	"net/http"
)

func WriterExposeHeader(w http.ResponseWriter, name string, value any) error {
	s := ""
	switch v := value.(type) {
	case string:
		s = v
	default:
		b, err := json.Marshal(value)
		if err != nil {
			return err
		}
		s = string(b)
	}

	w.Header().Add("Access-Control-Expose-Headers", name)
	w.Header().Add(name, s)
	return nil
}
