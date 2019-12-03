package server

var (
	// Music handler
	externalMusicHandler ExternalMusicHandler

	// Storage handler
	storageHandler StorageHandler

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
	storageHandler = StorageHandler{storage: dbStorage}

	// Authentication handlers init
	loginHandler = LoginHandler{StorageHandler: storageHandler}
	completeAuthHandler = CompleteAuthHandler{StorageHandler: storageHandler, MemberService: memberService, RoomService: roomService}
	spotifyTokenHandler = SpotifyTokenHandler{}

	// Music handlers init
	playlistsHandler = MyPlaylistsHandler{ExternalMusicHandler: externalMusicHandler}
	availableDevicesHandler = AvailableDevicesHandler{ExternalMusicHandler: externalMusicHandler}
	playbackHandler = PlaybackHandler{ExternalMusicHandler: externalMusicHandler}
	searchSongHandler = SearchSongHandler{ExternalMusicHandler: externalMusicHandler}
	spotifyPlayHandler = SpotifyPlayHandler{ExternalMusicHandler: externalMusicHandler}

	// Chlorine API handlers init
	roomHandler = RoomHandler{StorageHandler: storageHandler, MemberService: memberService, RoomService: roomService}
	memberHandler = MemberHandler{StorageHandler: storageHandler, MemberService: memberService}
	roomMembersHandler = RoomMembersHandler{StorageHandler: storageHandler, MemberService: memberService, RoomService: roomService}
	roomSongsHandler = RoomSongsHandler{StorageHandler: storageHandler, ExternalMusicHandler: externalMusicHandler,
		SongService: songService, MemberService: memberService, RoomService: roomService}
	roomSongsSpotifiedHandler = RoomsSongsSpotifiedHandler{StorageHandler: storageHandler,
		ExternalMusicHandler: externalMusicHandler, MemberService: memberService, SongService: songService, RoomService: roomService}
}
