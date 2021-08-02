import React from 'react'
import InfoPages from '../components/InfoPages';
import { homeObjThree } from '../components/InfoPages/data';
// import Sidebar from '../components/Sidebar';
// import Sidebar from '../components/Sidebar'
// import { homeObjOne } from '../components/InfoPages/data'
// import InfoPages from '../components/InfoPages'

const SigninPage = props => {
    return (
      <div className="popup-box">
        <div className="box">
          <span className="close-icon" onClick={props.handleClose}>x</span>
          {props.content}
          {/* <Sidebar isOpen={isOpen} toggle={toggle}/> */}
          <InfoPages {...homeObjThree} />  
        </div>
      </div>

    );
  };
   
  export default SigninPage;

