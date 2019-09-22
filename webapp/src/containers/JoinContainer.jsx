import React from 'react';
import styled from 'styled-components';
import breakpoint from 'styled-components-breakpoint';

const JoinContainer = props => (
  <JoinOuterContainer>
    <JoinInnerContainer>{props.children}</JoinInnerContainer>
  </JoinOuterContainer>
);

const JoinOuterContainer = styled.div`
  display: flex;
  min-height: 35rem;
  justify-content: center;

  ${breakpoint('mobile')`
    flex-direction: column;
  `}

  ${breakpoint('desktop')`
    flex-direction: ${props => props.direction || 'row'};
    margin: 10px;
  `}
`;

const JoinInnerContainer = styled.div`
  min-height: 35rem;

  ${breakpoint('desktop')`
    min-width: 35rem;
  `}
`;

export default JoinContainer;
