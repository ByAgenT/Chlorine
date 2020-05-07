import styled from 'styled-components';
import { only } from 'styled-breakpoints';

import PartyContainer from './PartyContainer';

const RootPartyContainer = styled(PartyContainer)`
  ${only('desktop')} {
    margin-left: 20px;
    margin-top: 10px;
    margin-right: 20px;
  }
`;

export default RootPartyContainer;
