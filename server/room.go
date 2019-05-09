package server

import (
	"chlorine/apierror"
	"chlorine/auth"
	"chlorine/cl"
	"chlorine/storage"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/zmb3/spotify"
)

// RoomHandler handle room creation and retrieving inforamtion about rooms.
type RoomHandler struct {
	auth.Session
	StorageHandler
}

func (h RoomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	switch r.Method {
	case "GET":
		memberID, ok := session.Values["MemberID"].(storage.ID)
		if !ok {
			jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
			return
		}
		member, err := h.storage.GetMember(memberID)
		if err != nil {
			log.Printf("server: MemberHandler: cannot retrieve member: %s", err)
			return
		}
		room, err := h.storage.GetRoom(storage.ID(member.RoomID))
		if err != nil {
			log.Printf("server: RoomHandler: %s", err.Error())
			jsonWriter.Error(apierror.APIServerError, 500)
			return
		}
		jsonWriter.WriteJSONObject(room)
		return
	}

}

// RoomMembersHandler handle serving members of the current room
type RoomMembersHandler struct {
	auth.Session
	StorageHandler
}

func (h RoomMembersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	switch r.Method {
	case "GET":
		memberID, ok := session.Values["MemberID"].(storage.ID)
		if !ok {
			jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
			return
		}
		member, err := h.storage.GetMember(memberID)
		if err != nil {
			log.Printf("server: RoomMemberHandler: cannot retrieve member: %s", err)
			return
		}
		room, err := h.storage.GetRoom(storage.ID(member.RoomID))
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
		return
	}
}

// RoomSongsHandler handle serving songs that are assinged to the current room.
type RoomSongsHandler struct {
	auth.Session
	StorageHandler
	ExternalMusicHandler
}

func (h RoomSongsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	switch r.Method {
	case "GET":
		memberID, ok := session.Values["MemberID"].(storage.ID)
		if !ok {
			jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
			return
		}
		member, err := h.storage.GetMember(memberID)
		if err != nil {
			log.Printf("server: RoomSongsHandler: cannot retrieve member: %s", err)
			return
		}
		room, err := h.storage.GetRoom(storage.ID(member.RoomID))
		if err != nil {
			log.Printf("server: RoomSongsHandler: %s", err.Error())
			jsonWriter.Error(apierror.APIServerError, 500)
			return
		}
		songs, err := h.storage.GetRoomSongs(room)
		jsonWriter.WriteJSONObject(songs)
		return
	case "POST":
		memberID, ok := session.Values["MemberID"].(storage.ID)
		if !ok {
			jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
			return
		}
		member, err := h.storage.GetMember(memberID)
		room, err := h.storage.GetRoom(storage.ID(member.RoomID))
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
		song, err := cl.CreateSong(songData.SpotifyID, int(*room.ID), songData.PreviousSongID, songData.NextSongID, member, h.storage)
		if err != nil {
			log.Printf("server: RoomSongsHandler: cannot create song: %s", err)
			jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
			return
		}
		jsonWriter.WriteJSONObject(song)
		return
	case "PUT":
		memberID, ok := session.Values["MemberID"].(storage.ID)
		if !ok {
			jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
			return
		}
		member, err := h.storage.GetMember(memberID)
		room, err := h.storage.GetRoom(storage.ID(member.RoomID))
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
		song, err := cl.UpdateSong(songData.ID, songData.SpotifyID, int(*room.ID), songData.PreviousSongID, songData.NextSongID, member, h.storage)
		if err != nil {
			log.Printf("server: RoomSongsHandler: cannot create song: %s", err)
			jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
			return
		}
		jsonWriter.WriteJSONObject(song)
		return
	}
}

// RoomsSongsSpotifiedHandler is a struct that serves songs from Spotify.
type RoomsSongsSpotifiedHandler struct {
	auth.Session
	StorageHandler
	ExternalMusicHandler
}

func (h RoomsSongsSpotifiedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	switch r.Method {
	case "GET":
		memberID, ok := session.Values["MemberID"].(storage.ID)
		if !ok {
			jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
			return
		}
		member, err := h.storage.GetMember(memberID)
		if err != nil {
			log.Printf("server: RoomSongsHandler: cannot retrieve member: %s", err)
			return
		}
		room, err := h.storage.GetRoom(storage.ID(member.RoomID))
		if err != nil {
			log.Printf("server: RoomSongsHandler: %s", err.Error())
			jsonWriter.Error(apierror.APIServerError, 500)
			return
		}
		songs, err := h.storage.GetRoomSongs(room)
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
		jsonWriter.WriteJSONObject(tracks)
		return
	}
}
