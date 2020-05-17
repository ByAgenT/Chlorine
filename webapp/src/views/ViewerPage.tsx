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
import { useTranslation } from 'react-i18next';
import AddSongsModal from '../containers/AddSongsModal';

const PartyPage: React.FC = () => {
  const { t } = useTranslation();
  const [members, updateMembers] = useMembersList();
  const [isModalShowed, setModalShowed] = useState<boolean>(false);
  const { searchResult, setSongQuery } = useSongSearch();
  const {
    spotifyTrackInfo,
    appendSong,
    fetchPlaylist,
    fetchSpotifyTrackInfo,
  } = useSpotifyPlaylist();

  const updateSongQuery = debounce((event: React.ChangeEvent<HTMLTextAreaElement>): void => {
    setSongQuery(event.target.value);
  }, 200);

  function onSearchModalChange(event: React.ChangeEvent<HTMLTextAreaElement>): void {
    event.persist();
    updateSongQuery(event);
  }

  function updatePlaylist(): void {
    Promise.all([fetchPlaylist(), fetchSpotifyTrackInfo()]);
  }

  return (
    <RootPartyContainer>
      <PartyContainer direction='column'>
        <Panel name={t('playlist')}>
          <SpotifyPlaylist
            onAddSongClick={() => setModalShowed(!isModalShowed)}
            playlist={spotifyTrackInfo}
            onUpdate={updatePlaylist}
          />
        </Panel>
      </PartyContainer>
      <PartyContainer direction='column'>
        <Panel name={t('members')}>
          <MembersList members={members} onUpdate={updateMembers} />
        </Panel>
      </PartyContainer>
      <AddSongsModal
        isShowed={isModalShowed}
        onClose={setModalShowed}
        onSearchValueChange={onSearchModalChange}
        onSongAdd={appendSong}
        songs={searchResult}
      />
    </RootPartyContainer>
  );
};

export default withRouter(PartyPage);
