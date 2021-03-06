package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kuronosu/animeflv-api/pkg/db"
)

// InternalError make an json response with error message
func InternalError(w http.ResponseWriter, err string) {
	JSONResponse(w, ErrorResponse{err, http.StatusInternalServerError}, http.StatusInternalServerError)
}

// JSONResponse create an http response in json format
func JSONResponse(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)
}

func formatOptions(options db.Options) string {
	if options.SortField == "_id" && options.SortValue == 1 {
		return ""
	}
	if options.SortField == "_id" {
		options.SortField = "flvid"
	}
	if options.SortValue >= 0 {
		return "&order=" + options.SortField
	}
	return "&order=-" + options.SortField
}

func assembleAnimesPageLink(result db.PaginatedAnimeResult, next bool, options db.Options) *string {
	newURI := AnimesPath + "?page=%d"
	if next && result.Page < result.TotalPages {
		newURI = fmt.Sprintf(newURI, result.Page+1)
	} else if !next && result.Page >= 2 {
		newURI = fmt.Sprintf(newURI, result.Page-1)
	} else {
		return nil
	}
	newURI = newURI + formatOptions(options)
	return &newURI
}

var animeOrderOptions = [...]string{"name", "state", "score", "votes",
	"-name", "-state", "-score", "-votes"}

func validSortField(field string) (string, int) {
	if field == "flvid" {
		return "_id", 1
	} else if field == "-flvid" {
		return "_id", -1
	}
	for index, op := range animeOrderOptions {
		if op == field {
			if index >= 4 {
				return op[1:], -1
			}
			return op, 1
		}
	}
	return "_id", 1
}

func trimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
