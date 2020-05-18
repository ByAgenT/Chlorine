import * as React from 'react';
import styled from 'styled-components';
import LinkButton from './common/LinkButton';
import TrackListItem from './TrackListItem';
import List from './common/List';
import { SpotifyTrack } from '../models/chlorine';
import { useTranslation } from 'react-i18next';
import { toTrackTime } from '../utils/time';

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
}) => {
  const { t } = useTranslation();
  return (
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
                    track.album.images.filter((image) => image.width > 50 && image.width < 100)[0]
                      .url
                  }
                  duration={toTrackTime(track.durationMs)}
                />
              );
            })
          : ''}
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
