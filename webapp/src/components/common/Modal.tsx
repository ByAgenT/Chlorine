import * as React from 'react';
import styled from 'styled-components';
import LinkButton from './LinkButton';
import { useTranslation } from 'react-i18next';

interface ModalProps extends React.HTMLAttributes<HTMLDivElement> {
  display: [boolean, (newValue: boolean) => void];
  title: string;
}

const Modal: React.FC<ModalProps> = ({ display, children, className, title }) => {
  const { t } = useTranslation();
  let [displayStatus, setDisplayStatus] = display;
  return (
    <ModalContainer display={displayStatus ? 'block' : 'none'}>
      <ModalContent className={className}>
        <ModalHeader>
          <ModalTitle>{title}</ModalTitle>
        </ModalHeader>
        <ModalBody>{children}</ModalBody>
        <ModalBottomBar>
          <LinkButton onClick={() => setDisplayStatus(false)}>{t('modal_exit')}</LinkButton>
        </ModalBottomBar>
      </ModalContent>
    </ModalContainer>
  );
};

interface ModalContainerProps {
  display: string;
}

const ModalContainer = styled.div<ModalContainerProps>`
  display: ${(props) => props.display};
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
  margin: 5% auto; /* 5% from the top and centered */
  border: 1px solid #888;
  width: 40rem; /* Could be more or less, depending on screen size */
  & > ${LinkButton} {
    margin: 0;
    padding-top: 0.5em;
  }
`;

const ModalBody = styled.div`
  padding: 20px;
`;

const ModalHeader = styled.div`
  display: flex;
  height: 3rem;
  color: white;
  border-bottom: 1px dashed #616467;
  background-color: #292929;
  align-items: center;
`;

const ModalTitle = styled.h1`
  color: white;
  padding: 0.5rem 0.75rem;
`;

const ModalBottomBar = styled.div`
  display: flex;
  height: 2.5rem;
  color: white;
  border-top: 1px dashed #616467;
  background-color: #292929;
  align-items: center;
`;

export default Modal;
