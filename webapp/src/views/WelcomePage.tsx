import * as React from 'react';
import styled from 'styled-components';
import { RouteComponentProps } from 'react-router-dom';
import Button from '../components/common/Button';

const WelcomePage: React.FC<RouteComponentProps> = ({ history }) => {
  return (
    <WelcomePageContainer>
      <Name>CHLORINE</Name>
      <Description>
        Chlorine helps you provide a music to your party that everyone enjoy. Just create a music
        room using your Spotify account or join already created account. Your friends could join you
        and add their own music without any registration.
      </Description>
      <ButtonContainer>
        <Button
          onClick={() => {
            window.location.href = '/login';
          }}
        >
          Create a room
        </Button>
        <Button
          onClick={() => {
            history.push('/join');
          }}
        >
          Join a room
        </Button>
      </ButtonContainer>
    </WelcomePageContainer>
  );
};

const WelcomePageContainer = styled.div`
  display: flex;
  min-height: 35rem;
  justify-content: center;
  flex-direction: column;
  align-items: center;
`;

const Name = styled.h1`
  padding: 0;
  margin: 0 0 0 1rem;

  font-size: 2.5rem;
  font-weight: 600;
  color: white;
  user-select: none;
`;

const Description = styled.p`
  color: white;
  max-width: 40rem;
  font-size: 1.5rem;
  font-weight: lighter;
`;

const ButtonContainer = styled.div`
  margin-top: 3rem;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  width: 30rem;
`;

export default WelcomePage;
