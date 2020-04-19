import { useCallback, useEffect, useState } from 'react';
import { connectPlayer, SpotifyPlayer } from '../services/spotifyPlaybackService';
import { ChlorineService } from '../services/chlorineService';
import shuffle from 'lodash/shuffle';
import { Song, SpotifyTrack, Token } from '../models/chlorine';

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
  const [playlist, setPlaylist] = useState<Song[]>([]);
  const [spotifyTrackInfo, setSpotifyTrackInfo] = useState<SpotifyTrack[]>([]);

  const fetchPlaylist = useCallback(async function () {
    try {
      const fetchedSongs = await new ChlorineService().retrieveRoomSongs();
      setPlaylist(fetchedSongs);
    } catch (error) {
      console.error(error);
    }
  }, []);

  const fetchSpotifyTrackInfo = useCallback(async function () {
    try {
      const fetchedInfo = await new ChlorineService().retrieveRoomsSongsFromSpotify();
      setSpotifyTrackInfo(fetchedInfo);
    } catch (error) {
      console.error(error);
    }
  }, []);

  // TODO: revisit and optimize
  async function appendSong(spotifyId: string): Promise<Song> {
    const chlorineService = new ChlorineService();
    await Promise.all([fetchPlaylist(), fetchSpotifyTrackInfo()]);
    const lastSong = playlist.filter((song) => song.nextSongId == null);
    if (lastSong[0]) {
      try {
        const newSong = await chlorineService.addSong(spotifyId, lastSong[0].id, null);
        await chlorineService.updateSong(lastSong[0].id, {
          spotifyId: lastSong[0].spotifyId,
          prevSongId: lastSong[0].previousSongId,
          nextSongId: newSong.id,
        });
        await Promise.all([fetchPlaylist(), fetchSpotifyTrackInfo()]);
        return newSong;
      } catch (error) {
        console.error(error);
        return;
      }
    }
    const song = await chlorineService.addSong(spotifyId, null, null);
    await Promise.all([fetchPlaylist(), fetchSpotifyTrackInfo()]);
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
    setSpotifyTrackInfo(shuffle(spotifyTrackInfo));
  }

  useEffect(() => {
    Promise.all([fetchPlaylist(), fetchSpotifyTrackInfo()]);
  }, [fetchPlaylist, fetchSpotifyTrackInfo]);

  return {
    playlist,
    spotifyTrackInfo,
    fetchPlaylist,
    fetchSpotifyTrackInfo,
    appendSong,
    startPlay,
    doShuffle,
  };
}

export { useSpotifyPlayer, useSpotifyPlaylist, usePlaybackInformation, SimplePlaybackInformation };
