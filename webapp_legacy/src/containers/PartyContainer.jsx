import styled from 'styled-components';
import breakpoint from 'styled-components-breakpoint';

const PartyContainer = styled.div`
  display: flex;
  min-height: 35rem;
  flex-grow: 1;

  ${breakpoint('mobile')`
    flex-direction: column;
  `}

  ${breakpoint('desktop')`
    flex-direction: ${props => props.direction || 'row'};
    margin: 10px;
  `}
`;

export default PartyContainer;
