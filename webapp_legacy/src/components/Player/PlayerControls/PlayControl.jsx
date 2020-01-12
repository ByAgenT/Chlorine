import styled from 'styled-components';
import Control from './Control';

const PlayControl = styled(Control)`
  background-image: url('icons/52px/play-white.png');

  &:hover {
    background-image: url('icons/52px/play-green.png');
  }
`;

export default PlayControl;
