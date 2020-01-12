import styled from 'styled-components';
import Control from './Control';

const BackControl = styled(Control)`
  background-image: url('icons/52px/start-white.png');

  &:hover {
    background-image: url('icons/52px/start-green.png');
  }
`;

export default BackControl;
