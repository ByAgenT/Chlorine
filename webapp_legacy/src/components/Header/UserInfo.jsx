import React from 'react';
import styled from 'styled-components';


const UserInfo = ({name}) => (
  <UserInfoSpan>
    {`Hello, ${name}`}
  </UserInfoSpan>
);

const UserInfoSpan = styled.span`
  font-size: 1.15rem;
  color: white;
  margin-right: 1rem;
  margin-left: 1rem;
`;

export default UserInfo;