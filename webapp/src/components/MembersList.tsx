import * as React from 'react';
import styled from 'styled-components';
import List from './common/List';
import LinkButton from './common/LinkButton';
import ListItem from './common/ListItem';
import { Member } from '../models/chlorine';
import { useTranslation } from "react-i18next";

interface MembersListProps {
  members: Member[];
  onUpdate: (event: React.MouseEvent<HTMLAnchorElement, MouseEvent>) => void;
}

const MembersList: React.FC<MembersListProps> = ({ members, onUpdate }) => {
  const {t} = useTranslation();
  return (
    <MemberListContainer>
      <List>
        {members.map((member) => {
          return <MembersListItem key={member.id}>{member.name}</MembersListItem>;
        })}
      </List>
      <MemberListBottomBar>
        <RoomButton>{t('your_room')} {members[0] ? members[0].roomId : ''}</RoomButton>
        <LinkButton
          onClick={() => {
            alert('In development');
          }}
        >
          {t('settings')}
        </LinkButton>
        <LinkButton onClick={onUpdate}>{t('refresh')}</LinkButton>
      </MemberListBottomBar>
    </MemberListContainer>
  );
};

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
