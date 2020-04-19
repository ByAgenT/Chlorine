import * as React from 'react';
import styled from 'styled-components';
import { only } from 'styled-breakpoints';

const JoinContainer: React.FC = (props) => (
  <JoinOuterContainer>
    <JoinInnerContainer>{props.children}</JoinInnerContainer>
  </JoinOuterContainer>
);

const JoinOuterContainer = styled.div`
  display: flex;
  min-height: 35rem;
  justify-content: center;

  ${only('tablet')} {
    flex-direction: column;
  }

  ${only('desktop')} {
    flex-direction: 'row';
    margin: 10px;
  }
`;

const JoinInnerContainer = styled.div`
  min-height: 35rem;

  ${only('desktop')} {
    min-width: 35rem;
  }
`;

export default JoinContainer;
