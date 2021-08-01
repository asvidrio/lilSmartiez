import React from 'react'
import {FaBars} from 'react-icons/fa'
import logo from '../../icons/logo.svg'
import { Nav, NavbarContainer, NavLogo, MobileIcon, NavMenu, NavItem, NavLinks, NavBtn, NavBtnLink } from './elements';

const Navbar = ({ toggle }) => {
    return (
        <>
            <Nav>
                <NavbarContainer>
                    <NavLogo to='/'>
                        <img src={logo} className="App-logo" alt="logo" width="auto" height="100" margin-top= "20px" margin-left= "20px" position='absolute'/>
                    </NavLogo>
                    <MobileIcon onClick={toggle}>
                        <FaBars />
                    </MobileIcon>
                    <NavMenu>
                        <NavItem>
                            <NavLinks to="home">Home</NavLinks>
                        </NavItem>
                        <NavItem>
                            <NavLinks to="resources">Resources</NavLinks>
                        </NavItem>
                        <NavItem>
                            <NavLinks to="about">About</NavLinks>
                        </NavItem>
                        <NavBtn>
                            <NavBtnLink to="profile">Profile</NavBtnLink>
                        </NavBtn>
                    </NavMenu>
                </NavbarContainer>

            </Nav>
        </>
    )
}

export default Navbar
