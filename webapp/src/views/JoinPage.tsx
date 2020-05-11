import * as React from 'react';
import { useState } from 'react';
import JoinContainer from '../containers/JoinContainer';
import styled from 'styled-components';
import Panel from '../components/common/Panel';
import TextInput from '../components/common/TextInput';
import LinkButton from '../components/common/LinkButton';
import { RouteComponentProps, withRouter } from 'react-router-dom';
import { ChlorineService } from '../services/chlorineService';

interface JoinPageProps extends RouteComponentProps {
  refreshMember: () => void;
}

const JoinPage: React.FC<JoinPageProps> = ({ history, refreshMember }) => {
  const [roomID, setRoomID] = useState<number | null>(null);
  const [name, setName] = useState<string>('');

  return (
    <JoinContainer>
      <JoinPanel name='Join a Room'>
        <TextInput
          onChange={(event) => {
            setRoomID(Number(event.currentTarget.value));
          }}
          placeholder='Room ID'
        />
        <TextInput
          onChange={(event) => {
            setName(event.currentTarget.value);
          }}
          placeholder='Your Name'
        />
        <LinkButton
          onClick={async () => {
            try {
              await new ChlorineService().joinRoom(roomID, name);
              console.log(refreshMember);
              refreshMember();
              history.push('/viewer');
            } catch (error) {
              console.error(error);
            }
          }}
        >
          Join
        </LinkButton>
      </JoinPanel>
    </JoinContainer>
  );
};

const JoinPanel = styled(Panel)`
  & * {
    margin: 0.7rem 1rem;
  }
`;

export default withRouter(JoinPage);
