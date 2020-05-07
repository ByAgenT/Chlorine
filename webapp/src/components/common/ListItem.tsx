import styled from 'styled-components';

const ListItem = styled.div`
  display: flex;
  min-height: 3rem;
  border-bottom: 1px solid #616467;
  color: white;
  align-items: center;
  padding-left: 10px;
  padding-right: 10px;
  &:hover {
    background-color: #222326;
  }
`;

export default ListItem;
