import * as React from 'react';
import { useEffect, useState } from 'react';
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

const JoinPage: React.FC<JoinPageProps> = ({ history, location, refreshMember }) => {
  const [roomID, setRoomID] = useState<number | null>(null);
  const [name, setName] = useState<string>('');

  useEffect(() => {
    // If 'room' query parameter is provided, we need to prefill Room Id text input with this value.
    const queryParams = new URLSearchParams(location.search);
    setRoomID(Number(queryParams.get('room')));
  }, [location]);

  return (
    <JoinContainer>
      <JoinPanel name='Join a Room'>
        <TextInput
          value={roomID ? roomID.toString() : ''}
          onChange={(event) => {
            const targetValue = Number(event.currentTarget.value);
            if (!isNaN(targetValue)) {
              setRoomID(targetValue);
            }
          }}
          placeholder='Room ID'
        />
        <TextInput
          value={name}
          onChange={(event) => {
            setName(event.currentTarget.value);
          }}
          placeholder='Your Name'
        />
        <LinkButton
          onClick={async () => {
            try {
              await new ChlorineService().joinRoom(roomID, name);
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
  & > * {
    margin: 1rem 1.5rem;
  }
`;

export default withRouter(JoinPage);
