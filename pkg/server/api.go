package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kuronosu/animeflv-api/pkg/db"
	"go.mongodb.org/mongo-driver/mongo"
)

const static = "/static/"

// New create new server
func New(client *mongo.Client, port int) Server {
	fmt.Println(TypeDetailsPath)
	a := &API{DBManager: db.Manager{Client: client}, port: port}
	r := mux.NewRouter()
	r.PathPrefix(static).Handler(http.StripPrefix(static, http.FileServer(http.Dir("."+static))))
	r.HandleFunc("/", a.HandleIndex).Methods(http.MethodGet)
	r.HandleFunc(APIPath, a.HandleAPIIndex).Methods(http.MethodGet)
	r.HandleFunc(TypesPath, a.HandleTypes).Methods(http.MethodGet)
	r.HandleFunc(TypeDetailsPath, a.HandleTypeDetails).Methods(http.MethodGet)
	r.HandleFunc(StatesPath, a.HandleStates).Methods(http.MethodGet)
	r.HandleFunc(GenresPath, a.HandleGenres).Methods(http.MethodGet)
	r.HandleFunc(AnimesPath, a.HandleAnimes).Methods(http.MethodGet)
	r.HandleFunc(DirectoryPath, a.HandleDirectory).Methods(http.MethodGet)
	r.HandleFunc(LatestEpisodesPath, a.HandleLatestEpisodes).Methods(http.MethodGet)
	r.HandleFunc(AnimeDetailsPath, a.HandleAnimeDetails).Methods(http.MethodGet)
	r.HandleFunc(EpisodeListPath, a.HandleEpisodeList).Methods(http.MethodGet)
	r.HandleFunc(EpisodeDetailsPath, a.HandleEpisodeDetails).Methods(http.MethodGet)
	r.HandleFunc(VideoPath, a.HandleEpisodeVideo).Methods(http.MethodGet)
	r.HandleFunc(VideoLangPath, a.HandleEpisodeVideo).Methods(http.MethodGet)
	a.router = r
	return a
}
