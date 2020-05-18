import { RouteComponentProps, withRouter } from 'react-router-dom';
import RootPartyContainer from '../containers/RootPartyContainer';
import * as React from 'react';
import { useCallback, useEffect, useState } from 'react';
import PartyContainer from '../containers/PartyContainer';
import Panel from '../components/common/Panel';
import { PlaylistTrack, usePlaybackInformation, useSpotifyPlayer, useSpotifyPlaylist } from '../hooks/player';
import { useMembersList } from '../hooks/membership';
import { ChlorineService } from '../services/chlorineService';
import MembersList from '../components/MembersList';
import Player from '../components/Player';
import SpotifyPlaylist from '../components/SpotifyPlaylist';
import { useTranslation } from 'react-i18next';
import AddSongsModal from '../containers/AddSongsModal';
import styled from 'styled-components';
import Settings from '../components/Settings';
import { Member } from '../models/chlorine';
import useChlorineWebSocket from '../hooks/websocket';

interface PlayerPageProps extends RouteComponentProps {
  member: Member | null;
}

const PlayerPage: React.FC<PlayerPageProps> = ({ member }) => {
  const { t } = useTranslation();
  const player = useSpotifyPlayer();
  const [members, updateMembers] = useMembersList();
  const playback = usePlaybackInformation(player);
  const [isModalShowed, setModalShowed] = useState<boolean>(false);
  const webSocketConnection = useChlorineWebSocket();
  const {
    playlist,
    fetchTracks,
    appendSong,
    startPlay,
    doShuffle,
  } = useSpotifyPlaylist();

  const updatePlaylist = useCallback(() => {
    return fetchTracks();
  }, [fetchTracks]);

  const claimPlayback = useCallback(() => {
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
  }, [player]);

  const trackDelete = useCallback(async (track: PlaylistTrack) => {
    await new ChlorineService().deleteSong(track.id);
    await fetchTracks();
  }, [fetchTracks]);

  useEffect(claimPlayback);

  useEffect(() => {
    webSocketConnection.onBroadcast('SongAdded', () => {
      updatePlaylist();
    });
    webSocketConnection.onBroadcast('MemberAdded', () => {
      updateMembers();
    });

    return () => {
      webSocketConnection.removeOnBroadcastListener('SongAdded');
      webSocketConnection.removeOnBroadcastListener('MemberAdded');
    };
  }, [updatePlaylist, webSocketConnection, updateMembers]);

  const roomId = member ? member.roomId : NaN;
  return (
    <RootPartyContainer>
      <PartyContainer direction='column'>
        <Panel name={t('playlist')}>
          <SpotifyPlaylist
            onAddSongClick={() => setModalShowed(!isModalShowed)}
            onStartPlay={startPlay}
            onShuffle={doShuffle}
            playlist={playlist}
            onUpdate={updatePlaylist}
            onDelete={trackDelete}
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
      <AddSongsModal isShowed={isModalShowed} onClose={setModalShowed} onSongAdd={appendSong} />
    </RootPartyContainer>
  );
};

const ManagementContainer = styled.div`
  display: flex;
  flex-direction: row;
  height: 100%;
`;

export default withRouter(PlayerPage);
