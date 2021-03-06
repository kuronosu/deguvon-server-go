package server

// APIPath uri endpoint
const APIPath = "/api"

// TypesPath uri endpoint
const TypesPath = APIPath + "/types"

// TypeDetailsPath uri endpoint
const TypeDetailsPath = TypesPath + "/{id:[0-9]+}"

// StatesPath uri endpoint
const StatesPath = APIPath + "/states"

// StateDetailsPath uri endpoint
const StateDetailsPath = StatesPath + "/{id:[0-9]+}"

// GenresPath uri endpoint
const GenresPath = APIPath + "/genres"

// GenreDetailsPath uri endpoint
const GenreDetailsPath = GenresPath + "/{id:[0-9]+}"

// AnimesPath uri endpoint
const AnimesPath = APIPath + "/animes"

// LatestEpisodesPath uri endpoint
const LatestEpisodesPath = APIPath + "/latest"

// AnimeDetailsPath uri endpoint
const AnimeDetailsPath = AnimesPath + "/{flvid:[0-9]+}"

// SearchAnimePath uri endpoint
const SearchAnimePath = AnimesPath + "/search"

// EpisodeListPath uri endpoint
const EpisodeListPath = AnimeDetailsPath + "/episodes"

// EpisodeDetailsPath uri endpoint
const EpisodeDetailsPath = EpisodeListPath + `/{eNumber}`

// VideoPath uri endpoint
const VideoPath = EpisodeDetailsPath + `/{server}`

// VideoLangPath uri endpoint
const VideoLangPath = VideoPath + `/{lang}`

// DirectoryPath uri endpoint
const DirectoryPath = APIPath + "/directory"

// ScreenshotsPath image endpoint
const ScreenshotsPath = `/screenshots/{a:[0-9]+}/{e:[0-9]+}/th_3.jpg`

// CoversPath image endpoint
const CoversPath = `/uploads/animes/covers/{a:[0-9]+}.jpg`

// BannersPath image endpoint
const BannersPath = `/uploads/animes/banners/{a:[0-9]+}.jpg`

// ThumbsPath image endpoint
const ThumbsPath = `/uploads/animes/thumbs/{a:[0-9]+}.jpg`

// AllPaths array with all paths (uris)
var AllPaths = []string{APIPath, TypesPath, StatesPath, GenresPath, AnimesPath, LatestEpisodesPath, DirectoryPath}

// AllPathsWithoutIndex array with all paths (uris) except the index
var AllPathsWithoutIndex = []string{TypesPath, StatesPath, GenresPath, AnimesPath, LatestEpisodesPath, DirectoryPath}
