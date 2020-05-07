import debounce from 'lodash/debounce';
import React, { useState } from 'react';
import { withRouter } from 'react-router-dom';
import { useMembersList } from '../hooks/membership';
import { useSongSearch } from '../hooks/search';
import { useSpotifyPlaylist } from '../hooks/player';
import RootPartyContainer from '../containers/RootPartyContainer';
import PartyContainer from '../containers/PartyContainer';
import Panel from '../components/common/Panel';
import SpotifyPlaylist from '../components/SpotifyPlaylist';
import MembersList from '../components/MembersList';
import Modal from '../components/common/Modal';
import TextInput from '../components/common/TextInput';
import SongSearchResultList from '../components/SongSearchResultList';

const PartyPage = () => {
  const [members, updateMembers] = useMembersList();
  const [isModalShowed, setModalShowed] = useState(false);
  const { searchResult, setSongQuery } = useSongSearch();
  const {
    spotifyTrackInfo,
    appendSong,
    fetchPlaylist,
    fetchSpotifyTrackInfo,
  } = useSpotifyPlaylist();

  const updateSongQuery = debounce((event) => {
    setSongQuery(event.target.value);
  }, 200);

  function onSearchModalChange(event) {
    event.persist();
    updateSongQuery(event);
  }

  function updatePlaylist() {
    Promise.all([fetchPlaylist(), fetchSpotifyTrackInfo()]);
  }

  return (
    <RootPartyContainer>
      <PartyContainer direction='column'>
        <Panel name='Playlist'>
          <SpotifyPlaylist
            onAddSongClick={() => setModalShowed(!isModalShowed)}
            playlist={spotifyTrackInfo}
            onUpdate={updatePlaylist}
          />
        </Panel>
      </PartyContainer>
      <PartyContainer direction='column'>
        <Panel name='Members'>
          <MembersList members={members} onUpdate={updateMembers} />
        </Panel>
      </PartyContainer>
      <Modal display={[isModalShowed, setModalShowed]}>
        <h1>Search Songs</h1>
        <TextInput placeholder='Enter Track' onChange={onSearchModalChange} />
        <SongSearchResultList onSongAdd={appendSong} songs={searchResult} />
      </Modal>
    </RootPartyContainer>
  );
};

export default withRouter(PartyPage);
