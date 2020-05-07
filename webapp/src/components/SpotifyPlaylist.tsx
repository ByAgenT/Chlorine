import * as React from 'react';
import styled from 'styled-components';
import LinkButton from './common/LinkButton';
import TrackListItem from './TrackListItem';
import List from './common/List';
import { SpotifyTrack } from '../models/chlorine';

function toTrackTime(milliseconds: number): string {
  let date = new Date(milliseconds);
  return `${date.getMinutes()}:${date.getSeconds()}`;
}

interface SpotifyPlaylistProps {
  onAddSongClick: (event: React.MouseEvent) => void;
  onStartPlay?: (event: React.MouseEvent) => void;
  onShuffle?: (event: React.MouseEvent) => void;
  onUpdate: (event: React.MouseEvent) => void;
  playlist: SpotifyTrack[];
}

const SpotifyPlaylist: React.FC<SpotifyPlaylistProps> = ({
  onAddSongClick,
  playlist,
  onStartPlay,
  onShuffle,
  onUpdate,
}) => (
  <SpotifyPlaylistContainer>
    <PlaylistList>
      {playlist
        ? playlist.map((track) => {
            return (
              <TrackListItem
                key={track.id}
                title={track.name}
                artist={track.artists.map((artist) => artist.name).join(', ')}
                img={
                  track.album.images.filter((image) => image.width > 50 && image.width < 100)[0].url
                }
                duration={toTrackTime(track.durationMs)}
              />
            );
          })
        : ''}
    </PlaylistList>
    <PlaylistBottomBar>
      <LinkButton onClick={onAddSongClick}>Add Songs</LinkButton>
      {onShuffle ? <LinkButton onClick={onShuffle}>Shuffle</LinkButton> : ''}
      {onStartPlay ? <LinkButton onClick={onStartPlay}>Start Play</LinkButton> : ''}
      <LinkButton onClick={onUpdate}>Refresh</LinkButton>
    </PlaylistBottomBar>
  </SpotifyPlaylistContainer>
);

const SpotifyPlaylistContainer = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  flex-grow: 1;
`;

const PlaylistBottomBar = styled.div`
  display: flex;
  height: 2.5rem;
  color: white;
  background-color: #292929;
  align-items: center;
`;

const PlaylistList = styled(List)`
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  overflow-y: scroll;
  max-height: 50em;
`;

export default SpotifyPlaylist;
