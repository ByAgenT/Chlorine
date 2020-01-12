import React from 'react';
import styled from 'styled-components';
import PlayControl from './PlayerControls/PlayControl';
import BackControl from './PlayerControls/BackControl';
import PauseControl from './PlayerControls/PauseControl';
import NextControl from './PlayerControls/NextControl';
import ControlContainer from './ControlContainer';
import SongLine from './PlayerControls/SongLine';
import breakpoint from 'styled-components-breakpoint';

const Player = ({ player, playback }) => (
  <PlayerContainer>
    <TrackCover src={playback.albumCoverURL} />
    <ControlContainer>
      <SongTitle>
        <span>{playback.artistTitle}</span>
        {' – '}
        <span>{playback.songTitle}</span>
      </SongTitle>
      <SongLine
        duration={playback.duration}
        now={playback.now}
      />
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

const PlayerContainer = styled.div`
  display: flex;
  justify-content: space-evenly;
  color: white;
  padding: 20px;
`;

const TrackCover = styled.img`
  ${breakpoint('mobile')`
    width: 50px;
    height: 50px;
  `}

  ${breakpoint('desktop')`
    width: 170px;
    height: 170px;
  `}
`;

const ButtonsContainer = styled.div`
  display: flex;
  justify-content: space-evenly;
  width: 100%;
`;

const SongTitle = styled.div`
  font-weight: 600;
  margin-bottom: 1em;

  ${breakpoint('mobile')`
    font-size: 1rem;
  `}

  ${breakpoint('desktop')`
    font-size: 1.6rem;
  `}
`;

export default Player;
