import React, { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import List from 'src/components/core/List';
import ListItem from 'src/components/core/ListItem';
import { getScripts } from 'src/store/api/actions';

import './Home.scss';

function Home(): JSX.Element {
  const dispatch = useDispatch();

  useEffect(() => {
    (async () => {
      const data = await dispatch(getScripts());
      console.log(data);
    })();

    // fetch('http://localhost:9000/scripts')
    //   .then(res => res.json())
    //   .then(console.log)
    //   .catch(console.log);
  });

  return (
    <div className="home-layout">
      <List id="ddd">
        <ListItem>xxx</ListItem>
        <ListItem>yyy</ListItem>
      </List>
    </div>
  );
}

export default Home;
