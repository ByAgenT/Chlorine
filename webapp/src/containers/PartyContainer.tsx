import styled from 'styled-components';
import { down, up } from 'styled-breakpoints';

interface PartyContainerProps {
  direction?: string;
}

const PartyContainer = styled.div<PartyContainerProps>`
  display: flex;
  min-height: 35rem;
  flex-grow: 1;
  ${down('tablet')} {
    flex-direction: column;
  }

  ${up('desktop')} {
    flex-direction: ${(props) => props.direction || 'row'};
    margin: 10px;
  }
`;

export default PartyContainer;
