import { useCallback, useEffect, useState } from 'react';
import { connectPlayer, SpotifyPlayer } from '../services/spotifyPlaybackService';
import { ChlorineService } from '../services/chlorineService';
import shuffle from 'lodash/shuffle';
import { Album, Artist, Song, SpotifyTrack, Token } from '../models/chlorine';

interface PlaylistTrack {
  id: number;
  spotifyId: string;
  name: string;
  artists: Artist[];
  album: Album;
  durationMs: number;
}

interface SimplePlaybackInformation {
  now: number;
  duration: number;
  artistTitle: string;
  songTitle: string;
  albumCoverURL: string;
}

function useSpotifyPlayer(): SpotifyPlayer {
  const [player, setPlayer] = useState<SpotifyPlayer | null>(null);

  useEffect(() => {
    const getSpotifyToken = async function () {
      return new ChlorineService().getToken();
    };

    async function getSpotifyPlayer(token: Token): Promise<SpotifyPlayer> {
      return new Promise((resolve) => {
        const playerReceiveInterval = setInterval(() => {
          let player = connectPlayer(token.accessToken);
          if (player) {
            clearInterval(playerReceiveInterval);
            resolve(player);
          }
        }, 1000);
      });
    }

    async function prepare() {
      try {
        const token = await getSpotifyToken();
        const player = await getSpotifyPlayer(token);
        await player.connect();
        setPlayer(player);
      } catch (error) {
        console.log(error);
      }
    }

    prepare();
  }, []);

  return player;
}

function usePlaybackInformation(player: SpotifyPlayer | null) {
  const [playbackInformation, setPlaybackInformation] = useState<SimplePlaybackInformation>({
    albumCoverURL: '',
    artistTitle: '',
    duration: 0,
    now: 0,
    songTitle: '',
  });

  useEffect(() => {
    async function prepare() {
      try {
        const playback = await new ChlorineService().getPlaybackInformation();

        const playbackInfo = {
          now: playback.progressMs,
          duration: playback.item ? playback.item.durationMs : NaN,
          artistTitle: playback.item
            ? playback.item.artists.map((artist) => artist.name).join(', ')
            : '',
          songTitle: playback.item ? playback.item.name : '',
          albumCoverURL: playback.item
            ? playback.item.album.images.filter(
                (image) => image.width > 200 && image.width < 400
              )[0].url
            : '',
        };
        setPlaybackInformation(playbackInfo);
      } catch (error) {
        console.error(error);
      }
    }

    prepare();
  }, []);

  useEffect(() => {
    if (player !== null) {
      player.onPlayerStateChanged((state: Spotify.PlaybackState) => {
        if (state) {
          if (state.track_window) {
            const currentTrack = state.track_window.current_track;
            const playbackInfo = {
              now: state.position,
              duration: state.duration,
              artistTitle: currentTrack.artists.map((artist) => artist.name).join(', '),
              songTitle: currentTrack.name,
              albumCoverURL: currentTrack.album.images.filter(
                (image) => image.width > 200 && image.width < 400
              )[0].url,
            };
            setPlaybackInformation(playbackInfo);
          }
        }
      });
    }
  }, [player]);

  return playbackInformation;
}

function useSpotifyPlaylist() {
  const [playlist, setPlaylist] = useState<PlaylistTrack[]>();
  const [songPlaylist, setSongPlaylist] = useState<Song[]>([]);
  const [spotifyTrackInfo, setSpotifyTrackInfo] = useState<SpotifyTrack[]>([]);

  const fetchTracks = useCallback(async () => {
    try {
      const chlorineService = new ChlorineService();
      const fetchedSongs = await chlorineService.retrieveRoomSongs();
      setSongPlaylist(fetchedSongs);
      const spotifySongs = await chlorineService.retrieveRoomsSongsFromSpotify();
      setSpotifyTrackInfo(spotifySongs);
      const fullPlaylist = fetchedSongs.map<PlaylistTrack>((song, index) => {
        return {
          id: song.id,
          spotifyId: spotifySongs[index].id,
          name: spotifySongs[index].name,
          artists: spotifySongs[index].artists,
          album: spotifySongs[index].album,
          durationMs: spotifySongs[index].durationMs
        };
      });
      setPlaylist(fullPlaylist);
    } catch (error) {
      console.error(error);
    }
  }, [setSongPlaylist, setSpotifyTrackInfo, setPlaylist]);

  // TODO: revisit and optimize
  async function appendSong(spotifyId: string): Promise<Song> {
    const chlorineService = new ChlorineService();
    await fetchTracks();
    const lastSong = songPlaylist.filter((song) => song.nextSongId == null);
    if (lastSong[0]) {
      try {
        const newSong = await chlorineService.addSong(spotifyId, lastSong[0].id, null);
        await chlorineService.updateSong(lastSong[0].id, {
          spotifyId: lastSong[0].spotifyId,
          prevSongId: lastSong[0].previousSongId,
          nextSongId: newSong.id,
        });
        await fetchTracks();
        return newSong;
      } catch (error) {
        console.error(error);
        return;
      }
    }
    const song = await chlorineService.addSong(spotifyId, null, null);
    await fetchTracks();
    return song;
  }

  async function startPlay(): Promise<void> {
    try {
      await new ChlorineService().play(spotifyTrackInfo.map((track) => track.uri));
    } catch (error) {
      console.error(error);
    }
  }

  async function doShuffle() {
    setPlaylist(shuffle(playlist));
  }

  useEffect(() => {
    fetchTracks();
  }, [fetchTracks]);

  return {
    playlist,
    spotifyTrackInfo,
    fetchTracks,
    appendSong,
    startPlay,
    doShuffle,
  };
}

export { useSpotifyPlayer, useSpotifyPlaylist, usePlaybackInformation, SimplePlaybackInformation, PlaylistTrack };
