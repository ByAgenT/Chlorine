package server

var (
	// Music handler
	externalMusicHandler = ExternalMusicHandler{MusicService: musicService, AuthenticationProvider: authenticationProvider}

	// Storage handler
	storageHandler = StorageHandler{}

	// Authentication handlers
	loginHandler        = LoginHandler{StorageHandler: storageHandler}
	completeAuthHandler = CompleteAuthHandler{StorageHandler: storageHandler}
	spotifyTokenHandler = SpotifyTokenHandler{}

	// Music handlers
	playlistsHandler        = MyPlaylistsHandler{ExternalMusicHandler: externalMusicHandler}
	availableDevicesHandler = AvailableDevicesHandler{ExternalMusicHandler: externalMusicHandler}
	playbackHandler         = PlaybackHandler{ExternalMusicHandler: externalMusicHandler}
	searchSongHandler       = SearchSongHandler{ExternalMusicHandler: externalMusicHandler}
	spotifyPlayHandler      = SpotifyPlayHandler{ExternalMusicHandler: externalMusicHandler}

	// Chlorine API handlers
	roomHandler               = RoomHandler{StorageHandler: storageHandler}
	memberHandler             = MemberHandler{StorageHandler: storageHandler}
	roomMembersHandler        = RoomMembersHandler{StorageHandler: storageHandler}
	roomSongsHanlder          = RoomSongsHandler{StorageHandler: storageHandler, ExternalMusicHandler: externalMusicHandler}
	roomSongsSpotifiedHandler = RoomsSongsSpotifiedHandler{StorageHandler: storageHandler, ExternalMusicHandler: externalMusicHandler}
)
