import { RouteComponentProps, withRouter } from 'react-router-dom';
import RootPartyContainer from '../containers/RootPartyContainer';
import * as React from 'react';
import { useEffect, useState } from 'react';
import PartyContainer from '../containers/PartyContainer';
import Panel from '../components/common/Panel';
import { usePlaybackInformation, useSpotifyPlayer, useSpotifyPlaylist } from '../hooks/player';
import { useMembersList } from '../hooks/membership';
import { useSongSearch } from '../hooks/search';
import { ChlorineService } from '../services/chlorineService';
import debounce from 'lodash/debounce';
import MembersList from '../components/MembersList';
import Player from '../components/Player';
import SpotifyPlaylist from '../components/SpotifyPlaylist';
import { useTranslation } from 'react-i18next';
import AddSongsModal from '../containers/AddSongsModal';
import styled from 'styled-components';
import Settings from '../components/Settings';
import { Member } from '../models/chlorine';

interface PlayerPageProps extends RouteComponentProps {
  member: Member | null;
}

const PlayerPage: React.FC<PlayerPageProps> = ({ member }) => {
  const { t } = useTranslation();
  const player = useSpotifyPlayer();
  const [members, updateMembers] = useMembersList();
  const playback = usePlaybackInformation(player);
  const [isModalShowed, setModalShowed] = useState<boolean>(false);
  const { searchResult, setSongQuery } = useSongSearch();
  const {
    spotifyTrackInfo,
    fetchPlaylist,
    fetchSpotifyTrackInfo,
    appendSong,
    startPlay,
    doShuffle,
  } = useSpotifyPlaylist();

  const updateSongQuery = debounce((event: React.ChangeEvent<HTMLTextAreaElement>) => {
    setSongQuery(event.target.value);
  }, 200);

  function onSearchModalChange(event: React.ChangeEvent<HTMLTextAreaElement>) {
    event.persist();
    updateSongQuery(event);
  }

  function updatePlaylist() {
    return Promise.all([fetchPlaylist(), fetchSpotifyTrackInfo()]);
  }

  function claimPlayback() {
    if (player) {
      player.onPlayerReady(async () => {
        const chlorineService = new ChlorineService();
        try {
          const devices = await chlorineService.getDevicesInformation();
          const chlorine = devices.filter((device) => device.name === 'Chlorine');
          if (chlorine[0] !== undefined) {
            await chlorineService.transferPlayback(chlorine[0].id, true);
          }
        } catch (error) {
          console.error(error);
        }
      });
    }
  }

  useEffect(claimPlayback);

  const roomId = member ? member.roomId : NaN;
  return (
    <RootPartyContainer>
      <PartyContainer direction='column'>
        <Panel name={t('playlist')}>
          <SpotifyPlaylist
            onAddSongClick={() => setModalShowed(!isModalShowed)}
            onStartPlay={startPlay}
            onShuffle={doShuffle}
            playlist={spotifyTrackInfo}
            onUpdate={updatePlaylist}
          />
        </Panel>
      </PartyContainer>
      <PartyContainer direction='column'>
        <Panel name={t('members')}>
          <ManagementContainer>
            <MembersList members={members} onUpdate={updateMembers} />
            <Settings roomId={roomId} />
          </ManagementContainer>
        </Panel>
        <Panel name={t('player')}>{<Player player={player} playback={playback} />}</Panel>
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

const ManagementContainer = styled.div`
  display: flex;
  flex-direction: row;
  height: 100%;
`;

export default withRouter(PlayerPage);
