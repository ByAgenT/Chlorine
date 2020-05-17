import * as React from 'react';
import styled from 'styled-components';
import { SpotifyTrack } from '../models/chlorine';
import ListItem from './common/ListItem';
import { AlignedCenterFlex } from '../containers/common/AlignedCenterFlex';

interface SongSearchResultListProps {
  songs: SpotifyTrack[];
  onSongAdd: (id: string) => void;
}

const SongSearchResultList: React.FC<SongSearchResultListProps> = ({ songs, onSongAdd }) => (
  <ListContainer>
    <SongsList>
      {songs.map((song) => (
        <SongSearchItem
          img={song.album.images[0].url}
          onClick={() => {
            onSongAdd(song.id);
          }}
          title={song.name}
          artist={song.artists.map((artist) => artist.name).join(', ')}
          key={song.id}
        />
      ))}
    </SongsList>
  </ListContainer>
);

const ListContainer = styled.div``;

const SongsList = styled.div`
  display: flex;
  align-items: baseline;
  flex-direction: column;
  padding-top: 2em;
  padding-left: 0.3em;
  overflow-y: scroll;
  height: 25rem;
`;

interface SongSearchItemProps {
  onClick(): void;

  img: string;
  title: string;
  artist: string;
}

const SongSearchItem: React.FC<SongSearchItemProps> = ({ img, title, artist, onClick }) => {
  return (
    <SearchItemContainer onClick={onClick}>
      <AlignedCenterFlex>
        <TrackImage src={img} />
        <TrackDescriptionContainer>
          <TrackTitle>{title}</TrackTitle>
          <TrackArtist>{artist}</TrackArtist>
        </TrackDescriptionContainer>
      </AlignedCenterFlex>
    </SearchItemContainer>
  );
};

const TrackImage = styled.img`
  width: 40px;
  height: 40px;
`;

const SearchItemContainer = styled(ListItem)`
  min-height: 4rem;
  justify-content: space-between;
  user-select: none;
  width: 100%;
  padding: 0;
  &:hover {
    background-color: #6f7a7d;
  }
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

export default SongSearchResultList;
