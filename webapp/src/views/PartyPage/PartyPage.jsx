import debounce from 'lodash/debounce';
import React, { useEffect, useState } from 'react';
import { withRouter } from 'react-router-dom';
import MembersList from '../../components/MembersList';
import Modal from '../../components/Modal';
import Panel from '../../components/Panel';
import Player from '../../components/Player';
import SongSearchResultList from '../../components/SongSearchResultList';
import SpotifyPlaylist from '../../components/SpotifyPlaylist/SpotifyPlaylist';
import TextInput from '../../components/TextInput';
import PartyContainer from '../../containers/PartyContainer';
import RootPartyContainer from '../../containers/RootPartyContainer';
import {
  getDevicesInformation,
  transferPlayback
} from '../../services/ChlorineService';
import {
  useMembersList,
  usePlaybackInformation,
  useSongSearch,
  useSpotifyPlayer,
  useSpotifyPlaylist
} from './hooks';

const PartyPage = () => {
  const player = useSpotifyPlayer();
  const [members, updateMembers] = useMembersList();
  const playback = usePlaybackInformation(player);
  const [isModalShowed, setModalShowed] = useState(false);
  const { searchResult, setSongQuery } = useSongSearch();
  const {
    spotifyTrackInfo,
    fetchPlaylist,
    fetchSpotifyTrackInfo,
    appendSong,
    startPlay,
    doShuffle
  } = useSpotifyPlaylist();

  useEffect(claimPlayback);

  const updateSongQuery = debounce(event => {
    setSongQuery(event.target.value);
  }, 200);

  function onSearchModalChange(event) {
    event.persist();
    updateSongQuery(event);
  }

  function updatePlaylist() {
    fetchPlaylist();
    fetchSpotifyTrackInfo();
  }

  function claimPlayback() {
    if (player) {
      player.onPlayerReady(async () => {
        try {
          const devices = await getDevicesInformation();
          const chlorine = devices.filter(device => device.name === 'Chlorine');
          if (chlorine[0] !== undefined) {
            console.log(`transferring to ${chlorine[0].id}`);
            await transferPlayback(chlorine[0].id, true);
          }
        } catch (error) {
          console.error(error);
        }
      });
    }
  }

  return (
    <RootPartyContainer>
      <PartyContainer direction="column">
        <Panel name="Playlist">
          <SpotifyPlaylist
            onAddSongClick={() => setModalShowed(!isModalShowed)}
            onStartPlay={startPlay}
            onShuffle={doShuffle}
            playlist={spotifyTrackInfo}
            onUpdate={updatePlaylist}
          />
        </Panel>
      </PartyContainer>
      <PartyContainer direction="column">
        <Panel name="Members">
          <MembersList members={members} onUpdate={updateMembers}/>
        </Panel>
        <Panel name="Player">
          <Player player={player} playback={playback} />
        </Panel>
      </PartyContainer>
      <Modal display={[isModalShowed, setModalShowed]}>
        <h1>Search Songs</h1>
        <TextInput placeholder="Enter Track" onChange={onSearchModalChange} />
        <SongSearchResultList onSongAdd={appendSong} songs={searchResult} />
      </Modal>
    </RootPartyContainer>
  );
};

export default withRouter(PartyPage);
