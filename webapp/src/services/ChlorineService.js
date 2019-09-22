async function joinRoom(roomID, name) {
  return await fetch('api/member', {
    method: 'POST',
    credentials: 'include',
    body: JSON.stringify({ name: name, room_id: Number(roomID) })
  });
}

async function getToken() {
  const response = await fetch('api/token', {
    credentials: 'include'
  });

  return response.json();
}

async function getMemberInfo() {
  const response = await fetch('api/member', {
    credentials: 'include'
  });

  return response.json();
}

async function getRoomMembers() {
  const response = await fetch('api/room/members', {
    credentials: 'include'
  });

  return response.json();
}

async function getPlaybackInformation() {
  const response = await fetch('api/me/player/', {
    credentials: 'include'
  });

  return response.json();
}

async function searchTracks(query) {
  const url = new URL('/api/search', window.location.href);
  const params = { q: query };
  url.search = new URLSearchParams(params);

  const response = await fetch(url, {
    credentials: 'include'
  });

  return response.json();
}

async function getDevicesInformation() {
  const response = await fetch('api/me/player/devices', {
    credentials: 'include'
  });

  return response.json();
}

async function transferPlayback(deviceId, play) {
  const response = await fetch('api/me/player/', {
    method: 'PUT',
    credentials: 'include',
    body: JSON.stringify({
      device_id: deviceId,
      play: play
    })
  });

  return response.json();
}

async function retrieveRoomSongs() {
  const response = await fetch('api/room/songs', {
    credentials: 'include'
  });

  return response.json();
}

async function retrieveRoomsSongsFromSpotify() {
  const response = await fetch('api/room/songs/spotify', {
    credentials: 'include'
  });

  return response.json();
}

async function addSong(spotifyId, prevSongId, nextSongId) {
  console.log(`adding song ${spotifyId}, ${prevSongId}, ${nextSongId}`);
  const response = await fetch('api/room/songs', {
    credentials: 'include',
    method: 'POST',
    body: JSON.stringify({
      spotify_id: spotifyId,
      previous_song_id: prevSongId,
      next_song_id: nextSongId
    })
  });
  return response.json();
}

async function updateSong(song_id, values) {
  console.log('updating song');
  const response = await fetch('api/room/songs', {
    credentials: 'include',
    method: 'PUT',
    body: JSON.stringify({
      id: song_id,
      spotify_id: values.spotifyId,
      previous_song_id: values.prevSongId,
      next_song_id: values.nextSongId
    })
  });

  return response.json();
}

async function play(spotify_uris) {
  console.log('will play', spotify_uris);
  const response = await fetch('api/play', {
    credentials: 'include',
    method: 'POST',
    body: JSON.stringify({
      uris: spotify_uris
    })
  });

  return response.json();
}

export {
  joinRoom,
  getToken,
  getMemberInfo,
  getRoomMembers,
  searchTracks,
  getPlaybackInformation,
  getDevicesInformation,
  transferPlayback,
  retrieveRoomSongs,
  retrieveRoomsSongsFromSpotify,
  addSong,
  updateSong,
  play
};
