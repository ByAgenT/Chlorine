package server

import "net/http"

var (
	// Music handler
	externalMusicHandler ExternalMusicHandler

	// Authentication handlers
	loginHandler        LoginHandler
	completeAuthHandler CompleteAuthHandler
	spotifyTokenHandler SpotifyTokenHandler

	// Music handlers
	playlistsHandler        MyPlaylistsHandler
	availableDevicesHandler AvailableDevicesHandler
	playbackHandler         PlaybackHandler
	searchSongHandler       SearchSongHandler
	spotifyPlayHandler      SpotifyPlayHandler

	// Chlorine API handlers
	roomHandler               RoomHandler
	memberHandler             MemberHandler
	roomMembersHandler        RoomMembersHandler
	roomSongsDetailDispatcher DispatchableHandler
	roomSongsDetailHandler    RoomSongsDetailHandler
	roomSongsHandler          RoomSongsHandler
	roomSongsSpotifiedHandler RoomsSongsSpotifiedHandler

	// WebSocket handlers
	wsHandler WebSocketInitHandler
)

func initHandlers() {
	externalMusicHandler = ExternalMusicHandler{MusicService: musicService,
		AuthenticationProvider: authenticationProvider}

	// Authentication handlers init
	loginHandler = LoginHandler{}
	completeAuthHandler = CompleteAuthHandler{MemberService: memberService, RoomService: roomService}
	spotifyTokenHandler = SpotifyTokenHandler{}

	// Music handlers init
	playlistsHandler = MyPlaylistsHandler{ExternalMusicHandler: externalMusicHandler}
	availableDevicesHandler = AvailableDevicesHandler{ExternalMusicHandler: externalMusicHandler}
	playbackHandler = PlaybackHandler{ExternalMusicHandler: externalMusicHandler}
	searchSongHandler = SearchSongHandler{ExternalMusicHandler: externalMusicHandler}
	spotifyPlayHandler = SpotifyPlayHandler{ExternalMusicHandler: externalMusicHandler}

	// Chlorine API handlers init
	roomHandler = RoomHandler{MemberService: memberService, RoomService: roomService}
	memberHandler = MemberHandler{MemberService: memberService, TokenService: tokenService}
	roomMembersHandler = RoomMembersHandler{MemberService: memberService, RoomService: roomService}
	roomSongsHandler = RoomSongsHandler{ExternalMusicHandler: externalMusicHandler,
		SongService: songService, MemberService: memberService, RoomService: roomService}
	roomSongsDetailHandler = RoomSongsDetailHandler{
		ExternalMusicHandler: externalMusicHandler,
		SongService:          songService,
		MemberService:        memberService,
		RoomService:          roomService}
	roomSongsDetailDispatcher = DispatchableHandler{
		DeleteMethodHandler: roomSongsDetailHandler,
	}
	roomSongsSpotifiedHandler = RoomsSongsSpotifiedHandler{
		ExternalMusicHandler: externalMusicHandler, MemberService: memberService, SongService: songService, RoomService: roomService}

	// WebSocket handlers init
	wsHandler = WebSocketInitHandler{MemberService: memberService}
}

type DeleteMethodHandler interface {
	Delete(w http.ResponseWriter, r *http.Request)
}

type GetMethodHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type PostMethodHandler interface {
	Post(w http.ResponseWriter, r *http.Request)
}

type PutMethodHandler interface {
	Put(w http.ResponseWriter, r *http.Request)
}

type DispatchableHandler struct {
	DeleteMethodHandler
	GetMethodHandler
	PostMethodHandler
	PutMethodHandler
}

func (h DispatchableHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if h.GetMethodHandler != nil {
			h.GetMethodHandler.Get(w, r)
			return
		}
	case http.MethodPost:
		if h.PostMethodHandler != nil {
			h.PostMethodHandler.Post(w, r)
			return
		}
	case http.MethodDelete:
		if h.DeleteMethodHandler != nil {
			h.DeleteMethodHandler.Delete(w, r)
			return
		}
	case http.MethodPut:
		if h.PutMethodHandler != nil {
			h.PutMethodHandler.Put(w, r)
			return
		}
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
