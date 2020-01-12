import React, { useState } from 'react';
import { withRouter } from 'react-router-dom';
import styled from 'styled-components';
import JoinContainer from '../containers/JoinContainer';
import Panel from '../components/Panel';
import TextInput from '../components/TextInput';
import LinkButton from '../components/LinkButton';
import { joinRoom } from '../services/ChlorineService';

const JoinPage = ({ history }) => {
  const [roomID, setRoomID] = useState('');
  const [name, setName] = useState('');

  return (
    <JoinContainer>
      <JoinPanel name="Join a Room">
        <TextInput
          onChange={event => {
            setRoomID(event.target.value);
          }}
          placeholder="Room ID"
        />
        <TextInput
          onChange={event => {
            setName(event.target.value);
          }}
          placeholder="Your Name"
        />
        <LinkButton
          onClick={async () => {
            try {
              await joinRoom(roomID, name);
              history.push('/viewer');
            } catch (error) {
              console.error(error);
            }
          }}
        >
          Authorize
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
