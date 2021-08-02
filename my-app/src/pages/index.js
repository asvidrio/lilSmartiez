import React, {useState} from 'react'
import Sidebar from '../components/Sidebar'
import Navbar from '../components/Navbar'
import Hero from '../Hero'
import { homeObjOne, homeObjThree, homeObjTwo } from '../components/InfoPages/data'
import InfoPages from '../components/InfoPages'

const Home = () => {
    const [isOpen, setIsOpen] = useState(false)

    const toggle = () => {
        setIsOpen(!isOpen)
    }

    return (
        <>
          <Sidebar isOpen={isOpen} toggle={toggle}/>
          <Navbar toggle={toggle} />
          <Hero /> 
          <InfoPages {...homeObjOne}/>  
          <InfoPages {...homeObjTwo} />  
          <InfoPages {...homeObjThree} />  
        </>
    )
}

export default Home
