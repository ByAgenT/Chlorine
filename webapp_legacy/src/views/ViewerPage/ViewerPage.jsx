import debounce from 'lodash/debounce';
import React, { useState } from 'react';
import { withRouter } from 'react-router-dom';
import MembersList from '../../components/MembersList';
import Modal from '../../components/Modal';
import Panel from '../../components/Panel';
import SongSearchResultList from '../../components/SongSearchResultList';
import SpotifyPlaylist from '../../components/SpotifyPlaylist/SpotifyPlaylist';
import TextInput from '../../components/TextInput';
import PartyContainer from '../../containers/PartyContainer';
import RootPartyContainer from '../../containers/RootPartyContainer';
import {
  useMembersList,
  useSongSearch,
  useSpotifyPlaylist
} from '../PartyPage/hooks';

const PartyPage = () => {
  const [members, updateMembers] = useMembersList();
  const [isModalShowed, setModalShowed] = useState(false);
  const { searchResult, setSongQuery } = useSongSearch();
  const { spotifyTrackInfo, appendSong, fetchPlaylist, fetchSpotifyTrackInfo } = useSpotifyPlaylist();

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

  return (
    <RootPartyContainer>
      <PartyContainer direction="column">
        <Panel name="Playlist">
          <SpotifyPlaylist
            onAddSongClick={() => setModalShowed(!isModalShowed)}
            playlist={spotifyTrackInfo}
            onUpdate={updatePlaylist}
          />
        </Panel>
      </PartyContainer>
      <PartyContainer direction="column">
        <Panel name="Members">
          <MembersList members={members} onUpdate={updateMembers}/>
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
