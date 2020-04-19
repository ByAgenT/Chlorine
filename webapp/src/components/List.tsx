import * as React from 'react';
import styled from 'styled-components';

const List: React.FC<React.HTMLAttributes<HTMLDivElement>> = (props) => (
  <ListContainer className={props.className}>{props.children}</ListContainer>
);

const ListContainer = styled.div``;

export default List;
