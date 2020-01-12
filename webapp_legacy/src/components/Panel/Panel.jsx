import React from 'react';
import styled from 'styled-components';
import breakpoint from 'styled-components-breakpoint';

const Panel = props => (
  <PanelContainer>
    <PanelHeader>{props.name}</PanelHeader>
    <PanelBody className={props.className}>{props.children}</PanelBody>
  </PanelContainer>
);

const PanelContainer = styled.div`
  border: 1px solid #616467;
  margin: 1rem;
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  background: #111111;
`;

const PanelBody = styled.div`
  flex-grow: 1;
  display: flex;
  flex-direction: column;

  ${breakpoint('mobile')`
    min-height: 0;
  `}

  ${breakpoint('desktop')`
    min-height: 7rem;
  `}
`;

const PanelHeader = styled.div`
  background-color: #292929;
  padding: 10px 15px;
  font-size: 1rem;
  border-bottom: 1px dashed #616467;
  color: white;
`;

export default Panel;
