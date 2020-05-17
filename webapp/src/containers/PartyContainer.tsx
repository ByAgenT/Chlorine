import styled from 'styled-components';
import { down, up } from 'styled-breakpoints';

interface PartyContainerProps {
  direction?: string;
  horizontalAlign?: string;
}

const PartyContainer = styled.div<PartyContainerProps>`
  display: flex;
  height: 40rem;
  flex-grow: 1;
  align-items: center;

  ${down('desktop')} {
    flex-direction: column;
  }

  ${up('desktop')} {
    flex-direction: ${(props) => props.direction || 'row'};
    margin: 10px;
  }

  & > * {
    margin: 1rem;
    width: 40rem;

    ${up('desktop')} {
      width: 50rem;
    }
  }
`;

export default PartyContainer;
