/* eslint-disable @typescript-eslint/camelcase */
import {
  Album,
  Artist,
  Device,
  Image,
  Member,
  PlaybackInformation,
  Song,
  SpotifyTrack,
  Token,
} from '../models/chlorine';

class ChlorineServiceError extends Error {
  // TODO: incorporate more details in error.
  constructor(message: string) {
    super();
    this.message = message;
    Object.setPrototypeOf(this, ChlorineServiceError.prototype);
  }
}

interface ArtistResponse {
  name: string;
  id: string;
  uri: string;
  href: string;
}

interface AlbumResponse {
  name: string;
  artists: ArtistResponse[];
  album_group: string;
  album_type: string;
  id: string;
  uri: string;
  available_markets: string[];
  href: string;
  images: ImagesResponse[];
  release_date: string;
  release_date_precision: string;
}

interface ImagesResponse {
  height: number;
  width: number;
  url: string;
}

interface SpotifyTrackResponse {
  artists: ArtistResponse[];
  available_markets: string[];
  disc_number: number;
  duration_ms: number;
  explicit: false;
  id: string;
  name: string;
  uri: string;
  album: AlbumResponse;
  popularity: number;
}

interface DeviceResponse {
  id: string;
  is_active: boolean;
  is_restricted: boolean;
  name: string;
  type: string;
  volume_percent: number;
}

class ChlorineResponseMapper {
  public mapToken(tokenResponse: {
    access_token: string;
    token_type: string;
    refresh_type: string;
    expiry: string;
  }): Token {
    return {
      accessToken: tokenResponse.access_token,
      tokenType: tokenResponse.token_type,
      refreshToken: tokenResponse.refresh_type,
      expiry: new Date(tokenResponse.expiry),
    };
  }

  public mapMember(memberResponse: {
    id: number;
    name: string;
    room_id: number;
    role: number;
    created_date: string;
  }): Member {
    return {
      id: memberResponse.id,
      name: memberResponse.name,
      roomId: memberResponse.room_id,
      role: memberResponse.role,
      createdDate: new Date(memberResponse.created_date),
    };
  }

  public mapImage(imageResponse: ImagesResponse): Image {
    return { height: imageResponse.height, url: imageResponse.url, width: imageResponse.width };
  }

  public mapAlbum(albumResponse: AlbumResponse): Album {
    return {
      albumGroup: albumResponse.album_group,
      albumType: albumResponse.album_type,
      artists: albumResponse.artists.map((artist) => this.mapArtist(artist)),
      availableMarkets: albumResponse.available_markets,
      href: albumResponse.href,
      id: albumResponse.id,
      images: albumResponse.images.map((image) => this.mapImage(image)),
      name: albumResponse.name,
      releaseDate: new Date(albumResponse.release_date),
      releaseDatePrecision: albumResponse.release_date_precision,
      uri: albumResponse.uri,
    };
  }

  public mapArtist(artistResponse: ArtistResponse): Artist {
    return {
      href: artistResponse.href,
      id: artistResponse.id,
      name: artistResponse.name,
      uri: artistResponse.uri,
    };
  }

  public mapSpotifyTrack(songResponse: SpotifyTrackResponse): SpotifyTrack {
    return {
      album: this.mapAlbum(songResponse.album),
      artists: songResponse.artists.map((artist) => this.mapArtist(artist)),
      availableMarkets: songResponse.available_markets,
      discNumber: songResponse.disc_number,
      durationMs: songResponse.duration_ms,
      explicit: songResponse.explicit,
      id: songResponse.id,
      name: songResponse.name,
      popularity: songResponse.popularity,
      uri: songResponse.uri,
    };
  }

  public mapDevice(deviceResponse: DeviceResponse): Device {
    return {
      id: deviceResponse.id,
      isActive: deviceResponse.is_active,
      isRestricted: deviceResponse.is_restricted,
      name: deviceResponse.name,
      type: deviceResponse.type,
      volumePercent: deviceResponse.volume_percent,
    };
  }

  public mapPlaybackInformation(playbackInformationResponse: {
    timestamp: number;
    progress_ms: number;
    is_playing: boolean;
    Item: SpotifyTrackResponse;
    device: DeviceResponse;
    shuffle_state: boolean;
    repeat_state: string;
  }): PlaybackInformation {
    return {
      device: this.mapDevice(playbackInformationResponse.device),
      isPlaying: playbackInformationResponse.is_playing,
      item: playbackInformationResponse.Item
        ? this.mapSpotifyTrack(playbackInformationResponse.Item)
        : null,
      progressMs: playbackInformationResponse.progress_ms,
      repeatState: playbackInformationResponse.repeat_state,
      shuffleState: playbackInformationResponse.shuffle_state,
      timestamp: playbackInformationResponse.timestamp,
    };
  }

  public async mapSong(songResponse: {
    id: number;
    spotify_id: string;
    room_id: number;
    previous_song_id: number | null;
    next_song_id: number | null;
    member_added_id: number;
    created_date: string;
  }): Promise<Song> {
    return {
      nextSongId: songResponse.next_song_id,
      previousSongId: songResponse.previous_song_id,
      id: songResponse.id,
      spotifyId: songResponse.spotify_id,
      roomId: songResponse.room_id,
      memberAddedId: songResponse.member_added_id,
      createdDate: new Date(songResponse.created_date),
    };
  }
}

class ChlorineService {
  private responseMapper: ChlorineResponseMapper;

  constructor() {
    // TODO: implement dependency injection here.
    this.responseMapper = new ChlorineResponseMapper();
  }

  public async joinRoom(roomID: number, name: string): Promise<void> {
    const response = await fetch('api/member', {
      method: 'POST',
      credentials: 'include',
      body: JSON.stringify({ name: name, room_id: Number(roomID) }),
    });

    if (!response.ok) {
      throw new ChlorineServiceError('Error joining room.');
    }
  }

  public async getToken(): Promise<Token> {
    const response = await fetch('api/token', {
      credentials: 'include',
    });

    if (!response.ok) {
      throw new ChlorineServiceError('Cannot retrieve token');
    }

    const token = await response.json();
    return this.responseMapper.mapToken(token);
  }

  public async getMemberInfo(): Promise<Member> {
    const response = await fetch('api/member', {
      credentials: 'include',
    });

    if (!response.ok) {
      throw new ChlorineServiceError('Cannot get member info');
    }

    const member = await response.json();
    return this.responseMapper.mapMember(member);
  }

  public async getRoomMembers(): Promise<Member[]> {
    const response = await fetch('api/room/members', {
      credentials: 'include',
    });

    if (!response.ok) {
      throw new ChlorineServiceError('Cannot get room members info');
    }

    const members = await response.json();
    return members.map((member) => this.responseMapper.mapMember(member));
  }

  public async getPlaybackInformation(): Promise<PlaybackInformation> {
    const response = await fetch('api/me/player/', {
      credentials: 'include',
    });

    if (!response.ok) {
      throw new ChlorineServiceError('Cannot get room members info');
    }
    const playbackInformation = await response.json();
    return this.responseMapper.mapPlaybackInformation(playbackInformation);
  }

  public async searchTracks(query: string): Promise<SpotifyTrack[]> {
    const url = new URL('/api/search', window.location.href);
    const params = { q: query };
    url.search = new URLSearchParams(params).toString();

    const response = await fetch(url.toString(), {
      credentials: 'include',
    });

    const data = await response.json();

    return data.tracks.items.map((track: SpotifyTrackResponse) =>
      this.responseMapper.mapSpotifyTrack(track)
    );
  }

  public async getDevicesInformation(): Promise<Device[]> {
    const response = await fetch('api/me/player/devices', {
      credentials: 'include',
    });

    const devices = await response.json();
    return devices.map((device) => this.responseMapper.mapDevice(device));
  }

  public async transferPlayback(deviceId: string, play: boolean): Promise<void> {
    const response = await fetch('api/me/player/', {
      method: 'PUT',
      credentials: 'include',
      body: JSON.stringify({
        device_id: deviceId,
        play: play,
      }),
    });
    if (!response.ok) {
      throw new ChlorineServiceError('Error transferring playback.');
    }
  }

  public async retrieveRoomSongs(): Promise<Song[]> {
    const response = await fetch('api/room/songs', {
      credentials: 'include',
    });

    const songs = await response.json();
    return songs.map((song) => this.responseMapper.mapSong(song));
  }

  public async retrieveRoomsSongsFromSpotify(): Promise<SpotifyTrack[]> {
    const response = await fetch('api/room/songs/spotify', {
      credentials: 'include',
    });

    const tracks = await response.json();
    return tracks.map((track) => this.responseMapper.mapSpotifyTrack(track));
  }

  public async addSong(
    spotifyId: string,
    prevSongId: number | null,
    nextSongId: number | null
  ): Promise<Song> {
    const response = await fetch('api/room/songs', {
      credentials: 'include',
      method: 'POST',
      body: JSON.stringify({
        spotify_id: spotifyId,
        previous_song_id: prevSongId,
        next_song_id: nextSongId,
      }),
    });
    const song = await response.json();

    return this.responseMapper.mapSong(song);
  }

  public async updateSong(
    song_id: number,
    values: {
      spotifyId: string;
      prevSongId: number;
      nextSongId: number;
    }
  ): Promise<Song> {
    const response = await fetch('api/room/songs', {
      credentials: 'include',
      method: 'PUT',
      body: JSON.stringify({
        id: song_id,
        spotify_id: values.spotifyId,
        previous_song_id: values.prevSongId,
        next_song_id: values.nextSongId,
      }),
    });

    const song = await response.json();
    return this.responseMapper.mapSong(song);
  }

  public async play(spotify_uris: string[]): Promise<void> {
    const response = await fetch('api/play', {
      credentials: 'include',
      method: 'POST',
      body: JSON.stringify({
        uris: spotify_uris,
      }),
    });

    if (!response.ok) {
      throw new ChlorineServiceError('Error while starting playing');
    }
  }
}

export { ChlorineService, ChlorineServiceError };
