import * as React from 'react';
import styled from 'styled-components';
import SongLine from './SongLine';
import { BackControl, NextControl, PauseControl, PlayControl } from './PlayerControls';
import { down, only, up } from 'styled-breakpoints';
import { SpotifyPlayer } from '../services/spotifyPlaybackService';
import { SimplePlaybackInformation } from '../hooks/player';

interface PlayerProps {
  player: SpotifyPlayer;
  playback: SimplePlaybackInformation;
}

const Player: React.FC<PlayerProps> = ({ player, playback }) => (
  <PlayerContainer>
    <TrackCover src={playback.albumCoverURL} />
    <ControlContainer>
      <SongTitle>
        <span>{playback.artistTitle}</span>
        {' – '}
        <span>{playback.songTitle}</span>
      </SongTitle>
      <SongLine duration={playback.duration} now={playback.now} />
      <ButtonsContainer>
        <BackControl
          onClick={() => {
            player.previousTrack();
          }}
        />
        <PlayControl
          onClick={() => {
            player.resume();
          }}
        />
        <PauseControl
          onClick={() => {
            player.pause();
          }}
        />
        <NextControl
          onClick={() => {
            player.nextTrack();
          }}
        />
      </ButtonsContainer>
    </ControlContainer>
  </PlayerContainer>
);

const ControlContainer = styled.div`
  display: flex;
  flex-grow: 1;
  justify-content: space-evenly;
  flex-direction: column;
  padding: 0 2em;
  align-items: center;
`;

const PlayerContainer = styled.div`
  display: flex;
  justify-content: space-evenly;
  color: white;
  padding: 20px;
`;

const TrackCover = styled.img`
  width: 170px;
  height: 170px;
  ${down('tablet')} {
    width: 50px;
    height: 50px;
  }

  ${up('desktop')} {
    width: 170px;
    height: 170px;
  }
`;

const ButtonsContainer = styled.div`
  display: flex;
  justify-content: space-evenly;
  width: 100%;
`;

const SongTitle = styled.div`
  font-weight: 600;
  margin-bottom: 1em;
  font-size: 1.6rem;
  ${down('tablet')} {
    font-size: 1rem;
  }

  ${up('desktop')} {
    font-size: 1.6rem;
  }
`;

export default Player;
