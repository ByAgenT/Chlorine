import styled from 'styled-components';
import { down, up } from 'styled-breakpoints';

const Control = styled.div`
  border-style: none;
  background-size: cover;
  background-repeat: no-repeat;
  ${up('desktop')} {
    width: 52px;
    height: 52px;
  }

  ${down('tablet')} {
    width: 26px;
    height: 26px;
  }
`;

const BackControl = styled(Control)`
  background-image: url('icons/52px/start-white.png');
  &:hover {
    background-image: url('icons/52px/start-green.png');
  }
`;

const NextControl = styled(Control)`
  background-image: url('icons/52px/end-white.png');
  &:hover {
    background-image: url('icons/52px/end-green.png');
  }
`;

const PauseControl = styled(Control)`
  background-image: url('icons/52px/pause-white.png');
  &:hover {
    background-image: url('icons/52px/pause-green.png');
  }
`;

const PlayControl = styled(Control)`
  background-image: url('icons/52px/play-white.png');
  &:hover {
    background-image: url('icons/52px/play-green.png');
  }
`;

export { Control, BackControl, NextControl, PauseControl, PlayControl };
