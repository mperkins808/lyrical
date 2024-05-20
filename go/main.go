package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	LYRICS_ENDPOINT = "https://api.lyrics.ovh/v1"
)

type SortedTrack struct {
	Artists []string `json:"artists"`
	Album   string   `json:"album"`
	Name    string   `json:"name"`
	Lyrics  string   `json:"lyrics,omitempty"`
}

type LyricResp struct {
	Lyrics string `json:"lyrics"`
}

func main() {
	ctx := context.Background()

	artist := flag.String("artist", "", "Name of the artist")
	dir := flag.String("directory", "", "directory to save the lyrics")
	i := flag.String("id", "", "Spotify app id")
	secret := flag.String("secret", "", "Spotify app secret")

	flag.StringVar(artist, "a", "", "Name of the artist")
	flag.StringVar(dir, "d", "", "Directory to save the lyrics (shorthand)")
	flag.StringVar(i, "i", "", "Spotify app id (shorthand)")
	flag.StringVar(secret, "s", "", "Spotify app secret (shorthand)")
	flag.Parse()

	if *dir == "" {
		flag.Usage()
		log.Fatal("must supply a directory to store the lyrics")

	}

	id, sec := getSpotifyAppCreds(*i, *secret)

	client, err := createSpotifyClient(ctx, id, sec)
	if err != nil {
		log.Fatal(err)
	}

	artistObj, err := getArtistObj(ctx, client, *artist)
	if err != nil {
		log.Fatal(err)
	}

	albums, err := getArtistAlbums(ctx, client, artistObj)
	if err != nil {
		log.Fatal(err)
	}

	sortedTracks, err := getAlbumsSongs(ctx, client, albums)
	if err != nil {
		log.Fatal(err)
	}

	history := make([]string, 0)
	existing, err := LoadSongs(*dir)
	if err == nil {
		for _, track := range existing {
			history = append(history, track.Name)
		}
	}

	sortedTracks, _, err = fetchLyrics(sortedTracks, history)
	if err != nil {
		log.Fatal(err)
	}
	SaveJson(*artist, *dir, sortedTracks)

}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func fetchLyrics(sortedTracks []SortedTrack, history []string) ([]SortedTrack, []string, error) {
	// newSorted := sortedTracks
	client := http.Client{}
	newHistory := history
	newTracks := make([]SortedTrack, 0)
	for _, track := range sortedTracks {
		if contains(newHistory, track.Name) {
			log.Printf("skipping lyric fetch for %s", track.Name)
			continue
		}

		for _, artist := range track.Artists {
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/%v/%v", LYRICS_ENDPOINT, artist, track.Name), nil)
			if err != nil {
				return nil, nil, fmt.Errorf("could not create client to find lyrics for %v", track.Name)
			}
			resp, err := client.Do(req)
			if err != nil {
				return nil, nil, fmt.Errorf("request failed to find lyrics for %v", track.Name)
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				continue
			}

			log.Printf("fetched the lyrics for %v", track.Name)

			var lyricsResp LyricResp
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("failed to read response from song %v: %v", track.Name, err)
				continue
			}

			err = json.Unmarshal(body, &lyricsResp)
			if err != nil {
				log.Printf("failed to read response from song %v: %v", track.Name, err)
				continue
			}

			track.Lyrics = lyricsResp.Lyrics
			newTracks = append(newTracks, track)

			newHistory = append(newHistory, track.Name)
			break
		}

	}
	return newTracks, newHistory, nil
}

func getAlbumTracks(ctx context.Context, client *spotify.Client, album spotify.SimpleAlbum) ([]spotify.SimpleTrack, error) {
	log.Printf("fetching tracks for album %v", album.Name)
	pager, err := client.GetAlbumTracks(ctx, album.ID)
	if err != nil {
		return nil, err
	}
	tracks := make([]spotify.SimpleTrack, 0)
	tracks = append(tracks, pager.Tracks...)
	for page := 1; ; page++ {
		err = client.NextPage(ctx, pager)
		if err == spotify.ErrNoMorePages {
			tracks = append(tracks, pager.Tracks...)
			return tracks, nil
		}
		if err != nil {
			return tracks, err
		}

	}
}

func sortAlbumTracks(tracks []spotify.SimpleTrack, album spotify.SimpleAlbum) []SortedTrack {
	sorted := make([]SortedTrack, 0)

	for _, track := range tracks {
		artists := make([]string, 0)
		for _, artist := range track.Artists {
			artists = append(artists, artist.Name)
		}

		tmp := SortedTrack{
			Album:   album.Name,
			Name:    track.Name,
			Artists: artists,
		}
		sorted = append(sorted, tmp)
	}
	return sorted
}

func getAlbumsSongs(ctx context.Context, client *spotify.Client, albums []spotify.SimpleAlbum) ([]SortedTrack, error) {
	sortedTracks := make([]SortedTrack, 0)
	for _, alb := range albums {
		tracks, err := getAlbumTracks(ctx, client, alb)
		if err != nil {
			return nil, err
		}
		sortedTracks = append(sortedTracks, sortAlbumTracks(tracks, alb)...)

	}
	return sortedTracks, nil
}

func getArtistAlbums(ctx context.Context, client *spotify.Client, artist *spotify.FullArtist) ([]spotify.SimpleAlbum, error) {

	log.Printf("fetching albums for %v", artist.Name)
	pager, err := client.GetArtistAlbums(ctx, artist.ID, []spotify.AlbumType{spotify.AlbumTypeAlbum, spotify.AlbumTypeSingle})
	if err != nil {
		log.Fatal(err)

	}
	albums := make([]spotify.SimpleAlbum, 0)
	albums = append(albums, pager.Albums...)
	for page := 1; ; page++ {
		err = client.NextPage(ctx, pager)
		if err == spotify.ErrNoMorePages {
			albums = append(albums, pager.Albums...)
			return albums, nil
		}
		if err != nil {
			return albums, err
		}

	}

	// SaveJson(albs)
}

func getArtistObj(ctx context.Context, client *spotify.Client, artist string) (*spotify.FullArtist, error) {
	log.Printf("searching for %v", artist)
	result, err := client.Search(ctx, artist, spotify.SearchTypeArtist)

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Artists.Artists {
		if item.Name == artist {
			log.Printf("found %v", artist)
			return &item, nil
		}
	}
	return nil, fmt.Errorf("failed to find artist %v", artist)
}

func getSpotifyAppCreds(i string, s string) (string, string) {

	if i != "" || s != "" {
		return i, s
	}

	id := os.Getenv("SPOTIFY_CLIENT_ID")
	secret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	if id == "" && i == "" {
		log.Fatal("SPOTIFY_CLIENT_ID is not set")
	}
	if secret == "" && s == "" {
		log.Fatal("SPOTIFY_CLIENT_SECRET is not set")
	}

	return id, secret
}

func LoadSongs(dir string) ([]SortedTrack, error) {
	var allSongs []SortedTrack

	// Read the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Iterate over each file in the directory
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(dir, file.Name())
			jsonFile, err := os.Open(filePath)
			if err != nil {
				log.Printf("Error opening file %s: %v\n", filePath, err)
				continue
			}

			// Read the JSON file
			byteValue, err := ioutil.ReadAll(jsonFile)
			if err != nil {
				log.Printf("Error reading file %s: %v\n", filePath, err)
				jsonFile.Close()
				continue
			}

			var songs []SortedTrack
			err = json.Unmarshal(byteValue, &songs)
			if err != nil {
				log.Printf("Error unmarshalling JSON from file %s: %v\n", filePath, err)
				jsonFile.Close()
				continue
			}

			// Append the songs to the aggregate slice
			allSongs = append(allSongs, songs...)

			// Close the JSON file
			jsonFile.Close()
		}
	}

	return allSongs, nil
}

func createSpotifyClient(ctx context.Context, clientID string, clientSecret string) (*spotify.Client, error) {
	cfg := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotifyauth.TokenURL,
	}
	tok, err := cfg.Token(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get access_token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, tok)
	client := spotify.New(httpClient)
	return client, nil
}

// saveHabit serializes the Habit structure to JSON and saves it to a file.
func SaveJson(name string, dir string, obj any) error {
	// Serialize the Habit structure to JSON
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	// Create the directory if it doesn't already exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
	}

	// Define the file path where the JSON will be saved
	filePath := filepath.Join(dir, fmt.Sprintf("%s.json", name))

	// Write the JSON data to the file
	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("error writing JSON file: %w", err)
	}

	return nil
}
