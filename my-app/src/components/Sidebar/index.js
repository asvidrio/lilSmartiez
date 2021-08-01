import React from 'react'
import { SidebarContainer, Icon, CloseIcon, SidebarWrapper, SidebarMenu, SidebarLink, SideBtnWrap, SidebarRoute } from './elements'

const Sidebar = ({isOpen, toggle}) => {
    return (
        <SidebarContainer isOpen={isOpen} onClick={toggle}>
            <Icon onClick={toggle}>
                <CloseIcon />
            </Icon>
            <SidebarWrapper>
                <SidebarMenu>
                    <SidebarLink to="home" onClick={toggle}>Home</SidebarLink>
                    <SidebarLink to="resources" onClick={toggle}>Resources</SidebarLink>
                    <SidebarLink to="about" onClick={toggle}>About</SidebarLink>
                    <SidebarLink to="profile" onClick={toggle}>Profile</SidebarLink>
                </SidebarMenu>
                <SideBtnWrap>
                    <SidebarRoute to='/signin'>Sign Up</SidebarRoute>
                </SideBtnWrap>
            </SidebarWrapper>
        </SidebarContainer>
    )
}

export default Sidebar;
