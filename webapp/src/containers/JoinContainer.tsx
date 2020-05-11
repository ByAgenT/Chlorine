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

  ${down('tablet')} {
    flex-direction: column;
  }

  ${up('desktop')} {
    flex-direction: 'row';
    margin: 10px;
  }
`;

const JoinInnerContainer = styled.div`
  min-height: 35rem;

  ${up('desktop')} {
    min-width: 35rem;
  }
`;

export default JoinContainer;
