import * as React from 'react';
import styled from 'styled-components';
import { down, up } from 'styled-breakpoints';

interface PanelProps extends React.HTMLProps<HTMLInputElement> {
  name: string;
}

const Panel: React.FC<PanelProps> = (props) => (
  <PanelContainer>
    <PanelHeader>{props.name}</PanelHeader>
    <PanelBody className={props.className}>{props.children}</PanelBody>
  </PanelContainer>
);

const PanelContainer = styled.div`
  border: 1px solid #616467;
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  background: #111111;
`;

const PanelBody = styled.div`
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  ${down('tablet')} {
    min-height: 0;
  }

  ${up('desktop')} {
    min-height: 7rem;
  }
`;

const PanelHeader = styled.div`
  background-color: #292929;
  padding: 0.5rem 0.75rem;
  font-size: 1rem;
  border-bottom: 1px dashed #616467;
  color: white;
`;

export default Panel;
