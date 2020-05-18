import * as React from 'react';
import styled from 'styled-components';
import { up, down } from 'styled-breakpoints';

const JoinContainer: React.FC = (props) => (
  <JoinOuterContainer>
    <JoinInnerContainer>{props.children}</JoinInnerContainer>
  </JoinOuterContainer>
);

const JoinOuterContainer = styled.div`
  display: flex;
  min-height: 35rem;
  justify-content: center;
  width: max-content;
  margin: 1.5rem auto;
`;

const JoinInnerContainer = styled.div`
  ${up('desktop')} {
    min-width: 35rem;
  }
`;

export default JoinContainer;
