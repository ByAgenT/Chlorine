package server

import (
	"chlorine/apierror"
	"chlorine/auth"
	"chlorine/cl"
	"chlorine/ws"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/zmb3/spotify"
)

// RoomHandler handle room creation and retrieving information about rooms.
type RoomHandler struct {
	auth.Session
	RoomService   cl.RoomService
	MemberService cl.MemberService
}

func (h RoomHandler) Get(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	member, ok := getMemberIfAuthorized(h.MemberService, session)
	if !ok {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
		return
	}
	room, err := h.RoomService.GetRoom(int(member.RoomID))
	if err != nil {
		log.Printf("server: RoomHandler: %s", err.Error())
		jsonWriter.Error(apierror.APIServerError, 500)
		return
	}
	jsonWriter.WriteJSONObject(room)
}

// RoomMembersHandler handle serving members of the current room
type RoomMembersHandler struct {
	auth.Session
	MemberService cl.MemberService
	RoomService   cl.RoomService
}

func (h RoomMembersHandler) Get(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	member, ok := getMemberIfAuthorized(h.MemberService, session)
	if !ok {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
		return
	}
	room, err := h.RoomService.GetRoom(int(member.RoomID))
	if err != nil {
		log.Printf("server: RoomMemberHandler: %s", err.Error())
		jsonWriter.Error(apierror.APIServerError, 500)
		return
	}
	members, err := room.GetMembers()
	if err != nil {
		log.Printf("server: RoomMemberHandler: %s", err.Error())
		jsonWriter.Error(apierror.APIServerError, 500)
		return
	}
	jsonWriter.WriteJSONObject(members)
}

// RoomSongsHandler handle serving songs that are assigned to the current room.
type RoomSongsHandler struct {
	auth.Session
	ExternalMusicHandler
	SongService   cl.SongService
	MemberService cl.MemberService
	RoomService   cl.RoomService
}

func (h RoomSongsHandler) Get(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}
	member, ok := getMemberIfAuthorized(h.MemberService, session)
	if !ok {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
		return
	}
	songs, err := h.SongService.GetRoomSongs(int(member.RoomID))
	if err != nil {
		log.Printf("server: RoomSongsHandler: %s", err)
		jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
		return
	}
	jsonWriter.WriteJSONObject(songs)
}

func (h RoomSongsHandler) Post(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}
	member, ok := getMemberIfAuthorized(h.MemberService, session)
	if !ok {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
		return
	}
	room, err := h.RoomService.GetRoom(int(member.RoomID))
	if err != nil {
		log.Printf("server: RoomSongsHandler: %s", err.Error())
		jsonWriter.Error(apierror.APIServerError, 500)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("server: RoomSongsHandler: error reading requset body: %s", err)
		jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
		return
	}
	songData := &struct {
		SpotifyID      string `json:"spotify_id,omitempty"`
		PreviousSongID int    `json:"previous_song_id,omitempty"`
		NextSongID     int    `json:"next_song_id,omitempty"`
	}{}
	err = json.Unmarshal(body, &songData)
	if err != nil {
		log.Printf("server: RoomSongsHandler: json parse error: %s", err)
		jsonWriter.Error(apierror.APIInvalidRequest, http.StatusBadRequest)
		return
	}
	song, err := h.SongService.CreateSong(cl.RawSong{
		SpotifyID:      songData.SpotifyID,
		RoomID:         int(*room.ID),
		PreviousSongID: songData.PreviousSongID,
		NextSongID:     songData.NextSongID,
		MemberCreated:  member,
	})
	if err != nil {
		log.Printf("server: RoomSongsHandler: cannot create song: %s", err)
		jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
		return
	}
	jsonWriter.WriteJSONObject(song)
	ws.Broadcast(roomWSConnections[int(member.RoomID)], &ws.Response{
		Type:        ws.TypeBroadcast,
		Status:      ws.StatusOK,
		Description: "SongAdded",
		Body: map[string]interface{}{
			"song": song,
		},
	})
}

func (h RoomSongsHandler) Put(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}
	member, ok := getMemberIfAuthorized(h.MemberService, session)
	if !ok {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
		return
	}
	room, err := h.RoomService.GetRoom(int(member.RoomID))
	if err != nil {
		log.Printf("server: RoomSongsHandler: %s", err.Error())
		jsonWriter.Error(apierror.APIServerError, 500)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("server: RoomSongsHandler: error reading requset body: %s", err)
		jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
		return
	}
	songData := &struct {
		ID             int    `json:"id,omitempty"`
		SpotifyID      string `json:"spotify_id,omitempty"`
		PreviousSongID int    `json:"previous_song_id,omitempty"`
		NextSongID     int    `json:"next_song_id,omitempty"`
	}{}
	err = json.Unmarshal(body, &songData)
	if err != nil {
		log.Printf("server: RoomSongsHandler: json parse error: %s", err)
		jsonWriter.Error(apierror.APIInvalidRequest, http.StatusBadRequest)
		return
	}
	song, err := h.SongService.UpdateSong(songData.ID, cl.RawSong{
		SpotifyID:      songData.SpotifyID,
		RoomID:         int(*room.ID),
		PreviousSongID: songData.PreviousSongID,
		NextSongID:     songData.NextSongID,
		MemberCreated:  member,
	})
	if err != nil {
		log.Printf("server: RoomSongsHandler: cannot create song: %s", err)
		jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
		return
	}
	jsonWriter.WriteJSONObject(song)
}

// RoomsSongsSpotifiedHandler is a struct that serves songs from Spotify.
type RoomsSongsSpotifiedHandler struct {
	auth.Session
	ExternalMusicHandler
	MemberService cl.MemberService
	RoomService   cl.RoomService
	SongService   cl.SongService
}

func (h RoomsSongsSpotifiedHandler) Get(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}
	member, ok := getMemberIfAuthorized(h.MemberService, session)
	if !ok {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
		return
	}
	room, err := h.RoomService.GetRoom(int(member.RoomID))
	if err != nil {
		log.Printf("server: RoomSongsHandler: %s", err.Error())
		jsonWriter.Error(apierror.APIServerError, 500)
		return
	}
	songs, err := h.SongService.GetRoomSongs(int(*room.ID))
	client, err := h.GetClient(session)
	if err != nil {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
		return
	}
	trackIDs := make([]spotify.ID, 0)
	for _, song := range songs {
		trackIDs = append(trackIDs, spotify.ID(song.SpotifyID))
	}
	tracks, err := client.GetTracks(trackIDs...)
	if tracks == nil {
		jsonWriter.WriteJSONObject([]spotify.FullTrack{})
		return
	}
	jsonWriter.WriteJSONObject(tracks)
}

type RoomSongsDetailHandler struct {
	auth.Session
	ExternalMusicHandler
	SongService   cl.SongService
	MemberService cl.MemberService
	RoomService   cl.RoomService
}

func (h RoomSongsDetailHandler) Delete(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}
	member, ok := getMemberIfAuthorized(h.MemberService, session)
	if !ok {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	songID, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonWriter.Error(apierror.APIInvalidRequest, http.StatusBadRequest)
		return
	}
	err = h.SongService.DeleteSong(songID)
	if err != nil {
		log.Printf("error deleting song: %s", err)
		jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
		return
	}
	ws.Broadcast(roomWSConnections[int(member.RoomID)], &ws.Response{
		Type:        ws.TypeBroadcast,
		Status:      ws.StatusOK,
		Description: "SongDeleted",
		Body: map[string]interface{}{
			"songID": songID,
		},
	})
}
