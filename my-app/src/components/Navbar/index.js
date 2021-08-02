import React, {useState} from 'react'
import {FaBars} from 'react-icons/fa'
import logo from '../../icons/logo1.png'
import { Nav, NavbarContainer, NavLogo, MobileIcon, NavMenu, NavItem, NavLinks } from './elements';

const Navbar = ({ toggle }) => {
    const [hover, setHover] = useState(false)

    const onHover = () => {
        setHover(!hover)
    }

    return (
        <>
            <Nav>
                <NavbarContainer>
                    <NavLogo to='/'>
                        <img src={logo} className="App-logo" alt="logo" width="auto" height="120" margin-top= "10px" margin-left= "20px" position='absolute'/>
                        <h1 color='white' font-weight='bold' width="auto" height="120" margin-top= "10px" margin-left= "20px" position='absolute'>Lil Smartiez</h1>
                    </NavLogo>
                    <MobileIcon onClick={toggle}>
                        <FaBars />
                    </MobileIcon>
                    <NavMenu>
                        <NavItem>
                            <NavLinks to='' onMouseEnter={onHover} onMouseLeave={onHover} primary='true' dark='false'>Home</NavLinks>
                        </NavItem>
                        <NavItem>
                            <NavLinks to='intro' onMouseEnter={onHover} onMouseLeave={onHover} primary='true' dark='false'>Intro</NavLinks>
                        </NavItem>
                        <NavItem>
                            <NavLinks to='resources' onMouseEnter={onHover} onMouseLeave={onHover} primary='true' dark='false'>About</NavLinks>
                        </NavItem>
                        <NavItem>
                            <NavLinks to='login' onMouseEnter={onHover} onMouseLeave={onHover} primary='true' dark='false' border-radius='20px'>Login</NavLinks>
                        </NavItem>
                    </NavMenu>
                </NavbarContainer>

            </Nav>
        </>
    )
}

export default Navbar
