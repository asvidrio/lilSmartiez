import { Component } from 'react';
import Resources from './resources';

export const homeObjOne = {
    id: 'intro',
    lightBg: true,
    lightText: false,
    lightTextDesc: false,
    topLine: 'Title',
    headline: 'Free database of resources',
    description: 'Get access to our database filled with resources that can help you get your student on track with STEM-principles. Materials range from online courses, video playlists, motivational films, and fun hands-on exercises!',
    buttonLabel: 'Get started',
    imgStart: false,
    img: require('../../icons/books.svg').default,
    alt: 'Books',
    dark: false,
    primary: false,
    darkText: true
}

export const homeObjTwo = {
    id: 'resources',
    lightBg: false,
    lightText: true,
    lightTextDesc: true,
    topLine: 'Resources',
    headline: 'List of resources',
    description: 'Get access to our database filled with resources that can help you get your student on track with STEM-principles. Materials range from online courses, video playlists, motivational films, and fun hands-on exercises!',
    buttonLabel: 'Edit',
    imgStart: true,
    img: require('../../icons/books.svg'),
    alt: '',
    dark: true,
    primary: true,
    darkText: false
}

export const homeObjThree = {
  id: 'login',
  lightBg: true,
  lightText: false,
  lightTextDesc: false,
  topLine: 'Sign In',
  headline: 'Log in or create a new account to save your favorite resources.',
  description: '',
  TextField: 'Enter your username here',
  color: 'white',
  buttonLabel: 'Log in',
  imgStart: true,
  img: require('../../icons/login_pic.svg').default,
  alt: '',
  dark: false,
  primary: false,
  darkText: true
}

class App extends Component {
    render() {
      return (
        <Resources resources={this.state.resources} />
      )
    }
  }

export default App;
