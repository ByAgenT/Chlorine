import * as React from 'react';
import styled from 'styled-components';
import LinkButton from './common/LinkButton';
import TrackListItem from './TrackListItem';
import List from './common/List';
import { useTranslation } from 'react-i18next';
import { PlaylistTrack } from '../hooks/player';

interface SpotifyPlaylistProps {
  onAddSongClick: (event: React.MouseEvent) => void;
  onStartPlay?: (event: React.MouseEvent) => void;
  onShuffle?: (event: React.MouseEvent) => void;
  onDelete?: (track: PlaylistTrack) => void;
  onUpdate: (event: React.MouseEvent) => void;
  playlist: PlaylistTrack[];
}

const SpotifyPlaylist: React.FC<SpotifyPlaylistProps> = ({
  onAddSongClick,
  playlist,
  onStartPlay,
  onShuffle,
  onDelete,
  onUpdate,
}) => {
  const { t } = useTranslation();
  return (
    <SpotifyPlaylistContainer>
      <PlaylistList>
        {playlist &&
          playlist.map((track) => {
            return <TrackListItem key={track.id} track={track} onDelete={onDelete} />;
          })}
      </PlaylistList>
      <PlaylistBottomBar>
        <LinkButton onClick={onAddSongClick}>{t('add_songs')}</LinkButton>
        {onShuffle ? <LinkButton onClick={onShuffle}>{t('shuffle')}</LinkButton> : ''}
        {onStartPlay ? <LinkButton onClick={onStartPlay}>{t('start_play')}</LinkButton> : ''}
        <LinkButton onClick={onUpdate}>{t('refresh')}</LinkButton>
      </PlaylistBottomBar>
    </SpotifyPlaylistContainer>
  );
};

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
  max-height: 35rem;
`;

export default SpotifyPlaylist;
