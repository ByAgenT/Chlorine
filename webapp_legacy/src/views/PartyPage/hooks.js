import { useEffect, useState, useCallback } from 'react';
import {
  getPlaybackInformation,
  getRoomMembers,
  getToken,
  searchTracks,
  retrieveRoomSongs,
  retrieveRoomsSongsFromSpotify,
  addSong,
  updateSong,
  play
} from '../../services/ChlorineService';
import { connectPlayer } from '../../services/SpotifyPlaybackService';
import shuffle from 'lodash/shuffle';

function useMembersList() {
  const [members, setMembers] = useState([]);

  async function updateMembers() {
    try {
      const members = await getRoomMembers();
      setMembers(members);
    } catch (error) {
      console.error(error);
    }
  }

  useEffect(() => {
    async function prepare() {
      try {
        const members = await getRoomMembers();
        setMembers(members);
      } catch (error) {
        console.error(error);
      }
    }

    prepare();
  }, []);

  return [members, updateMembers];
}

function useSpotifyPlayer() {
  const [player, setPlayer] = useState(null);

  useEffect(() => {
    const getSpotifyToken = async function() {
      return await getToken();
    };

    async function getSpotifyPlayer(token) {
      return new Promise(resolve => {
        const playerReceiveInterval = setInterval(() => {
          let player = connectPlayer(token.access_token);
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
        player.connect();
        setPlayer(player);
      } catch (error) {
        console.log(error);
      }
    }

    prepare();
  }, []);

  return player;
}

function useSongSearch() {
  const [songQuery, setSongQuery] = useState('');
  const [searchResult, setSearchResult] = useState([]);

  useEffect(
    () => {
      async function prepare() {
        if (songQuery) {
          try {
            const songs = await searchTracks(songQuery);
            setSearchResult(songs.tracks.items);
          } catch (error) {
            console.log('Song search error');
            console.error(error);
          }
        }
      }

      prepare();
    },
    [songQuery]
  );

  return { searchResult, setSongQuery };
}

function usePlaybackInformation(player) {
  const [playbackInformation, setPlaybackInformation] = useState({});

  useEffect(() => {
    async function prepare() {
      try {
        const playback = await getPlaybackInformation();

        console.log(playback);
        const playbackInfo = {
          now: playback.progress_ms,
          duration: playback.Item ? playback.Item.duration_ms : NaN,
          artistTitle: playback.Item
            ? playback.Item.artists.map(artist => artist.name).join(', ')
            : '',
          songTitle: playback.Item ? playback.Item.name : '',
          albumCoverURL: playback.Item
            ? playback.Item.album.images.filter(
              image => image.width > 200 && image.width < 400
            )[0].url
            : ''
        };
        setPlaybackInformation(playbackInfo);
      } catch (error) {
        console.error(error);
      }
    }

    prepare();
  }, []);

  useEffect(
    () => {
      if (player !== null) {
        player.onPlayerStateChanged(state => {
          if (state) {
            if (state.track_window) {
              const currentTrack = state.track_window.current_track;
              const playbackInfo = {
                now: state.position,
                duration: state.duration,
                artistTitle: currentTrack.artists
                  .map(artist => artist.name)
                  .join(', '),
                songTitle: currentTrack.name,
                albumCoverURL: currentTrack.album.images.filter(
                  image => image.width > 200 && image.width < 400
                )[0].url
              };
              setPlaybackInformation(playbackInfo);
            }
          }
        });
      }
    },
    [player]
  );

  return playbackInformation;
}

function useSpotifyPlaylist() {
  const [playlist, setPlaylist] = useState([]);
  const [spotifyTrackInfo, setSpotifyTrackInfo] = useState([]);

  const fetchPlaylist = useCallback(async function() {
    try {
      const fetchedSongs = await retrieveRoomSongs();
      console.log(fetchedSongs);
      setPlaylist(fetchedSongs);
    } catch (error) {
      console.error(error);
    }
  }, []);

  const fetchSpotifyTrackInfo = useCallback(async function() {
    try {
      const fetchedInfo = await retrieveRoomsSongsFromSpotify();
      setSpotifyTrackInfo(fetchedInfo);
    } catch (error) {
      console.error(error);
    }
  }, []);

  async function appendSong(spotifyId) {
    fetchPlaylist();
    fetchSpotifyTrackInfo();
    const lastSong = playlist.filter(
      song => song.next_song_id === undefined || song.next_song_id === null
    );
    if (lastSong[0]) {
      try {
        const newSong = await addSong(spotifyId, lastSong[0].id, null);
        console.log('perform updating');
        const updateResponse = await updateSong(lastSong[0].id, {
          spotifyId: lastSong[0].spotify_id,
          prevSongId: lastSong[0].previous_song_id,
          nextSongId: newSong.id
        });
        console.log('Updated song: ');
        console.log(updateResponse);
        fetchPlaylist();
        fetchSpotifyTrackInfo();
        return newSong;
      } catch (error) {
        console.error(error);
        return;
      }
    }
    console.log('add first element');
    const response = await addSong(spotifyId, null, null);
    fetchPlaylist();
    fetchSpotifyTrackInfo();
    return await response.json();
  }

  async function startPlay() {
    try {
      await play(spotifyTrackInfo.map(track => track.uri));
    } catch (error) {
      console.error('error while playing');
      console.error(error);
    }
  }

  async function doShuffle() {
    setSpotifyTrackInfo(shuffle(spotifyTrackInfo));
  }

  useEffect(
    () => {
      fetchPlaylist();
      fetchSpotifyTrackInfo();
    },
    [fetchPlaylist, fetchSpotifyTrackInfo]
  );

  return {
    playlist,
    spotifyTrackInfo,
    fetchPlaylist,
    fetchSpotifyTrackInfo,
    appendSong,
    startPlay,
    doShuffle
  };
}

export {
  useMembersList,
  useSpotifyPlayer,
  useSongSearch,
  usePlaybackInformation,
  useSpotifyPlaylist
};
