import React, {Component} from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import Home from './pages';
import { homeObjOne, homeObjThree, homeObjTwo } from './components/InfoPages/data';


class App extends Component {
  state = {
    resources: []
  }

  componentDidMount() {
    fetch('http://localhost:8080/api/resource/')
    .then(res => res.json())
    .then((data) => {
      this.setState({ resources: data })
    })
    .catch(console.log)
  }

  render () {
  return (
    <Router>
      <Switch>
         <Route exact path="/" component={Home} />
         <Route exact path="/intro" component={homeObjOne} />
         <Route exact path="/resources" component={homeObjTwo} />
         <Route exact path="/login" component={homeObjThree} />
      </Switch>
      {/* <Home /> */}
     </Router>
  );
  }
}

export default App;
