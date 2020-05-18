import styled from 'styled-components';
import { up } from 'styled-breakpoints';

import PartyContainer from './PartyContainer';

const RootPartyContainer = styled(PartyContainer)`
  ${up('desktop')} {
    margin: 0.5rem auto;
    width: max-content;
  }
`;

export default RootPartyContainer;
