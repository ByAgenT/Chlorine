import * as React from 'react';
import styled from 'styled-components';

interface TextInputProps {
  onChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
  placeholder?: string;
  value?: string;
  width?: string;
}

// TODO: fix types.
const TextInput: React.FC<TextInputProps> = ({ onChange, placeholder, value, width }) => (
  <div>
    <TextInputInput
      onChange={onChange as any}
      type='text'
      placeholder={placeholder}
      value={value}
      width={width}
    />
  </div>
);

const TextInputInput = styled.input`
  width: ${(props) => (props.width ? props.width : '-webkit-fill-available')};
  width: ${(props) => (props.width ? props.width : '-moz-available')};
  font-size: 2em;
  color: white;
  background: none;
  border: none;
  border-bottom: 1px dashed #616467;
  outline: none;
`;

export default TextInput;
