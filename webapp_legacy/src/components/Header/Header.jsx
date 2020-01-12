import React from 'react';
import { withRouter } from 'react-router-dom';
import styled from 'styled-components';
import Brand from './Brand';
import HeaderButton from './HeaderButton';
import HeaderMenu from './HeaderMenu';
import UserInfo from './UserInfo';

const Header = ({ member, refreshMember, history }) => (
  <HeaderContainer>
    <Brand>CHLORINE</Brand>
    {member ? (
      <div>
        <UserInfo name={member.name} />
        <HeaderButton
          onClick={() => {
            function deleteAllCookies() {
              var cookies = document.cookie.split(';');

              for (var i = 0; i < cookies.length; i++) {
                var cookie = cookies[i];
                var eqPos = cookie.indexOf('=');
                var name = eqPos > -1 ? cookie.substr(0, eqPos) : cookie;
                document.cookie =
                  name + '=;expires=Thu, 01 Jan 1970 00:00:00 GMT';
              }
            }
            deleteAllCookies();
            refreshMember();
            history.push('/');
          }}
        >
          Logout
        </HeaderButton>
      </div>
    ) : (
      <HeaderMenu>
        <HeaderButton href="/login">Create</HeaderButton>
        <HeaderButton href="/join">Join</HeaderButton>
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

export default withRouter(Header);
