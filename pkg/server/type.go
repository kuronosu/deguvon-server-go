package server

import (
	"fmt"
	"net/http"

	"github.com/kuronosu/animeflv-api/pkg/db"
	"github.com/kuronosu/animeflv-api/pkg/scrape"
	"github.com/kuronosu/animeflv-api/pkg/utils"
)

// API represents the api
type API struct {
	router    http.Handler
	port      int
	DBManager db.Manager
}

// Server represents the api
type Server interface {
	Router() http.Handler
	Run() error
}

// Router return the api router
func (a *API) Router() http.Handler {
	return LogMiddleware(TrailingSlashes(CaselessMatcher(CorsMiddleware(a.router))))
}

// Run start the server
func (a *API) Run() error {
	utils.InfoLog(fmt.Sprintf("Listen to :%d", a.port))
	return http.ListenAndServe(fmt.Sprint(":", a.port), a.Router())
}

// ErrorResponse rendered in json
type ErrorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

// AnimesResponse rendered anime data in json
type AnimesResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []scrape.Anime `json:"results"`
}

// EpisodesResponse rendered list of episodes data with extra anime data in json
type EpisodesResponse struct {
	AnimeID   int              `json:"animeID"`
	AnimeName string           `json:"animeName"`
	AnimeURL  string           `json:"animeURL"`
	Episodes  []scrape.Episode `json:"episodes"`
}

// EpisodeResponse rendered episode data with extra anime data in json
type EpisodeResponse struct {
	AnimeID   int            `json:"animeID"`
	AnimeName string         `json:"animeName"`
	AnimeURL  string         `json:"animeURL"`
	Episode   scrape.Episode `json:"episode"`
}
