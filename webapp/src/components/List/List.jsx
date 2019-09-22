import React from 'react';
import styled from 'styled-components';

const List = props => (
  <ListContainer className={props.className}>{props.children}</ListContainer>
);

const ListContainer = styled.div``;

export default List;
