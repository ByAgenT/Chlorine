import styled from 'styled-components';
import Control from './Control';

const PauseControl = styled(Control)`
  background-image: url('icons/52px/pause-white.png');

  &:hover {
    background-image: url('icons/52px/pause-green.png');
  }
`;

export default PauseControl;
