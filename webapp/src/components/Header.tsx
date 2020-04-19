import * as React from 'react';
import { RouteComponentProps, withRouter } from 'react-router-dom';
import styled from 'styled-components';
import LinkButton from './LinkButton';
import { Member } from '../models/chlorine';

interface UserInfoProps {
  name: string;
}

interface HeaderProps extends RouteComponentProps {
  member?: Member;
  refreshMember?: () => void;
}

const Header: React.FC<HeaderProps> = ({ member, refreshMember, history }) => (
  <HeaderContainer>
    <Brand>CHLORINE</Brand>
    {member ? (
      <div>
        <UserInfo name={member.name} />
        <HeaderButton
          onClick={() => {
            function deleteSessionCookie() {
              document.cookie = 'chlorine_session=;expires=Thu, 01 Jan 1970 00:00:01 GMT;';
            }
            deleteSessionCookie();
            refreshMember();
            history.push('/');
          }}
        >
          Logout
        </HeaderButton>
      </div>
    ) : (
      <HeaderMenu>
        <HeaderButton href='/login'>Create</HeaderButton>
        <HeaderButton href='/join'>Join</HeaderButton>
      </HeaderMenu>
    )}
  </HeaderContainer>
);

const HeaderContainer = styled.header`
  background-color: black;
  width: 100%;
  height: 4.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

const UserInfo: React.FC<UserInfoProps> = ({ name }) => (
  <UserInfoSpan>{`Hello, ${name}`}</UserInfoSpan>
);

const UserInfoSpan = styled.span`
  font-size: 1.15rem;
  color: white;
  margin-right: 1rem;
  margin-left: 1rem;
`;

const Brand = styled.span`
  padding: 0;
  margin: 0;
  margin-left: 1rem;

  font-size: 1.8rem;
  font-weight: 600;
  color: white;
  user-select: none;
`;

const HeaderButton = styled(LinkButton)``;

const HeaderMenu = styled.div`
  margin-right: 1rem;
  margin-left: 1rem;
`;

export default withRouter(Header);
