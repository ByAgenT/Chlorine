import styled from 'styled-components';
import breakpoint from 'styled-components-breakpoint';

const Control = styled.div`
  border: none;
  border-style: none;
  background-size: cover;
  background-repeat: no-repeat;

  ${breakpoint('desktop')`
    width: 52px;
    height: 52px;
  `}

  ${breakpoint('mobile')`
    width: 26px;
    height: 26px;
  `}
`;

export default Control;
