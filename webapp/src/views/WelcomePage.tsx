import * as React from 'react';
import styled from 'styled-components';
import { RouteComponentProps } from 'react-router-dom';
import Button from '../components/common/Button';
import { useTranslation } from 'react-i18next';

const WelcomePage: React.FC<RouteComponentProps> = ({ history }) => {
  const { t } = useTranslation();

  return (
    <WelcomePageContainer>
      <Name>{t('name')}</Name>
      <Description>{t('welcome_description')}</Description>
      <ButtonContainer>
        <Button
          onClick={() => {
            window.location.href = '/login';
          }}
        >
          {t('welcome_create_button')}
        </Button>
        <Button
          onClick={() => {
            history.push('/join');
          }}
        >
          {t('welcome_join_button')}
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
