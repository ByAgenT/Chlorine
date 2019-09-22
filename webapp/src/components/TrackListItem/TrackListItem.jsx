import React from 'react';
import styled from 'styled-components';

import ListItem from '../ListItem';

const TrackListItem = ({img, title, artist, duration}) => (
  <TrackListItemContainer>
    <TrackListItemInnerContainer>
      <TrackImage src={img} />
      <TrackDescriptionContainer>
        <TrackTitle>{title}</TrackTitle>
        <TrackArtist>{artist}</TrackArtist>
      </TrackDescriptionContainer>
    </TrackListItemInnerContainer>
    <TrackDuration>{duration}</TrackDuration>
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

const TrackDuration = styled.span`
  font-size: 1.3rem;
`;

export default TrackListItem;
