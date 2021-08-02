import React, {useState} from 'react'
import Sidebar from '../components/Sidebar'
import { homeObjOne } from '../components/InfoPages/data'
import InfoPages from '../components/InfoPages'

const List = () => {
    const [isOpen, setIsOpen] = useState(false)

    const toggle = () => {
        setIsOpen(!isOpen)
    }

    return (
        <>
          <Sidebar isOpen={isOpen} toggle={toggle}/>
          <InfoPages {...homeObjOne} />  
        </>
    )
}

export default List;
