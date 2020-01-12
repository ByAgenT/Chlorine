import styled from 'styled-components';

const LinkButton = styled.a`
  user-select: none;
  color: white;
  font-size: 1.15rem;
  font-weight: 600;
  margin-left: 1rem;
  margin-right: 1rem;

  text-decoration: none;
  text-align: center;

  &:hover {
    color: rgb(29, 185, 84);
  }
`;

export default LinkButton;
