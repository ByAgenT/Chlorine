import React, { useCallback, useEffect, useState } from 'react';
import { withRouter } from 'react-router-dom';
import { useMembersList } from '../hooks/membership';
import { useSpotifyPlaylist } from '../hooks/player';
import RootPartyContainer from '../containers/RootPartyContainer';
import PartyContainer from '../containers/PartyContainer';
import Panel from '../components/common/Panel';
import SpotifyPlaylist from '../components/SpotifyPlaylist';
import MembersList from '../components/MembersList';
import { useTranslation } from 'react-i18next';
import AddSongsModal from '../containers/AddSongsModal';
import useChlorineWebSocket from '../hooks/websocket';

const PartyPage: React.FC = () => {
  const { t } = useTranslation();
  const [members, updateMembers] = useMembersList();
  const [isModalShowed, setModalShowed] = useState<boolean>(false);
  const webSocketConnection = useChlorineWebSocket();
  const {
    spotifyTrackInfo,
    appendSong,
    fetchPlaylist,
    fetchSpotifyTrackInfo,
  } = useSpotifyPlaylist();

  const updatePlaylist = useCallback(() => {
    return Promise.all([fetchPlaylist(), fetchSpotifyTrackInfo()]);
  }, [fetchPlaylist, fetchSpotifyTrackInfo]);

  useEffect(() => {
    webSocketConnection.onBroadcast('SongAdded', () => {
      updatePlaylist();
    });
  }, [webSocketConnection, updatePlaylist]);

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
      <AddSongsModal isShowed={isModalShowed} onClose={setModalShowed} onSongAdd={appendSong} />
    </RootPartyContainer>
  );
};

export default withRouter(PartyPage);
