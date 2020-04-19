import * as React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import JoinPage from './views/JoinPage';
import AppStyle from './globalStyle';
import Header from './components/Header';
import PlayerPage from './views/PlayerPage';
import { useMemberInformation } from './hooks/membership';

const App = () => {
  let [member, refreshMember] = useMemberInformation();

  return (
    <Router>
      <div>
        <Header member={member} refreshMember={refreshMember} />
        <Route path='/player' exact component={PlayerPage} />
        <Route path='/' exact component={JoinPage} />
        <AppStyle />
      </div>
    </Router>
  );
};

export { App };
