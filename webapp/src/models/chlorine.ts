enum RoleType {
  Administrator = 1,
  Member = 2,
}

interface Member {
  id: number;
  name: string;
  roomId: number;
  role: RoleType;
  createdDate: Date;
}

interface Token {
  accessToken: string;
  tokenType: string;
  refreshToken: string;
  expiry: Date;
}

interface Song {
  id: number;
  spotifyId: string;
  roomId: number;
  previousSongId: number | null;
  nextSongId: number | null;
  memberAddedId: number;
  createdDate: Date;
}

interface PlaybackInformation {
  timestamp: number;
  progressMs: number;
  isPlaying: boolean;
  item: SpotifyTrack | null;
  device: Device;
  shuffleState: boolean;
  repeatState: string; // TODO make repeatState Enum.
}

interface SpotifyTrack {
  artists: Artist[];
  availableMarkets: string[];
  discNumber: number;
  durationMs: number;
  explicit: boolean;
  id: string;
  name: string;
  uri: string;
  album: Album;
  popularity: number;
}

interface Artist {
  name: string;
  id: string;
  uri: string;
  href: string;
}

interface Album {
  name: string;
  artists: Artist[];
  albumGroup: string;
  albumType: string;
  id: string;
  uri: string;
  availableMarkets: string[];
  href: string;
  images: Image[];
  releaseDate: Date;
  releaseDatePrecision: string; // TODO make enum.
}

interface Image {
  height: number;
  width: number;
  url: string;
}

interface Device {
  id: string;
  isActive: boolean;
  isRestricted: boolean;
  name: string;
  type: string; // TODO make type to be enum.
  volumePercent: number;
}

export { Member, Token, PlaybackInformation, Device, Song, Artist, Album, Image, SpotifyTrack };
