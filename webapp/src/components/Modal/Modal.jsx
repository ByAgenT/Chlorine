import React from 'react';
import styled from 'styled-components';
import LinkButton from '../LinkButton';

const Modal = ({ display, children, className }) => {
  let [displayStatus, setDisplayStatus] = display;
  displayStatus = displayStatus ? 'block' : 'none';
  return (
    <ModalContainer display={displayStatus}>
      <ModalContent className={className}>
        {children}
        <ModalBottomBar>
          <LinkButton onClick={() => setDisplayStatus(false)}>Exit</LinkButton>
        </ModalBottomBar>
      </ModalContent>
    </ModalContainer>
  );
};

const ModalContainer = styled.div`
  display: ${props => props.display};
  position: fixed; /* Stay in place */
  z-index: 1; /* Sit on top */
  left: 0;
  top: 0;
  width: 100%; /* Full width */
  height: 100%; /* Full height */
  overflow: auto; /* Enable scroll if needed */
  background-color: rgb(0, 0, 0); /* Fallback color */
  background-color: rgba(0, 0, 0, 0.4); /* Black w/ opacity */
`;

const ModalContent = styled.div`
  background-color: #222326;
  color: white;
  margin: 15% auto; /* 15% from the top and centered */
  padding: 20px;
  border: 1px solid #888;
  width: 40%; /* Could be more or less, depending on screen size */

  & > ${LinkButton} {
    margin: 0;
    padding-top: 0.5em;
  }
`;

const ModalBottomBar = styled.div`
  display: flex;
  height: 2.5rem;
  color: white;
  background-color: #292929;
  align-items: center;
`;

export default Modal;
