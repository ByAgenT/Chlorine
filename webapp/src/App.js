import React, { useEffect, useState } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import Header from './components/Header';
import AppStyle from './globalStyle';
import { getMemberInfo } from './services/ChlorineService';
import JoinPage from './views/JoinPage';
import PartyPage from './views/PartyPage';
import ViewerPage from './views/ViewerPage';

const App = () => {
  const [member, refreshMember] = useMemberInformation();

  return (
    <Router>
      <div>
        <Header member={member} refreshMember={refreshMember} />
        <Route path="/player" exact component={PartyPage} />
        <Route path="/join/" exact component={JoinPage} />
        <Route path="/viewer" exact component={ViewerPage} />
        <AppStyle />
      </div>
    </Router>
  );
};

function useMemberInformation() {
  const [member, setMember] = useState(null);

  async function refreshMember() {
    try {
      setMember(await getMemberInfo());
    } catch (error) {
      console.error(error);
    }
  }

  useEffect(() => {
    async function prepare() {
      try {
        setMember(await getMemberInfo());
      } catch (error) {
        console.error(error);
      }
    }
    prepare();
  }, []);

  return [member, refreshMember];
}

export default App;
