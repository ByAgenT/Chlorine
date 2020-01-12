import React from 'react';
import styled from 'styled-components';

function toTrackTime(milliseconds) {
  let date = new Date(milliseconds);
  return `${date.getMinutes()}:${date.getSeconds()}`;
}

const SongLine = props => (
  <SongLineContainer>
    <SongTime>{toTrackTime(props.now)}</SongTime>
    <SongLineTotal>
      <SongLineFg progress={(props.now / props.duration) * 100} />
    </SongLineTotal>
    <SongTime>{toTrackTime(props.duration)}</SongTime>
  </SongLineContainer>
);

const SongLineContainer = styled.div`
  display: flex;
  align-items: center;
  width: 100%;
`;

const SongTime = styled.span`
  font-size: 1.2rem;
  font-weight: 600;
`;

const SongLineTotal = styled.div`
  background-color: gray;
  width: 100%;
  margin: 0 10px;
  height: 5px;
  border-radius: 0.3em;
`;

const SongLineFg = styled.div.attrs(props => ({
  width: props.progress ? `${props.progress}%` : 0
}))`
  position: inherit;
  background-color: white;
  height: 5px;
  border-radius: 0.3em;
  width: ${props => props.width};
`;

export default SongLine;
