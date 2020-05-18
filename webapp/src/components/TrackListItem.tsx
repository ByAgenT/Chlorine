import * as React from 'react';
import styled from 'styled-components';
import ListItem from './common/ListItem';
import { AlignedCenterFlex } from '../containers/common/AlignedCenterFlex';
import { PlaylistTrack } from '../hooks/player';
import { toTrackTime } from '../utils/time';

interface TrackListItemProps {
  track: PlaylistTrack;
  onDelete?: (track: PlaylistTrack) => void;
}

const TrackListItem: React.FC<TrackListItemProps> = ({ track, onDelete }) => (
  <TrackListItemContainer>
    <AlignedCenterFlex>
      <TrackImage
        src={track.album.images.filter((image) => image.width > 50 && image.width < 100)[0].url}
      />
      <TrackDescriptionContainer>
        <TrackTitle>{track.name}</TrackTitle>
        <TrackArtist>{track.artists.map((artist) => artist.name).join(', ')}</TrackArtist>
      </TrackDescriptionContainer>
    </AlignedCenterFlex>
    <RightContainer>
      <TrackDuration>{toTrackTime(track.durationMs)}</TrackDuration>
      {onDelete && (
        <TrackDeleteButton
          onClick={() => {
            onDelete(track);
          }}
        />
      )}
    </RightContainer>
  </TrackListItemContainer>
);

const TrackListItemContainer = styled(ListItem)`
  min-height: 4rem;
  justify-content: space-between;
  &:hover {
    background-color: #222326;
  }
`;

const TrackImage = styled.img`
  width: 50px;
  height: 50px;
`;

const TrackDescriptionContainer = styled.div`
  display: flex;
  flex-direction: column;
  margin: 0 1rem;
  font-size: 1.1rem;
`;

const TrackTitle = styled.span`
  font-size: 1.2rem;
`;

const TrackArtist = styled.span`
  font-size: 1rem;
`;

const RightContainer = styled.div`
  display: flex;
  flex-direction: row;
  align-items: center;

  & > * {
    margin-left: 0.3rem;
  }
`;

const TrackDuration = styled.span`
  font-size: 1.3rem;
`;

const TrackDeleteButton = styled.div`
  background-image: url('icons/delete-24px.png');
  background-size: contain;
  height: 1.2rem;
  width: 1.2rem;
`;

export default TrackListItem;
