import styled from 'styled-components';
import breakpoint from 'styled-components-breakpoint';

import PartyContainer from './PartyContainer';

const RootPartyContainer = styled(PartyContainer)`
  ${breakpoint('desktop')`
    margin-left: 20px;
    margin-top: 10px;
    margin-right: 20px;
  `}
`;

export default RootPartyContainer;
