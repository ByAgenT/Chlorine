import { withRouter } from 'react-router-dom';
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
import TextInput from '../components/common/TextInput';
import Modal from '../components/common/Modal';
import SongSearchResultList from '../components/SongSearchResultList';

const PlayerPage: React.FC = () => {
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
            console.log(`transferring to ${chlorine[0].id}`);
            await chlorineService.transferPlayback(chlorine[0].id, true);
          }
        } catch (error) {
          console.error(error);
        }
      });
    }
  }

  useEffect(claimPlayback);

  return (
    <RootPartyContainer>
      <PartyContainer direction='column'>
        <Panel name='Playlist'>
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
        <Panel name='Members'>
          <MembersList members={members} onUpdate={updateMembers} />
        </Panel>
        <Panel name='Player'>{<Player player={player} playback={playback} />}</Panel>
      </PartyContainer>
      <Modal display={[isModalShowed, setModalShowed]}>
        <h1>Search Songs</h1>
        <TextInput placeholder='Enter Track' onChange={onSearchModalChange} />
        <SongSearchResultList onSongAdd={appendSong} songs={searchResult} />
      </Modal>
    </RootPartyContainer>
  );
};

export default withRouter(PlayerPage);
