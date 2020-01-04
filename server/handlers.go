package server

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
	roomSongsHandler          RoomSongsHandler
	roomSongsSpotifiedHandler RoomsSongsSpotifiedHandler
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
	roomSongsSpotifiedHandler = RoomsSongsSpotifiedHandler{
		ExternalMusicHandler: externalMusicHandler, MemberService: memberService, SongService: songService, RoomService: roomService}
}
