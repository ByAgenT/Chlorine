import * as React from 'react';
import styled from 'styled-components';
import { toTrackTime } from '../utils/time';

interface SongLineProps {
  now: number;
  duration: number;
}

const SongLine: React.FC<SongLineProps> = (props) => (
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
`;

const SongLineTotal = styled.div`
  background-color: gray;
  width: 100%;
  margin: 0 10px;
  height: 5px;
  border-radius: 0.3em;
`;

interface SongLineFgProps {
  progress: number;
}

const SongLineFg = styled.div.attrs<SongLineFgProps>((props) => ({
  style: { width: props.progress ? `${props.progress}%` : '0' },
}))<SongLineFgProps>`
  position: inherit;
  background-color: white;
  height: 5px;
  border-radius: 0.3em;
`;

export default SongLine;
