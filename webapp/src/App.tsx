import * as React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import JoinPage from './views/JoinPage';
import AppStyle from './globalStyle';
import Header from './components/Header';
import PlayerPage from './views/PlayerPage';
import { useMemberInformation } from './hooks/membership';
import ViewerPage from './views/ViewerPage';
import WelcomePage from './views/WelcomePage';

const App = () => {
  let [member, refreshMember] = useMemberInformation();

  return (
    <Router>
      <div>
        <Header member={member} refreshMember={refreshMember} />
        <Route path='/player' exact component={PlayerPage} />
        <Route path='/' exact component={WelcomePage} />
        <Route
          path='/join'
          exact
          render={(props) => <JoinPage {...props} refreshMember={refreshMember} />}
        />
        <Route path='/viewer' exact component={ViewerPage} />
        <AppStyle />
      </div>
    </Router>
  );
};

export { App };
