import * as React from 'react';
import { useEffect, useState } from 'react';
import styled from 'styled-components';
import SongLine from './SongLine';
import { BackControl, NextControl, PauseControl, PlayControl } from './PlayerControls';
import { down, up } from 'styled-breakpoints';
import { SpotifyPlayer } from '../services/spotifyPlaybackService';
import { SimplePlaybackInformation } from '../hooks/player';

interface PlayerProps {
  player: SpotifyPlayer | null;
  playback: SimplePlaybackInformation;
}

const Player: React.FC<PlayerProps> = ({ player, playback }) => {
  const [playerPlayback, setPlayerPlayback] = useState<Spotify.PlaybackState | null>(null);
  useEffect(() => {
    (async function () {
      if (player) {
        const state = await player.getCurrentState();
        setPlayerPlayback(state);
      }
    })();
  });

  const duration = playerPlayback ? playerPlayback.duration : playback.duration;
  const now = playerPlayback ? playerPlayback.position : playback.now;
  return (
    <PlayerContainer>
      <TrackCover src={playback.albumCoverURL} />
      <ControlContainer>
        <SongTitle>
          <span>{playback.artistTitle}</span>
          {' – '}
          <span>{playback.songTitle}</span>
        </SongTitle>
        <SongLine duration={duration} now={now} />
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
};

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
  font-size: 1.6rem;
  text-align: center;
  height: 4.5rem;
  ${down('desktop')} {
    font-size: 1rem;
  }

  ${up('desktop')} {
    font-size: 1.6rem;
  }
`;

export default Player;
