import React from 'react';
import styled from 'styled-components';

import LinkButton from '../LinkButton';

const SongSearchResultList = ({ songs, onSongAdd }) => (
  <ListContainer>
    <SongsList>
      {songs.map(song => (
        <LinkButton
          onClick={() => {
            console.log(`Add ${song.id}`);
            onSongAdd(song.id);
          }}
          key={song.id}
        >
          {song.artists.map(artist => artist.name).join(', ')} - {song.name}
        </LinkButton>
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
  padding-left: 1em;

  & > ${LinkButton} {
    margin: 0;
    padding-top: 0.5em;
    padding-bottom: 0.5em;
  }
`;

export default SongSearchResultList;
