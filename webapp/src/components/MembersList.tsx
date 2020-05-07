import * as React from 'react';
import styled from 'styled-components';
import List from './common/List';
import LinkButton from './common/LinkButton';
import ListItem from './common/ListItem';
import { Member } from '../models/chlorine';

interface MembersListProps {
  members: Member[];
  onUpdate: (event: React.MouseEvent<HTMLAnchorElement, MouseEvent>) => void;
}

const MembersList: React.FC<MembersListProps> = ({ members, onUpdate }) => (
  <MemberListContainer>
    <List>
      {members.map((member) => {
        return <MembersListItem key={member.id}>{member.name}</MembersListItem>;
      })}
    </List>
    <MemberListBottomBar>
      <RoomButton>Your Room: {members[0] ? members[0].roomId : ''}</RoomButton>
      <LinkButton
        onClick={() => {
          alert('In development');
        }}
      >
        Settings
      </LinkButton>
      <LinkButton onClick={onUpdate}>Refresh</LinkButton>
    </MemberListBottomBar>
  </MemberListContainer>
);

const RoomButton = styled(LinkButton)`
  &:hover {
    color: white;
  }
`;

const MembersListItem = styled(ListItem)`
  font-size: 1.5em;
`;

const MemberListBottomBar = styled.div`
  display: flex;
  height: 2.5rem;
  color: white;
  background-color: #292929;
  align-items: center;
`;

const MemberListContainer = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  flex-grow: 1;
`;

export default MembersList;
