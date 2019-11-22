import React from 'react';
import styled from 'styled-components';

const TextInput = ({ onChange, placeholder }) => (
  <div>
    <TextInputInput onChange={onChange} type="text" placeholder={placeholder} />
  </div>
);

const TextInputInput = styled.input`
  width: -webkit-fill-available;
  font-family: 'Josefin Sans', -apple-system, BlinkMacSystemFont, 'Segoe UI',
    'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans',
    'Helvetica Neue', sans-serif;
  font-size: 2em;
  color: green;
  background: none;
  border: none;
  outline: none;
`;

export default TextInput;
