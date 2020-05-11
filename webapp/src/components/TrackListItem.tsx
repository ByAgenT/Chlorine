import * as React from 'react';
import styled from 'styled-components';
import ListItem from './common/ListItem';

interface TrackListItemProps {
  img: string;
  title: string;
  artist: string;
  duration: string;
}

const TrackListItem: React.FC<TrackListItemProps> = ({ img, title, artist, duration }) => (
  <TrackListItemContainer>
    <TrackListItemInnerContainer>
      <TrackImage src={img} />
      <TrackDescriptionContainer>
        <TrackTitle>{title}</TrackTitle>
        <TrackArtist>{artist}</TrackArtist>
      </TrackDescriptionContainer>
    </TrackListItemInnerContainer>
    <RightContainer>
      <TrackDuration>{duration}</TrackDuration>
      <TrackDeleteButton
        onClick={() => {



          alert('he');
        }}
      />
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

const TrackListItemInnerContainer = styled.div`
  display: flex;
  align-items: center;
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
