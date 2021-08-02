// import React, {Component, useState} from 'react';
// import './App.css';
// import {BrowserRouter as Router, Switch} from 'react-router-dom'
// import Home from './pages';
// import axios from 'axios';

// import Resources from './components/InfoPages/resources';


// class App extends Component {
//   const [books, setBooks] = . useState(null);

//   const apiURL = "https://www.anapioficeandfire.com/api/books?pageSize=30";

//   const fetchData = async () => {
//     const response = await axios.get(apiURL)

//     setBooks(response.data) 
// }

//   return (
//     <div className="App">
//       <h1>Game of Thrones Books</h1>
//       <h2>Fetch a list from an API and display it</h2>

//   // {/_ Fetch data from API _/}
//     <div>
//       <button className="fetch-button" onClick={fetchData}>
//         Fetch Data
//       </button>
//     </div>

//   // {/_ Display data from API _/}
//     <div className="books">
//     // Data from API goes here
//         </div>

// </div>
// )
// }

//   componentDidMount() {
//     fetch('http://localhost:8080/api/resource/')
//     .then(res => res.json())
//     .then((data) => {
//       this.setState({ resources: data })
//     })
//     .catch(console.log)
//   }

//   render () {
//   return (
//     <Router>
//       <Switch>
//       </Switch>
//       <Home />
//     </Router>
//   );
//   }
// }

// export default App;
